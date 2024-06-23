### usecase

اگر بخواهیم کد رو بخش بندی کنیم ، به قسمتی از کد می گیم که فقط لاجیک بیزینسی در اون است 

این لایه باید جوری باشه که حتی کسی که زبان برنامه نویسی ما رو بلد نیست هم متوجه بشه

در این بخش ولیدیت ابتدایی داده یا یا بخش های خیلی سطح پایین برنامه نباید باشه تنها منطق و لاجیک اون کار

ورودی و خروجی باید یا entity باشه یا DTO اگر ورودی از استراکت به بایت تغییر کرد , usecase  ما تغییر نمی کنه چون فقط منطق اون تو هست 

در صورتی این لایه تغییر می کنه که منطق برنامه تغییر کنه  ـ نه به دیتا بیس ارتباط داره ، نه به پروتوکل های داده نه به کانورت ساختار ورودی و خروجی

همچنین نباید بخشی از منطق در آداپتور باشه بخشی در کنترلر ، منطق سراسر در یوزکیس هست

### repository pattern

یه design pattern هست که می گه کوییری باید از لاجیک جدا بشه و تمامی کوییری ها -crud - باید در یه لایه باشد و در لایه لاجیک تنها اونها کال بشن و با خواندن یوز کیس ، مشخص نباشه نوع دیتابیس یا متن کوییری ها .

ما همیشه از این دیزاین استفاده می کنیم بدون اینکه بفهمیم

### Clean Architecture Directory/Layers 
**Entities**: Represent the core business objects.
**Use Cases/Interactors**: Contain the application-specific business logic.
**Interfaces/Adapters**: Act as a bridge between the core application and the external systems.
**Frameworks/Drivers**: Include all external dependencies and infrastructure-related code.
