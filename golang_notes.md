وجه شود در صورتی که فیلد های استراکت پرایویت باشد ، اگر اینستنسی از آن مارشال شود ، فیلد های پرایویت برگردانده نمی شود ، پس با دقت فیلد ها را پرایوت کن . ( کلا پرایوت جالب نیست )

redis - json ------- اینجوری می تونیم دیکد کنیم :



	result := new(interface{})
	json.Unmarshal([]byte(cacheResponse.(string)), result)
 

## export path

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

    
