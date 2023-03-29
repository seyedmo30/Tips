<b> first_class_functions</b>

بعضی زبان های برنامه نویسی قابلیتی دارند که می توان، فانکشن را درون یک متغییر ریخت ، به قابلیت فرس کلاس فانکشن می گوند


https://golangbot.com/first-class-functions/


 <hr>

<b> Higher Order Functions</b>



 تابعی که بتوان  یک تابع به عنوان پارامتر بگیرد ، یا بتوان خروجی اش یک تابع باشد، از این قابلیت برخوردار است
 
 توجه : این به این معنی نیست که تابع مقدار دهی شده ( کانکریت) در ورودی یا خروجی استفاده شود ، این یعنی تایپ آن فانکشن است
 
 https://www.golangprograms.com/higher-order-functions-in-golang.html
 
 <hr>
 
         package main

        import (
            "fmt"
            "net/http"
        )

        func hello(w http.ResponseWriter, req *http.Request) {

            fmt.Fprintf(w, "hello\n")
        }

        func headers(w http.ResponseWriter, req *http.Request) {

            for name, headers := range req.Header {
                for _, h := range headers {
                    fmt.Fprintf(w, "%v: %v\n", name, h)
                }
            }
        }

        func main() {

            http.HandleFunc("/hello", hello)
            http.HandleFunc("/headers", headers)

            http.ListenAndServe(":8090", nil)
        }
 
