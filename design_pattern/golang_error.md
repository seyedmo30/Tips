

stack trace error   ردپای خطا

به این معناست که بفهمیم خطا از کجا شروع شده ، یعنی اگر پیام خطا در بیرونی ترین لایه ی برنامه نمایش داده شد ، بتوانیم اولین مرحله بروز خطا را پیدا کنیم

go get github.com/pkg/errors

در گولنگ این پکیج مناسب است امام مشکل اینجاست که برای تعریف ارور باید به روش پکیج ، ارور را تعریف کنیم . اگر خود سازنده ی پکیج هستیم ، می توان این ساختار را رعایت کرد اما اگر استفاده کننده از یه پکیج هستیم ، باید برای رپ کردن ابتدا ارور را به فرمت مورد نظر تبدیل کنیم: 

errors.Errorf("Could not write to file")




### Sentinel Errors خطای نگهبان

یکی از راه های پیش بینی خطا ، این است که خطا را ابتدا تعریف کنیم ( پیام خطا را ) و در کد ، متوانیم با استفاده از IS چک کنیم که خطای پیش آمده برابر خطای پیش بینی شده است یا خیر .

بدی این روش این است که باید تمامی خطا ها را از پکیج آن امپورت کنیم و شاید این روند خیلی پیچیده و بزرگ شود

### Custom Error Types تایپ دستی خطا

به این نتیجه رسیدیم تو لایه های ریپازیتوری یا http call ها بهتره علاوه بر کد هم داشته باشیم ، در این صورت فانکشنی که قرار ارور هندل کنه می تونه از این دستور استفاده کنه :

if errors.As(err, &customErr) {
با این روش تایپ اسرشن میکنیم به اون استراکت و از کد خطا استفاده میکنیم


یادمون باشه تست فقط کافی نیست ، تست میگه کار میکنه یا نه ولی نیاز به metric هم داریم تا بتونیم ببینیم با چه سرعتی کار میکنه



در این روش می توانیم با توجه با اینترفیس ارور ، علاوه بر پیام ارور ، مقدار های دیگری هم به استراکت بیفزاییم ( یعنی ساختار خطا را تغییر دهیم) . 

حال در صورت رخ دادن ارور ، با استفاده از errors.As چک کنیم ساختار ارور ، با ساختار پیش بینی شده یکسان است یا خیر .

در این روش هم می بایست تایپ های ارور را پابلیک کرد و در تمامی کد باید فراخوانی شود و این خود مشکل است .

https://earthly.dev/blog/golang-errors/


بهتره ارور ساختار زیر رو داشته باشه :

```go
// built-in interface
type error interface {
	Error() string
}

// custom struct error
type AppError struct {  
	Code        int  
	Description string  
	Message     string  
}  
```

دلایل :

+ گاهی در یک لایه امکان دارد چندیدن حالت از خطا رخ دهد و در لایه بالا تر باید تشخیص دهیم چه اکشنی نشان دهیم مثال ها :

+ +  در ریپو امکان داره getItem  بزنیم و پاسخ نیاد ، حال می تونیم یا خطا در سینتکس بخوریم ، امکان داره اون مورد در تیبل نباشه ، شایدم اون لحظه کانکشن قطع شده باشه ،* باید ۳ اکشن متفاوت داشته باشیم

+ + در apiCall ها حالت های مختلفی از خطا داشته باشیم

پس بهتره در لایه ی بیزینس لاجیک بدونیم هر لایه که خطا داره دلیلش چیه



### Error wrapping بسته بندی خطا

در این روش ، می توانیم هر جا که به ارور خوردیم در پیام آن ، یک لایه متن ، رپ کنیم و به مرحله بعد دهیم ، هر موقع کاربر خواست می تواند لایه های خطا را ببیند 

توجه شود که wrap کردن با Sentinel Errors و Error Types ناسازگار است  

در مجموع به نظر خودم بهترین راه ، این است که در هر مرحله که err != nil  ، همان جا را لاگ کنیم .اگر ارور را به فانکشن بالا ریترن کنیم باز هم لاگ می شود و در نتیجه رپ کردن در خود لاگ ها مشاهده می شود


### 1404.3.18

 به بهترین ارور هندل تو گولنگ  - البته تا اینجای کار - رسیدیم ، به طوری که توی هر لایه - ریپوزیتوری یا درخواست از بانک - همون جا ارور کامل میشه - متن فارسی - استتوس - دیسکریپشن که می تونه یا بادی پاسخ باشه و یا خطای پستگرس باشه - و در نهایت استک که از ران تایم میزاریم توش - و همچنین در موارد استسنا از لایه درونی مانند ریپازیتوری تا یوز کیس اگر خطا این بود که در مخزن یافت نشد - باید در لایه پایانی **overwrite** شه به خطای داخلی 500 رخ داده ، اما استک و دیسکریپشن تغییری نکنه
  

### انواع خطا در گولنگ

+ Runtime error 

خطا هایی که بعد از کامپایل شدن برنامه رخ می دهند یعنی مشکل سینتکسی یا غیر در شروع نبوده و متداول ترین ها ش : 

+ Undefined variable/function

متغیر یا فانکشنی اطتفاده شود در حالی که تعریف نشده

+ Multiple-value  in single-value 

اصولا زمانی که خروجی یک فانکشن درون یک متغییر ریخته شود و تعداد خروجی های فانکشن بیش از ظرفیت متغیر باشد

+ Undefined reference

زمانی که فانکشن پیدا نشود یا ایمپورت نشود

+ syntax error: unexpected علامت
زمانی که اشتباه سینتکسی تو کد باشه

+ Panic: runtime error: index out of range

زمانی که ایندکس مورد نظر در رشته نباشد


**github.com/pkg/errors**

```go
import "github.com/pkg/errors"

var ErrS1 = errors.New("operation failed due to a bad request")   // باید اول اسم حتما Err باشه
 
 	if err := riskyOperation(); err != nil {
		return errors.Wrap(err, "failed to do something") // wrapper
	}


    if errors.Is(err, ErrGG) { // Unwrap the error 
        fmt.Println("salam") // حتی اگه خطا رپ شده هم باشد باز هم خروجی میده
    }

```

![Description of the image](https://github.com/seyedmo30/Tips/blob/main/static/1_CcZQx10nlxPi2YD3x5fUpg.png)

## tips

### 

