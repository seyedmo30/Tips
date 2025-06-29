


#### export path

بعد از نصب گولنگ ،فایل اجرا کننده کد های گولینگ ،  فایلی که باید در bashrc , zshrc  باید اضافه بشه در این مسیر است - /usr/local/go/bin/



	export PATH=$PATH:/usr/local/go/bin
	export PATH=$PATH:$GOPATH/bin
 
همچنین باینری هایی ( تولز ) که استفاده میشه مثل pprof  در این آدرس هست 
/usr/local/go/pkg/tool/linux_amd64/

اما پکیج هایی که ایمپورت می کنیم 
/home/seyed/go/pkg/mod/github.com/


و یا باینری هایی مانند swag در این آدرس 
/home/seyed/go/bin/


دقت کنید تو مصاحبه اگر خواستید تو یه اسلایس بریزید ، اگر با توجه با ایندکس اون می خواید بریزید ، باید len , cap اون رو یکی بدید نت با توجه به ایندکس ، اون رو آپدیت کنید ولی اگه می خوایید به اون اپند کنید ، باید len ون رو صفر بزارید


		resultList := make([]int, lenList, lenList)
  		resultList[indexRes] = value


		resultList := make([]int, 0, lenList)
  		resultList =append(resultList, value)

    

#### Golang integer type
   می تونیم به چند روش، عدد درون تایپ اینت بریزیم ، توجه شود تمامی مقادیر ۱۵ هستند ، 

+ ۱۵___ دسیمال

+ اوکتال یا مبنای ۸ ___ 0o17 یا 017 

+ هگز یا مبنای ۱۶ ___ 0xF یا 0XF

+ باینری یا مبنای ۲ ___0b1111 

+ 1.23e2 == 123.0

+ rune== int32


  برای استفاده از utf8 هست و تمامی حروف زبان های جهان و کاراکتر ها در این هست
 

#### time.Time
این تایپ برای زمانه و اگر استرینگ با فرمت زیر بدیم ، می توینم unmarshal کنیم
```
type Event struct {
    Name      string
    Timestamp time.Time
}
...
    jsonStr := `{"Name":"Birthday","Timestamp":"2024-04-27T12:00:00Z"}`
    var event Event
    err := json.Unmarshal([]byte(jsonStr), &event)
```
#### signals

بهتره در صورتی که یه سرویس آپ می شه ، براش مدیریت سیگنال بنویسیم به این صورت که در صورت ورود SIGTERM یا SIGINT مدیریت کنیم ، به ترتیب در تمامی سرویس ها ، ابتدا باید این سیگنال ها رو بخونیم  سپس تمام تسک ها رو ببندیم یا کامکشن ها رو قطع کنیم و در نهایت EXIT(0) کنیم 

#### private git repository

اگر بخواهیم یه پکیبج درون گیت شرکت توسغه بدیم و تنها اونجا استفاده کنیم ، بادی مقدار زیر رو اکسپورت کنیم

یه نمونه آدرسی که درون گروه بک اند است
```
https://git.maani.app/maani/backend/kafka-wrapper
```
حالا می توان این رو به فایل /etc/environment اضافه کرد 
```
GOPRIVATE=*.example.app
```

همچنین اگر پروژه ذیل یه لایبری باشد مانند آدرس زیر ، باید درون gitconfig مقداری رو اضافه کنیم


nano ~/.gitconfig


```
[url "ssh://git@git.maani.app/"]
        insteadOf = https://git.maani.app/
```

#### 403 - 443 

یه سری وقتا یه پکیج نصب نمی شه ، شاید بشه با یه سری فلگ نصب کرد

`GOPROXY=https://proxy.golang.org,direct go get github.com/gin-gonic/gin`

`GODEBUG=netdns=go GOHOST=localhost GOINSECURE=nullprogram.com GONOSUMDB=nullprogram.com GO111MODULE=on go get -v github.com/gin-gonic/gin`

#### cobra vs multiple cmd folders

در صورتی که برنامه ی ما بیش از یک حالت ران شود یعنی همزمان هم rest api http  داره هم چند کرون جاب  ، می تونیم ۲ کار کنیم :

حالت اول این که چند cmd  و دایرکتوری داشته باشیم و درون هر کدوم ، main جدا ، خوبیش اینه خیلی مستقل هستن و نیاز به پکیج بیرونی نداره مانند کبرا همچنین به ازای هر کدوم ، یه داکر فایل باید داشته باشیم

حالت دوم استفاده از کبرا هست ، کلن یه بار بیلد میشه و یه فایل باینری داریم ولی در عوض می تونه با پارامتر هایی که میگیره ، 

### golang dont support

+ **type inheritance**

وراثت ساپورت نمی کنه ولی بجاش از `composition` همچنین اینترفیس ها می تونیم استفاده کنیم


+ **operator overloading**

تغییر علامت های محاسبات مانند موارد زیر در گو امکان پذیر نیست

`like +, -, *,`

+ **method overloading**

در یک پکیج متد هایی با نام مشابه اما ورودی های مختلف ، این هم در گولنگ امکان پذیر نیست بجاش می تونیم پسوند آخر اسمش بزاریم یا هر استراکت ، ریسیور خودش رو داشته باشه

+ **pointer arithmetic**

مثلن میشه عملیات روی پوینتر ها زد فرض کنید می گیم خونه ی بعدی یک پوینتر رو بده یا پلاس پلاس کن و در خونه بعدی یه چیز رو ذخیره کن و گولنگ اجازه نمی ده

+ **struct type in consts**

نمیشه یه استراکت ثابت ساخت برحلاف تایپ های بیسیک `like integers, floats, strings, and booleans` 


### tips


وجه شود در صورتی که فیلد های استراکت پرایویت باشد ، اگر اینستنسی از آن مارشال شود ، فیلد های پرایویت برگردانده نمی شود ، پس با دقت فیلد ها را پرایوت کن . ( کلا پرایوت جالب نیست )

redis - json ------- اینجوری می تونیم دیکد کنیم :



	result := new(interface{})
	json.Unmarshal([]byte(cacheResponse.(string)), result)
 
### golang conventional naming package

متاسفانه مرسومه که نام پکیج ها `lowercase` باشه و از `-` or `_` استفاده نکنیم

و این شد که من باید اسم پکیجم رو این بزارم :

`internalrequestservice`
