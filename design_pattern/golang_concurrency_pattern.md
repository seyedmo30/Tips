

## پترن جنراتور

* ساخت چنل درون جنراتور است - make
* جنراتور حداقل یک چنل خروجی دارد - return chan
* درون جنراتور گوروتین وجود دارد - go
* گلوز شدن چنل ، درون جنراتور است - close


همچنین هر چنل می توان استراکتی باشد که در خود ارور هم دارد ، و یکی از ورودی های جنراتور ، فلگ است و جنراتور در خود سلکت دارد ، مورد اول فلگ را بررسی می کند، در صورتی که دان شود دستور مربوطه اجرا می شود و در صورتی که نباشد روال عادی طی می شود



**بهتره ownership  یه چنل ، همونی باشه که توش write میکنه ، همچنین خودش اون رو close کنه.**

```go

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

به عبارت دیگر چندین گوروتین کار می کنند و نتیجه را از طریق یک چنل به یک گوروتین می دهند


چندین پردازش سند می کنن تو یک چنل

احتمالا قبل از این که چند گوروتین بخوان درون یک چنل بریزن ، ابتدا هر گوروتین در چنل های مجزا به merge method  میدن ، و این فانکشن درون یک چنل میریزه



## پترن FanOut

یک چنل گرفته و در چند چنل می ریزد . برای دیستریبیوت کردن بین چند ورکر مفید است .

به بیان دیگر یک گوروتین کار می کند و نتیجه را به چندین گوروتین می دهد



چندین گوروتین از یک چنل ریسیو می کنن

احتمالن این بخش شبیه worker pool هست ، تفاوتی که با ورکر پول دیدم ، این بود که اون یه حلقه داره ، گوروتین ها از یک چنل میخونن و به یک چنل میریزن ، اما در fan out هر گوروتین خروجی رو توی چنل مخصوص خودش میریزه

![alt text](https://github.com/seyedmo30/Tips/blob/main/static/iii3.png)

![alt text](https://github.com/seyedmo30/Tips/blob/main/static/iii4.png)

## پترن PipeLine

 خروجی یک جنراتور ، ورودی جنراتور بعدی هست :
 + همه ی فانکشن ها ، ورودی و خروجی اش چنل است ، به جز اولی و آخری
 + فانکشن میین یا کنترلر وظیفه دارد خروجی یک جنراتور را به جنراتور بعدی بدهد .


نکته ای که هست ، بهتره گوروتین های آپ استریم در پایان چنل رو کلوز کنن ، و گوروتین های داون استریم یه حلقه با range بخونن و در صورت پایان از حلقه خارج شن

![alt text](https://github.com/seyedmo30/Tips/blob/main/static/iii2.png)


## پترن Worker Pool

 تعدادی ورکر داریم که به صورت کانکارنت در گوروتین های مختلف کار می کنند :
 + تعداد ورکر ها را ابتدا مشخص می کنیم
 + یک حلقه به تعداد ورکر ها ایجاد کرده و در هر بار ، گوروتینی ایجاد می شود

```go
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
```

## پترن Queuing - concurrency-limiting pattern

 گاهی نیاز است پردازش همزمان  را محدود کنیم تا منابع پر نشود  :
 ```go
 
 		queue := make(chan struct{},limit)
 ```
 + هرظ اجرای پردازش، باید داشتن ظرفیت در بافر این چنل باشد
 + هر گوروتین که اجرا می شود به این چنل ، یک استراکت امپتی افزوده می شود و اگر تکمیل شد از آن برداشته می شود


## پترن For-Select-Done

گاهی منتظر چند چنل هستیم و هر کدام زود تر پاسخ دهد را استفاده می کنیم

گاهی نیاز است برای دریافت داده از چنل ، ددلاین بزاریم تا در صورت رخ دادن مشکل ، منتظر نمانیم :
+ فانکشن حداقل ۲ ورودی میگیرد، چنل ورودی و چنل دان ، می توانیم جای دان ، کانتکست استفاده کنیم
‍
```go
select {
case msg := <-ch1:
	fmt.Println("Received", msg)

case <-time.After(time.Second):
    fmt.Println("Timeout")

case  <-Done:

}

```
### drop pattern
یه روش برای زمانی که حفظ داده در اولویت نیست و می توان داده ی اضافی را دور ریخت و به این صورته که از select استفاده می کنیم و default و درصورتی که نتوانیم از chann بخونیم یا بنویسیم ، دیفالت اجرا شده و رد می کنی 

```go

func producer(queue chan<- int, id int) {
    for i := 0; i < 10; i++ {
        select {
        case queue <- i:
            fmt.Printf("Producer %d: Produced %d\n", id, i)
        default:
            fmt.Printf("Producer %d: Dropped %d\n", id, i)
        }
        time.Sleep(100 * time.Millisecond)
    }
}
```
9. Publish-Subscribe Pattern
Description: Allows multiple consumers (subscribers) to listen to messages published by producers.
Use Case: Event-driven architectures, notifications.

10. Timeout Pattern
Description: Uses a timer or timeout channel to handle operations that exceed a certain duration.
Use Case: Avoid blocking indefinitely in network calls or resource access.

11. Retry Pattern
Description: Retries a failed operation with exponential backoff or a fixed number of attempts.
Use Case: Fault-tolerant systems.

12. Resource Pool Pattern
Description: Maintains a pool of reusable resources, such as database connections or goroutines.
Use Case: Efficient resource management.
