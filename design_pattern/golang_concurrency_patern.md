

## پترن جنراتور

* ساخت چنل درون جنراتور است - make
* جنراتور حداقل یک چنل خروجی دارد - return chan
* درون جنراتور گوروتین وجود دارد - go
* گلوز شدن چنل ، درون جنراتور است - close


همچنین هر چنل می توان استراکتی باشد که در خود ارور هم دارد ، و یکی از ورودی های جنراتور ، فلگ است و جنراتور در خود سلکت دارد ، مورد اول فلگ را بررسی می کند، در صورتی که دان شود دستور مربوطه اجرا می شود و در صورتی که نباشد روال عادی طی می شود


```

type Result struct {
	data int
	err  error
}



func handler() {
	input := []int{1, 2, 3, 4, 5, 6}

	// Explicit cancellation
	doneCh := make(chan struct{})
	defer close(doneCh)

	inputCh := generator(doneCh, input)

	for data := range inputCh {
		if data == 1 {
			return
		}
	}
}



func generator(doneCh chan struct{}, input []int) chan int {
	inputCh := make(chan int)

	go func() {
		defer close(inputCh)

		for _, data := range input {
			select {
			case <-doneCh:
				return
			case inputCh <- data:
			}
		}
	}()

	return inputCh
}



func consumer(inputCh chan int, resultCh chan Result) {
	defer close(resultCh)

	for data := range inputCh {
		resp, err := callDatabase(data)

		result := Result{
			data: resp,
			err:  err,
		}

		resultCh <- result
	}
}


```


## پترن FanIn

نوعی جنراتور است که چندین چنل می گیرد و تجمیع کرده ، یک چنل خروجی می دهد .



## پترن FanOut

یک چنل گرفته و در چند چنل می ریزد . برای دیستریبیوت کردن بین چند ورکر مفید است .


## پترن PipeLine

 خروجی یک جنراتور ، ورودی جنراتور بعدی هست :
 + همه ی فانکشن ها ، ورودی و خروجی اش چنل است ، به جز اولی و آخری
 + فانکشن میین یا کنترلر وظیفه دارد خروجی یک جنراتور را به جنراتور بعدی بدهد .




## پترن Worker Pool

 تعدادی ورکر داریم که به صورت کانکارنت در گوروتین های مختلف کار می کنند :
 + تعدتد ورکر ها را ابتدا مشخص می کنیم
 + یک حلقه به تعداد ورکر ها ایجاد کرده و در هر بار ، گوروتینی ایجاد می شود

   

## پترن Queuing

 گاهی نیاز است پردازش همزمان  را محدود کنیم تا منابع پر نشود  :
 
 		queue := make(chan struct{},limit)
 + هرظ اجرای پردازش، باید داشتن ظرفیت در بافر این چنل باشد
 + هر گوروتین که اجرا می شود به این چنل ، یک استراکت امپتی افزوده می شود و اگر تکمیل شد از آن برداشته می شود


## پترن For-Select-Done

گاهی نیاز است برای دریافت داده از چنل ، ددلاین بزاریم تا در صورت رخ دادن مشکل ، منتظر نمانیم :
+ فانکشن حداقل ۲ ورودی میگیرد، چنل ورودی و چنل دان ، می توانیم جای دان ، کانتکست استفاده کنیم

### drop pattern
یه روش برای زمانی که حفظ داده در اولویت نیست و می توان داده ی اضافی را دور ریخت و به این صورته که از select استفاده می کنیم و default و درصورتی که نتوانیم از chann بخونیم یا بنویسیم ، دیفالت اجرا شده و رد می کنی 


