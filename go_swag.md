swagger-echo



این مثال بر اساس اکو نوشته شده است



ابتدا باید برای تابع مین داکیومنت های زیر را اضافه کنیم

```


    import (
    "github.com/swaggo/echo-swagger"
    _ "aggrigation/docs"
    )

    // @title          title API
    // @version        1.0
    // @description    description.com
    // @query.collection.format multi
    func main() {
    	e.GET("", func(c echo.Context) error { return c.Redirect(http.StatusSeeOther, "/swagger/index.html") })

	    e.GET("/swagger/*", echoSwagger.WrapHandler)
    
    }
```
