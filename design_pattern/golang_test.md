#  test


### unit test

بخش های بسیار کوچک و فانکشن ها را با این روش تست می کنیم .

در این روش ، ورودی و خروجی و نتیجه ی فانشکن را از قبل مشخص می کنیم و تست را اجرا می کنیم .

### integration test

گاهی اوقات تست به پیش نیاز هایی وابسته است ، مثلا برای تست یک ریکویست ، با وظیفه ثبت در دیتابیس یا خطا ، ما نیاز داریم که علاوه بر مشخص کردن ورودی و خروجی ، کانکتور ها ( یا dependency ها ) را هم مشخص کنیم . در این صورت می توانیم ابتدا تمامی کلاس ها یا استراکت ها را معرفی کنیم . مثلا تنظیمات کانکشن به دیتا بیس یا تنظیمات روتر . 
```
 
      type ParserTestSuite struct {
        suite.Suite
        Parser protocol.Parser
      }

      func (s *ParserTestSuite) SetupTest() {
        parser := viper_parser.Parser{}
        s.Parser = &parser

      }
      func TestExampleTestSuite(t *testing.T) {
        suite.Run(t, &ParserTestSuite{})
      }

      func (s *ParserTestSuite) TestMigration() {
        conn := NewConnectionPostgres(s.Parser)
        err := conn.Migration(context.Background())
        s.Nil(err)

      }

```

می توان در روت پروژه کد را زد و تمامی تست ها انجام شود : 
```
   go test -v ./... 
```

توجه شود می توان کانفیگ را به صورت nested پیدا کرد ، در این صورت نیازی نیست در تمامی دارکتوری ها config  را قرار داد .

# metrics

### pprof ـ portabl profile 
ابتدا باید این پکیج را نصب کنیم
   
  `  sudo apt install graphviz `
   
حال باید در یک گوروتین جدا ، لیسن به کد کنیم 

```go
    import (
    	"net/http"
    	_ "net/http/pprof"
    )

    go func() {
     http.ListenAndServe("localhost:6060", nil)
    }()

```    

پروفایل مموری: 

`http://localhost:6060/debug/pprof/heap`

پروفایل CPU:

 `http://localhost:6060/debug/pprof/profile?seconds=30`

پروفایل گوروتین‌ها: 

`http://localhost:6060/debug/pprof/goroutine`



و همزمان در حلقه چندین بار فانکشن هایی که احتمالا گلوگاه کد هستند را فراخانی کنیم . کد حداقل باید 30 ثانیه در حلقه تکرار شود ، همزمان کد زیر را اجرا می کنیم تا کد را از طریق ListenAndServe کد را ذخیره کند :

`go tool pprof http://localhost:6060/debug/pprof/profile`

اگر بتواند 30 ثانیه آنالیز کند ، آنگاه می توانیم به روش های گوناگون خروجی را بگیریم ولی متداول ترین روش ، وب است:

    (pprof) web


زمانی که بخواهیم توی وب ببینیم

`go tool pprof -http=:8080 http://localhost:6060/debug/pprof/heap`


### TotalAlloc


```go
import (
    "runtime"
    "log"
)

var memStats runtime.MemStats
runtime.ReadMemStats(&memStats)
log.Printf("Alloc = %v, TotalAlloc = %v, Sys = %v, NumGC = %v", memStats.Alloc, memStats.TotalAlloc, memStats.Sys, memStats.NumGC)
```

### pprof goroutines

گاهی شایدی تعداد زیادی گوروتین باز شده ولی استفاده نمی شه

با استفاده از **pprof** میشه دید چه تعداد گوروتین بازه

### Prometheus و Grafana

باید اطلاعات سیستم عامل به این ابزار فرستاده شود تا مصرف کنابع دیده شود

### repository layer test

+ **Isolate Tests**

بهتره تست ها جوری باشن که نیازی به ترتیب و پیش نیاز نباشه ، هر تست ، با هر ترتیبی بتونه اجرا بشه

+ **Use a Test Database**

بهتره روی یه دیتابیس دیگه تست انجام بشه


## tips

در مانی تست ها همه در دایرکتوری تست قرار دارد  ، خوبیش اینه که همه تست ها یه جاست و با خاندن جیسون ، تمام سید ها و داده های تست در دسترسه

ولی مشکل اینجاست که تنها global variable رو میشه تست کرد و نمی توان مقادیر داخلی رو ایمپورت کرد

### **tdd** test driven development

یک رویکرد برای تست نویسی هست و ابتدا تست نوشته میشه ، بعد کد زده میشه

### Table-Driven Tests

یه جدول هست شامل ورودی ها ، خروجی ها و درستی یا غلطی خروجی است

