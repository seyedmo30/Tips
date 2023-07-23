# bytes 

گاهی می خواهیم روی یک متغییر استرینگ تغییرات در دفعات زیادی اعمال کنیم ، در این صورت هر بار آدرس و مکان متغییر در مموری تغییر می کند زیرا تایپ استرینگ  ، ایمیوتبل هست . همچنین گاهی رشته ی ما آنقدر بزرگه که بهتر با اسلایس بایت اون رو مدیریت کرد .  

 کتاب خونه bytes ، برای دستکاری ( manipulation ) روی اسلایس بایت ها متد های زیادی داره . 

 در حقیقت بافر یک استراکت است که اسلایس باید  و طول آن را در خود ذخیره می کند و تعداد متد برای راحتی کار با اسلایس فراهم کرده است ، استراکت بافر :

type Buffer struct {
    buf      []byte // contents are the bytes buf[off : len(buf)]
    off      int    // read at &buf[off], write at &buf[len(buf)]
    lastRead readOp // last read operation, so that Unread* can work correctly.
}

func (*bytes.Buffer).Bytes() []byte
func (*bytes.Buffer).Cap() int
func (*bytes.Buffer).Grow(n int)
func (*bytes.Buffer).Len() int
func (*bytes.Buffer).Next(n int) []byte
func (*bytes.Buffer).Read(p []byte) (n int, err error)
func (*bytes.Buffer).ReadByte() (byte, error)
func (*bytes.Buffer).ReadBytes(delim byte) (line []byte, err error)
func (*bytes.Buffer).ReadFrom(r io.Reader) (n int64, err error)
func (*bytes.Buffer).ReadRune() (r rune, size int, err error)
func (*bytes.Buffer).ReadString(delim byte) (line string, err error)
func (*bytes.Buffer).Reset()
func (*bytes.Buffer).String() string
func (*bytes.Buffer).Truncate(n int)
func (*bytes.Buffer).UnreadByte() error
func (*bytes.Buffer).UnreadRune() error
func (*bytes.Buffer).Write(p []byte) (n int, err error)
func (*bytes.Buffer).WriteByte(c byte) error
func (*bytes.Buffer).WriteRune(r rune) (n int, err error)
func (*bytes.Buffer).WriteString(s string) (n int, err error)
func (*bytes.Buffer).WriteTo(w io.Writer) (n int64, err error)




 
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
    fmt.Fprintf(&strBuffer, "It is number: %d, This is a string: %v\n", 10, "Bridge")           افزودن به بافر
    strBuffer.Write([]byte("Hello "))         افزودن به بافر
   	strBuffer.WriteString("Ranjan")              افزودن به بافر
   	strBuffer.WriteString("Kumar")
    bugstr := strBuffer.String()                  گرفتن خروجی استرینگ
   	fmt.Println("The string buffer output is",bugstr)

    strBuffer.WriteTo(os.Stdout)            مثل پرینت
   }
```
