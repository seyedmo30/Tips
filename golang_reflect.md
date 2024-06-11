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






