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
