
# ۱ با احتیاط از مقدار دهی سریع استفاده کنیم

گاهی ممکن است یک داده چند بار مفدار دهی شود، بهتر است در این شرایط از مقدار دهی سریع استفاده نکینم


```
var client *http.Client ❶
if tracing {
 client, err := createClientWithTracing() ❷
 if err != nil {
 return err
 }
 log.Println(client)
} else {
 client, err := createDefaultClient() ❸
 if err != nil {
 return err
 }
 log.Println(client)
}
```

کد زیر بهتر است

```
var client *http.Client
var err error ❶
if tracing {
 client, err = createClientWithTracing() ❷
 if err != nil {
 return err
 }
} else {
 // Same logic
}

```






# ۲ برای خوانایی کد، بهتر است از کد های تو در تو پرهیز کنیم

در صورتی که امکان دارد می توانیم بجای استفاده از ایف و الس ، تنها از ایف استفاده کنیم


```
func join(s1, s2 string, max int) (string, error) {
 if s1 == "" {
 return "", errors.New("s1 is empty")
 } else {
 if s2 == "" {
 return "", errors.New("s2 is empty")
 } else {
 concat, err := concatenate(s1, s2) 
 if err != nil {
 return "", err
 } else {
 if len(concat) > max {
 return concat[:max], nil
 } else {
 return concat, nil
 }
 }
 }
 }
}
```
به شکل زیربنویسیم

```
func join(s1, s2 string, max int) (string, error) {
 if s1 == "" {
 return "", errors.New("s1 is empty")
 }
 if s2 == "" {
 return "", errors.New("s2 is empty")
 }
 concat, err := concatenate(s1, s2)
 if err != nil {
 return "", err
 }
 if len(concat) > max {
 return concat[:max], nil
 }
 return concat, nil
}
```



# ۳ باید با دقت از تابع اینیت استفاده کنیم ، زیرا چند مشکل احتمال دارد به وجود بیاید

الف ـ اینیت همیشه اجرا می شود و شاید کد ما در جایی که انتظار نداشته باشیم یا نیاز نباشد ، فرا خوانی شود و نا خواسته اینیت اجرا شود

ب ـ در اینیت نمی توانیم ارورهندلینگ را به خوبی پیاده سازی کنیم ، تنها می توانیم پنیک کنیم که شاید خیلی خطرناک باشد

ج ـ اینیت حتی در تست هم اجرا می شود ، در حالی که شاید ما می خواهیم با استفده از یونیت تست بخش کوچکی رو تست کنیم و اینیت رو نیاز نداریم اما اجرا می شود و شاید پنیک دهد

د ـ تابع اینیت چون نمی توان چیزی ریترن کرد ، می بایست در یک گلوبال وریبل داده خود را بریزد ، و این خود مشکل ساز می توان باشد زیرا هم توابع دیگر می توانند این وریبل را تغییر دهند ، هم در صورتی که یک تست بخواهد از این وریبل استفاده کند ، دیگر اوزوله نیست( یونیت نیست)

کدی که از اینیت بد استفاده کرده

```
var db *sql.DB
func init() {
 dataSourceName :=
 os.Getenv("MYSQL_DATA_SOURCE_NAME") 
 d, err := sql.Open("mysql", dataSourceName)
 if err != nil {
 log.Panic(err)
 }
 err = d.Ping()
 if err != nil {
 log.Panic(err)
 }
 db = d 
}
```

و اصلاح آن

```
func createClient(dsn string) (*sql.DB, error) { 
 db, err := sql.Open("mysql", dsn)
 if err != nil {
 return nil, err 
 }
 if err = db.Ping(); err != nil {
 return nil, err
 }
 return db, nil
}
```

