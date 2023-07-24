# byte

گاهی می خواهیم روی یک متغییر استرینگ تغییرات در دفعات زیادی اعمال کنیم ، در این صورت هر بار آدرس و مکان متغییر در مموری تغییر می کند زیرا تایپ استرینگ  ، ایمیوتبل هست . همچنین گاهی رشته ی ما آنقدر بزرگه که بهتر با اسلایس بایت اون رو مدیریت کرد .  
```
bt := []byte("salam")              لیترال اینیت

x := []int{10, 20, 30, 40, 50}     لیترال اینیت

byt := make([]byte, 0, 50)                     اینیت


```

# Reader

ریدر در مفهوم گولنگ ، یک اینترفیس است که باید پروتوتایپ زیر را پیاده کند . اصولا در کانکریت ، یک استراکتی وجود دارد که متدی به نام رید دارد که است فیلد های استراکت برمی دارد و بر روی اسلایس ورودی میریزد .

```
type Reader interface {
	Read(p []byte) (n int, err error)
}
```

# Writer

رایتر یک انترفیس است که متدی به نام رایت باید پیاده سازی کند که از اسلایس ورودی بخواند و به زیر لایه بریزد ، ( در استراکت خود اصولا فیلدی هست که بر روی آن بیفزاید ) 

```
type Writer interface {
	Write(p []byte) (n int, err error)
}
```

در مثال های زیر ریدر و رایتر پیاده سازی شدند

# Buffer

 کتاب خونه bytes ، برای دستکاری ( manipulation ) روی اسلایس بایت ها متد های زیادی داره . 

 در حقیقت بافر یک استراکت شامل اسلایس بایت ، طول اسلایس و تعداد متد برای راحتی کار با اسلایس فراهم کرده است ، استراکت بافر :


```
type Buffer struct {
    buf      []byte // contents are the bytes buf[off : len(buf)]
    off      int    // read at &buf[off], write at &buf[len(buf)]
    lastRead readOp // last read operation, so that Unread* can work correctly.
}

Bytes() []byte                                                     یک خروجی بایت می ده
Cap() int                                                ظرفیت
Grow(n int)                                         می تونیم به ظرفیت بیفزاییم
Len() int                                             طول 
Next(n int) []byte           اون تعداد که مشخص می کنیم از بافر ور می داره  در یه اسلایس می ریزه و خروجی می ده  
Read(p []byte) (n int, err error)      خیلی شبیه نکست هست ، از بافر می خونه و روی اسلایس ورودی می ریزه ، دقت کنید اضافه نمی کنه                 
ReadByte() (byte, error)
ReadBytes(delim byte) (line []byte, err error)
ReadFrom(r io.Reader) (n int64, err error)
ReadRune() (r rune, size int, err error)
ReadString(delim byte) (line string, err error)
Reset()               بافر رو خالی می کنه
String() string                            خروجی استرینگ می ده
Truncate(n int)                   یک عدد می گیره و به همون تعداد نگه می دار بقیه رو پاک می کنه
UnreadByte() error
UnreadRune() error
Write(p []byte) (n int, err error)            ورودی بایت گرفته و به سلایس می افزاید
WriteByte(c byte) error
WriteRune(r rune) (n int, err error)
WriteString(s string) (n int, err error)             ورودی استرینگ گرفته و به سلایس می افزاید
WriteTo(w io.Writer) (n int64, err error)                     
```





# NewReader (strings)

توجه شود اینترفیس ریدر و استراکت ریدر در پکیج استرینگ قاطی نشود :)



در بافر ، رید وجود دارد ، در زیر یک مثال از strings می بینیم که ریدر را پیاده سازی کرده . و متد رید دارد : 



ریدر یک استراکت شامل یک استرینگ ، طول رشته است که چند متد دارد ، در حقیقت شبیه با بافر است ، با این فرق که بافر درون خود اسلایس بایت دارد و ریدر استرینگ .

```
type Reader struct {
	s        string
	i        int64 // current reading index
	prevRune int   // index of previous rune; or < 0
}



Len() int                 طول استرینگ
Read(b []byte) (n int, err error)        از استرینگ می خونه و روی سلایس ورودی میریزه        
ReadAt(b []byte, off int64) (n int, err error)
ReadByte() (byte, error)
ReadRune() (ch rune, size int, err error)
Reset(s string)
Seek(offset int64, whence int) (int64, error)
Size() int64
UnreadByte() error
UnreadRune() error
WriteTo(w io.Writer) (n int64, err error)

```


# example 

 
```
   package main
   import (
   	"fmt"
   )
   //importing the bytes package so that buffer can be used
   import (
   	"bytes"
   )
   func main() {
   //Creating buffer variable to hold and manage the string data
   	var strBuffer bytes.Buffer          تعریف متغیر
    strBuffer.Grow(64)           به صورت دیفالت ظرفیت بافر ۶۴ است ، بعد اون ۸۰ ولی می تونیم با این همون اول ظرفیتشو مشخص کنیم
    fmt.Fprintf(&strBuffer, "It is number: %d, This is a string: %v\n", 10, "Bridge")           افزودن به بافر
    strBuffer.Write([]byte("Hello "))         افزودن به بافر
   	strBuffer.WriteString("Ranjan")              افزودن به بافر
   	strBuffer.WriteString("Kumar")
    bugstr := strBuffer.String()                  گرفتن خروجی استرینگ
   	fmt.Println("The string buffer output is",bugstr)

    strBuffer.WriteTo(os.Stdout)            مثل پرینت
   }
```
