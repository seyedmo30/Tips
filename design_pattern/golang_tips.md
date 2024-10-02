## کامل شود

 یه موضیع مبهم برام اینه که گاهی توی میکرو سرویس ها ، یه سرویس می نویسم که مثلن توی پروایدر یه Item Draft (like signature digital or promissory note) می سازیم ، یعنی پیش  نویس یه چیزی و قراره تو مراحل ( ...استپ ) بعد کامل شه ، در این حالت اگر سرویس درخواست کنند برای هر درخواست شناسه داشته باشه ، حتی بعد قطع ارتباط ، دوباره در خواست با شناسه می ده ، اگه جدید بود یعنی در خواست جدیده اما اگر قدیمی بود ، ادامه ی داستان رو میریم

 حالا فرض کنیم کلاینت ما **stateLess** هست و حاضر نیست شناسه جنریت و قبل درخواست ذخیره کنه

 حال ما مشکلمون اینه که نمی دونم درخواست برای بار اوله یا یه بار درخواست هندل شده ولی ریسپانس به کلاینت نرسیده در نتیجه درخواست تکراریه ، ۲ راه داریم :

 + با فرض اینکه می تونه تکراری باشه ، بسازیم ، استتوس **draft** بدیم و اگر دوباره درخواست اومد ، ادامه ی همین رو بریم ، در اینجا چون شناسه خاصی نداریم باید **find**  کنیم آخرین درخواست با استتوس مورد نظر

 + حتی اگر بار اول بود یا تکراری ، تاجایی که هزینه ی زیادی نداره ، ما هم استیت لس باشیم و تا چیزی قطعی نشده ، چیزی ثبت نکنیم

 اینجاست که باید بررسی کنیم هر پیش نویس آیا هزینه داره ، آیا پردازش سنگین داره ، یا پرایدر در پایان روز پیش نویس رو حذف می کنه  یا تعداد محدودی پیش نویس می تونیم تولید کنیم



### PREAPARED_STATEMENT

این روش برای پرفورمنس و حمله  SQL injection attacks محافظت می کند ، اما خیلی از کتابخانه ها و orm ها پشتیبانی نمی کنند ، مثلن نمی توان dynamic array  استفاده کرد همچنین از فانکشن های خوب sqlx دیگه نمی شه استفاده کرد

## json

+ یکی از ابزار های خوب برای کار با جیسون آنمارشال است اما گاهی جیسون ما خیلی بزگ است و یا گاهی می خواهیم بخشی از جیسون رو استفاده کنیم در این صورت نیازی نیست تمامی جیسون را در مموری نگه داریم و راه حل  استفاده از json.Decoder

+ **first_class_functions**

در زبان گو ، فانکشن ها first class citizen  هستند

بعضی زبان های برنامه نویسی قابلیتی دارند که می توان، فانکشن را درون یک متغییر ریخت ، به قابلیت فرس کلاس فانکشن می گوند


https://golangbot.com/first-class-functions/


+ **Higher Order Functions**

 تابعی که بتوان  یک تابع به عنوان پارامتر بگیرد ، یا بتوان خروجی اش یک تابع باشد، از این قابلیت برخوردار است
 
 توجه : این به این معنی نیست که تابع مقدار دهی شده ( کانکریت) در ورودی یا خروجی استفاده شود ، این یعنی تایپ آن فانکشن است
 
 https://www.golangprograms.com/higher-order-functions-in-golang.html

```
         package main

        import (
            "fmt"
            "net/http"
        )

        func hello(w http.ResponseWriter, req *http.Request) {

            fmt.Fprintf(w, "hello\n")
        }



        func main() {

            http.HandleFunc("/hello", hello)

            http.ListenAndServe(":8090", nil)
        }
 ```
تنها تایپ hello به عنوان ورودی به HandleFunc داده میشود   


+ **swagger-echo**

این مثال بر اساس اکو نوشته شده است

ابتدا باید برای تابع مین داکیومنت های زیر را اضافه کنیم

```


    import (
    "github.com/swaggo/echo-swagger"
    _ "aggrigation/docs"
    )

    // @title          title API
    // @version        1.0
    // @description    description.com
    // @query.collection.format multi
    func main() {
    	e.GET("", func(c echo.Context) error { return c.Redirect(http.StatusSeeOther, "/swagger/index.html") })

	    e.GET("/swagger/*", echoSwagger.WrapHandler)
    
    }
```


## refrence

### tips
https://github.com/func25/go-practical-tips/blob/main/tips.md#readability--maintainability


### type comperable

تایپ استراکت در گو comperable هستن اگر همه ی فیلد هاشون برابر باشه ، آنگاه == , != نتیجه میده
 . ولی تایپ فانکشن compareble    نیست
