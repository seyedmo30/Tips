
زبان گو بیست استاتیک تتیپ هست و بهتره compile time  کداش اجرا شه ، اما با استفاده از reflect می شه ایجاد تایپ یا تعریف تایپ در runtime  اجرا بشه و چند تا بدی داره 

+ سرعت و پرفورمنس گو رو میگیره
+ تایپ سیف نیست و نمیشه روی تایپاش اطمینان کرد
+ خوانایی کد رو خیلی پایین میاره و پیچیده میکنه کد رو
+ توی تایپ های پریویت _ کمل کیس _ نمیتونه کار کنه

**اما مجبوریم جاهای زیر ازش استفاده کنیم**

+ مارشال و آنمارشال
+ جنریک پروگرم
+ جنریک دیتا استراکچر
+ خوانایی





## tips


+ زمانی که یه استک ساخته میشه و مقدارش کمه ، توسط runtime بزرگ میشه

+ اگر یه جایی باید یه اینترفیس رو ایمپلمنت می کردیم ، حال نیاز به فیلد های آن استراکت داشتیم ، مثال یه error custom  در نظر بگیرید
درون استراکت علاوه بر مسیج ارور میخواهیم فیلد status code هم بزاریم برای ارور، در این صورت تنها راه رسیدن به این فیلد، استفاده از تایپ اسرشن 
هست



### fmt

توجه شوداین پکیج از رفلکشن استفاده میکند ، به این معنی که برای همه ی ورودی هایی که بهش میدیم ، سوییچ کیس تایپش رو در میاره پس در جاهایی از کد که زیاد استفاده میشه ، بهتره جایگزینش رو استفاده کنیم
مثلن زمانی که می خوایم عدد رو تبدیل به اسرینگ کنیم بهتره ازstrconv استفاده شه



با استفاده از این قابلیت در runtime  می توانیم نوع و مقدار و تمامی اطلاعات یک interface{}  را مشخص کنیم


    v := reflect.ValueOf(data)   // مقدار data را دریافت می کنیم


    t := v.Type()   // نوع data را مشخص می کنیم


    if v.Kind() != reflect.Func   // اگر نوع data یک function باشد

    // با استفاده از kind می توانیم می توانیم تایپ مورد نظر رو در بیاریم

    if v.Type().NumIn() != 1  // در صورتی که تایپ داده ، فانکشن باشد ، تعداد ورودی را چک می کنیم

    if v.Type().NumOut() != 2 // در صورتی که تایپ داده ، فانکشن باشد ، تعداد خروجی را چک می کنیم

    reqParamType := v.Type().In(0) // در صورتی که تایپ داده ، فانکشن باشد ، نوع ورودی را مشخص می کنیم ، در اینجا اولین ورودی را مشخص می کنیم

    reqParam := reflect.New(reqParamType) // در اینجا با توجه به نوع که بالا مشخص شد ، یک اینستنس از آن نوع ایجاد می کنیم

    reqParamElem = reqParam.Elem() //  

    reqParamInterface = reqParam.Interface() //  


    result := v.Call(args)  // در صورتی که تایپ داده ، فانکشن باشد ، در اینجا این فانکشن را فراخوانی می کنیم

    result[0] // خروجی اولین فانکشن

### tip

+ **Type safety**

گاهی مشخص میشود ساختار ورودی و تایپ آن ، در این صورت  برنامه strict می کند داده رو به تایپ مورد نظر و از قبل از اجرای برنامه و زمان **compile time**  تایپ را مشخص می کند این باعث افزایش سرعت برنامه می شود ولی داینامیک بودن رو خیلی دشوار می کند



در صورتی که ورودی جیسون است اما ساختار آن را نمی دانیم ، به طور معمول آن را درون مپ می ریزیم:


    	var reqParamMap map[string]interface{}
	    err := json.Unmarshal([]byte(jsonData), &reqParamMap)







```


func GenericValidateAndUnmarshalFunc(data interface{}, jsonData string) {

	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Func {
		fmt.Println("Input is not a function")
		return
	}
	if v.Type().NumIn() != 2 {
		fmt.Println("Function should have exactly 1 parameters")
		return
	}

	if v.Type().NumOut() != 2 {
		fmt.Println("Function should have NumOut exactly 2 parameters")
		return
	}

	reqParamType := v.Type().In(1)

	reqParam := reflect.New(reqParamType)

	reqParamInterface := reqParam.Interface()
	json.Unmarshal([]byte(jsonData), reqParamInterface)

	reqParamElem := reqParam.Elem()

	args := make([]reflect.Value, v.Type().NumIn())

	args[1] = reqParamElem
	ctx := context.Background()

	valueCtx := reflect.TypeOf(ctx)

	args[0] = reflect.New(valueCtx)

	result := v.Call(args)

	for _, v := range result {

		fmt.Printf("Result: %v, Type: %v\n", v, v.Type())

	}
}


```