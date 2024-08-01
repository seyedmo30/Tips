<div dir='rtl' align='right'>

### circuit breaker pattern

در صورتی که یکی از سرویس ها یا کامپوننت ها down شد یا پاسخ ها با تایم اوت بر میگشت ، همه ی درخواست ها رو نفرستیم که خطا رو هر بار نمایش بدیم ، همون اول جلو ورود درخواست رو می بندیم که منتظر پاسخ خطا از سوی سرویس نباشیم ، همچنین ۳ حالت دارد ، حالت باز که پاسخ ها موفق هست ، حالت بسته که درخواست دیگه نمیره به سرویس ، حالت نیمه باز که با یه لاجیکی در بازه زمانی ، تعدادی پاسخ محدود میفرستیم ، اگر درست شده بود ، باز میکنیم

### chatty io

میگه تا جایی که ممکنه کانکشن های زیاد نزنیم برای گرفتن داده ها  از یه منبع بجاش از بچ یا بالک یا ... استفاده کنیم ، به تعریف دیگر هر در خواست به خارج از سرویسمون ، خارج از گو نیاز به کانکشن و io  داره
باید اونها رو بهینه کرد 


### scope architecture

**namespace**

برای دسته بندی کردن و مرتب کردن کد های طوری که در اسم گذادی ها کانفلیکت نخوریم ، این مفهوم بیشتر در c , c++ , java  کاربرد دارد

همچنین وقتی صحبت در اسکوپ namespace می شه منظور در مقیاس یا دسترسی کلاس ها ، فانکشن ها و وریبل های یک ماژول صحبت میشه

**ماژول**

برای سازمان دهی فانکشن ها همچنین برای ری یوزبل کردن آنها که می توان در یک یا چند فایل بیان ، که از کلاس ها ، فانکشن ها و مقدار ها ساخته شده

**پکیج**

برای سازماندهی و مدیریت ارتباط چند ماژول است ، پکیج  می توان چند ماژول یا ساب پکیج در خود داشته باشد

**لایبری**

کد هایی از پیش نوشته شده هستند که می توان در پروژه ها استفاده نمود ، تمرکز لایبری ، ارتباط آسان و فراهم کردن فانکشن یا توضیحات برای راحت سازی و استفاده در پروژه های متنوع است
لیبری می توان به صورت تنها باشد یا در یک پکیج استفاده شود

### نحوه سنجش عملکرد کد

**Latency**

مدت زمان بین درخواست و پاسخ چند ثانیه است

**Troughput _ rps**

چه تعداد درخواست رو میتونه هندل کنه

**load test**

برای این منظور از load test , stress test  استفاده میکنیم

یکی از ابزار های خوب apache bench mark است 

همچنین میتونیم از ابزار locust استفاده کنیم ، گرافیکیه ، اطلاعات رو ذخیره میکنه و به صورت کانکارنسی تمام core ها رو درگیر میکنه


می توان latancy رو چند حالت اندازه گیری کرد ابتدایی ترین راه ، میانگین است یعنی مجموع ریکوست ها تقسیم بر تعداد اما راه بهتر upper 90
, upper 99 است


**Data dummy _ seed**

داده هایی که فقط برای پر کردن استفاده میشن 


### Scaleup 

**vertical**

در این حالت تنها single instance  داریم و بهش cpu . Ram . Hard بیشتر میدیم

**horizontal**

در این حالت تعداد instance هامون رو بالا میبریم



### saga

ساگا در ساختار تیبل ها تغییری ایجاد نمی کند ، یا فیلدی اضافه نمی کند ، بلکه یک رویکرد است و سرویس ساگا کوردینیتور باید با ارسال دستور به میکروسرویس ها ، آن ها را به مرحله ی بعد ببرد یا رول بک بزنند

در حقیقت باید هر میکرو سرویس ،قابلیت رول بک داشته باشه مثلن اگر داشته باشیم ثبت سفارش ، باید همچنین داشته باشیم کامپنسیت سفارش 

در یک مثال chatgpt

