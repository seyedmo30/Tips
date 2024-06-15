# reflect golang

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