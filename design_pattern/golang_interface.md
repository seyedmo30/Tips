برای مشخص شدن تایپ اینترفیس :

type assertion  interface conversion

str, ok := interfaceVar.(string)

if msg, ok := rawMsg.(*pgproto3.CopyData); ok {

موارد مصرف اینترفیس :



# Embedding interfaces in structs

این روش برای وابسته کردن یک کلاس ( استراکت ) به اینترفیس است ، در حقیقت در کانستراکتور این کلاس باید شی ساخته شود که تمام رسیور های اینترفیس را ایمپلمنت کرده باشد 
```

    type controller struct {
    	dataStore  	 protocol.DataStore
     	downloader       protocol.Downloader
    }
    
    func NewController(
    	dataStore protocol.DataStore ,
     	downloader protocol.Downloader
    	) *controller {
    	return &controller{
     		downloader:       downloader,
    		dataStore:  dataStore
    	}
    }
```
در حقیقت مجبوریم شی بسازیم که dataStore را ایمپلمنت کرده باشد

```
	type DataStore interface {
		GetLastIdList() (uint32, error)
		List(ctx context.Context) chan types.SeedLink
		Migration(ctx context.Context) error
		Store(ctx context.Context, ggg chan types.ggg) error
	
	}
```
و در نهایت کلاینت باید با توجه به سیگنیچر های اینترفیس آن را پیاده سازی کند

 # Pass an Interface as an argument to a function
 

 #### ۱ - کانستراکتور یک استراکت ، نیاز به اینترفیس داشته باشد :
  یکی از رایج ترین راه ها برای دیکوپلینگ در روابط است . در این روش بجای ایمپورت کردن و استفاده در کد ، وابستگی را درکانستراکتور نشان می دهیم . در این حال کد قابلیت تست یونیت دارد و می توانیم وابستگی را موک کنیم

 
#### نکته منفی interface 
توجه شود اگر ما مستقیم استراکت را پاس دهیم ، هم دسترسی به رسیور داریم ، هم فیلد و گاهی استراکت نستد است و ما دسترسی به رسیور های نستد استراکت را هم داریم ، اما اگر از اینترفیس استفاده کنیم ، تنها از ریسیور های آن استراکت می توانیم استفاده کنیم 


توجه شود از این روش بیش از حد استفاده نکنیم ، زیرا آن استراکتی که اینترفیس را ایمپلمنت کرده ، شاید یکی از فیلد هاش اینترفیسی باشه و استراکتی امبد کرده که خیلی نیازه ، ولی وقتی از این روش استفاده کنیم ، تنها دسترسی به سیگنیچر های اینترفیس داریم . پس شورشو در نیار ( یعنی اگر هر لایه اینترفیس داشته باشد ، خواص کنتکستی یا دسترسی به پروپرتی های سوپر کلاس امکان پذیر نیست )

روش هایی مانند استفاده از سینگلتون و ایمپورت کردن pkg هم هست

#### ۲ - ورودی یک فانکشن باشد :
 این حالت برای گرفتن ورودی های خیلی جنریک استفاده می شه و مثالش ، توابعی که io/reader می گیرن . هر استراکتی که پروتو تایپ read  داشته باشه ، می تونه ورودی باشه

### return interface 

این حالت خیلی کم استفاده می شود و در داکیومنت ها پیشنهاد نمی شود . در حقیقت زمانی که بخواهیم متد های استراکت خروجی را محدود کنیم، از این استفاده می کنیم . مثلا شاید استراکتی که ریترن شده ۵ ریسیور داشته باشد ، اما چون اینترفیسی با ۳ متد ریترن شده ، کلاینت تنها از ۳ متد آن می تواند استفاده کند .

نکته: یکی از مشکلاتی که با کد حامد برخوردم ، اینه که اومده  پارامتر های ورودی رو اینترفیس گذاشته ، خروجی هارو اینتر فیس گذاشته ، بعضی اسراکت ها پابلیکن ولی چون از طریق اینترفیس پیاده سازی می شن ، تنها دسترسی به متد ها رو داریم ( به فیلد های پابلیک دسترسی نداریم ) ، همچنین تا یه لایه استراکت می تونیم از رسیور هاش استفاده کنیم ، (الان من نمی تونم از رسیویر های پایه ردیس یا پیتگرس استفاده کنم )

#### نکات :

بعد از کلی زور زدن برای پیاده سازی یک تابع جریک که ورودی ، یک تایپ استراکت بگیرد ، به این نتیجه رسیدم که خیلی کار درستی نیست ، می توان با استفاده از reflect.Type تنها تایپ را به عنوان پارامتر به تابع بدهیم ، اما مشکلات زیادی دارد 

خیلی راحت می شه تابع را به عنوان پارامتر به یک تابع پاس داد ، ولی متد خیلی سخته و درسر های زیادی داره ،