```

package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
)

type OrderRequest struct {
    OrderID    int     `json:"order_id"`
    UserID     int     `json:"user_id"`
    ProductID  int     `json:"product_id"`
    Quantity   int     `json:"quantity"`
    TotalPrice float64 `json:"total_price"`
}

type InventoryRequest struct {
    ProductID int `json:"product_id"`
    Quantity  int `json:"quantity"`
}

type PaymentRequest struct {
    OrderID int     `json:"order_id"`
    Amount  float64 `json:"amount"`
}

func main() {
    order := OrderRequest{
        OrderID:    1,
        UserID:     1,
        ProductID:  1,
        Quantity:   2,
        TotalPrice: 200.00,
    }

    err := createOrderSaga(order)
    if err != nil {
        fmt.Printf("Saga failed: %v\n", err)
    } else {
        fmt.Println("Saga completed successfully")
    }
}

func createOrderSaga(order OrderRequest) error {
    // Step 1: Create Order
    if err := createOrder(order); err != nil {
        return err
    }

    // Step 2: Reserve Inventory
    if err := reserveInventory(order.ProductID, order.Quantity); err != nil {
        compensateOrder(order.OrderID)
        return err
    }

    // Step 3: Process Payment
    if err := createPayment(order.OrderID, order.TotalPrice); err != nil {
        compensateInventory(order.ProductID, order.Quantity)
        compensateOrder(order.OrderID)
        return err
    }

    // All steps succeeded
    return nil
}

func createOrder(order OrderRequest) error {
    orderURL := "http://order-service/orders"
    return sendRequest("POST", orderURL, order)
}

func reserveInventory(productID, quantity int) error {
    inventoryURL := "http://inventory-service/inventory"
    req := InventoryRequest{ProductID: productID, Quantity: quantity}
    return sendRequest("POST", inventoryURL, req)
}

func createPayment(orderID int, amount float64) error {
    paymentURL := "http://payment-service/payments"
    req := PaymentRequest{OrderID: orderID, Amount: amount}
    return sendRequest("POST", paymentURL, req)
}

func compensateOrder(orderID int) {
    orderURL := fmt.Sprintf("http://order-service/orders/%d", orderID)
    sendRequest("DELETE", orderURL, nil)
}

func compensateInventory(productID, quantity int) {
    inventoryURL := "http://inventory-service/inventory/compensate"
    req := InventoryRequest{ProductID: productID, Quantity: quantity}
    sendRequest("POST", inventoryURL, req)
}

func sendRequest(method, url string, body interface{}) error {
    jsonBody, err := json.Marshal(body)
    if err != nil {
        return err
    }

    req, err := http.NewRequest(method

```

### Auth Messages in Microservices

**trust zone**

در میعماری به محدوده ای که می دونیم داده در صورت ورود صحیح _ valid  _  است و نیاز نیست در تمامی میکرو سرویس ها هر بار داده چک ولید شود می گویند

**zero trust**

به علت اهمییت زیاد و این که مطمعن باشیم علاوه بر چک سرویس درخواست کننده ، خود هم دوباره از صحت داده مطمعن شویم ، دوباره ولید بودن داده را چک میکنیم و یعنی به هیچ سرویسی اعتماد نداریم

یه رویکرد امنیتیه که میگه بر خلاف دید سنتی ، تنها هنگام ورود از صحت چک نشه بلکه در هر سرویس خواست استفاده بشه ، چک بشه و جمله ی معروف :

never trust, always verify

**tokens**

یکی از راه های خوب برای احراز هوییت ، اینه که در هر مرحله که امنیت مهم است ، توکن درخواست را ولید کرده و با استفاده از اطلاعات غیر مستقیم توکن ، درسترسی را چک کنیم

 **مثال** 

توجه شود gateway تنها میتواند توکن رو چک کنه ، مثلن وقتی درخواست تنها بر اساسه uuid باشه ، نیازی نیست دابل چک بشه چون احتمال تولید چنین شناسه ای صفره ، ولی وقتی درخواست دریافت اطلاعات فلان کد ملی میشه ، علاوه بر چک gateway , باید سرویس دوباره چک کنه کاربر می تونه این درخواست رو بزنه یا نه
تنها چیزی که gateway تولید میکنه ، metadata هست و میشه به آن اطمینان کرد و با استفاده از این ، درخواست های دابل چک کرد ، به نظرم وظیفه ی ولید بودن درخواست در لایه ی یوزکیس نیست ، در لایه ی دلیوری یا فریمورک است




</div>