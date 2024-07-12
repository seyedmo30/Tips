
#### Dbms _ سیستم مدیریت دیتابیس

به نرم افزار هایی که برای ذخیره دیتا هستند میگن ، قبلا مجبور بودن توی فایل ذخیره کنن ، نه ایندکس بوده نه سرچو نه .... , بدیش هم اینه باید sql بلد باشید همچنین داده ها یه متمرکزن و اگر خراب شه ، تمام داده ها از دسترس خارج میشه

####  Dba _ ادمین دیتابیس

مسعول حفظ و مدل سازی و امنیت .... دیتابیس ها

#### (ORDBMS)  object model
علاوه بر ویژگی های دیتابیس رابطه ای ، تعداد کمی 
خواص شی گرایی هم داره:
+ User-Defined Data Types -  ـ می توان تایپ دلخواه را ساخت

+ Inheritance - می توان یک تیبل از تیبل دیگر ارث بری کرد ، مثلا تیبل دانشجو از تیبل یوزر ارث بری کرده

+ Polymorphism - می توان توابعی نوشت که هر تایپی را میگیرد یا خروجی دهد، anyelement

##  Transactions


### اسید

+ اتمیک

میگه یه استوریج باید یا همه ی تراکنش ها رخ بده یا  هیچکدوم تغییر نکنه ، به عبارت دیگر اتومیسیتی یعنی نمی تونیم یه چیز رو بشکونیم مثل اتم ، یک واحده

+ کانسیستنسی _ ثبات

باید استوریج داده را از حالت ولید تغییر بده و به حالت ولید بیاره ، یعنی با تغییر داده ، خراب نشه ، یا 4ناخوانا نشه ، همچنین شروط و محدودیت های دیتا بیس قبل از تراکنش و بعد از آن نقض نشود

+ ایزوله

اگر همزمان چندین داده تغییر کرد ، ترتیب آنها رعایت شه و ریس کاندیشن رخ نده ، همون لاک کردنه، این با mvcc خیلی ارتباط داره

+ دیوریبیلیتی

تضمین میکنه با خرابی سیستم عامل یا قطعی برق ، داده نپره



 
تو این لول حتی اگر داخل ترنساکشن ، یه ترنساکشن دیگه داده ها رو آپدیت کرد ، ما فرض میکنیم داده های قدیمی هنوز هستن 

توی ترنساکشن ها می تونیم در انتهای در خواست یه در اصطلاح برای خود ترنساکشن ها استفاده کنیم
Select * from balance where id =4 for update 
این به این معناست کهکه حتی اگر یه ترنساکشن دیگه خواست موجودی بگیره ، واسته تا این ترنساکشن کامیت بشه


### execute

شامل دستورات اصلی مانند begin , commit , rollback میشود

+ Savepoint

درصورت رول بک ، میتوانیم بگیم از کجای ترنساکشن دوباره شروع کند

        SAVEPOINT my_savepoint; 

        ROLLBACK TO my_savepoint;

### Concurrency Control _ TRANSACTION ISOLATION LEVEL
می توان سطح مشاهده داده همزمان چند ترنس اکشن را مدیریت کرد 

+  READ UNCOMMITTED 

 در این حالت ، میتوان تغییرات قبل از کامیت را مشاهده کرد

در این لول یه ترنساکشن حتی میتونه زمانی که یه ترنس اکشن دیگه کامیت نشده هم اون رو بخونه 
خیلی بی خود

        SET TRANSACTION ISOLATION LEVEL READ UNCOMMITTED 



+  READ COMMITTED 

 داده راتنها بعد از کامیت میتوان دید

در این حالت تا زمانی که ترنساکشن کامیت نشده ، داده های قدیمی رو میخونه


        SET TRANSACTION ISOLATION LEVEL READ COMMITTED 

+ repeatable read

تو این لول حتی اگر داخل ترنساکشن ، یه ترنساکشن دیگه داده ها رو آپدیت کرد ، ما فرض میکنیم داده های قدیمی هنوز هستن 

        SET TRANSACTION ISOLATION LEVEL repeatable read


#### Locking transaction

این قابلیت جلوگیری میکند از تغییرات همزمان و ترناکشن ها بر روی چند شارد ، انواع مود ها : 

FOR UPDATE, FOR NO KEY UPDATE, FOR SHARE,  FOR KEY SHARE

توی ترنساکشن ها می تونیم در انتهای در خواست یه در اصطلاح برای خود ترنساکشن ها استفاده کنیم

Select * from balance where id =4 for update 

این به این معناست کهکه حتی اگر یه ترنساکشن دیگه خواست موجودی بگیره ، واسته تا این ترنساکشن کامیت بشه


        BEGIN; 
        SELECT * FROM my_table WHERE id = 1 FOR UPDATE;
        COMMIT;

#### Schema
برای دسته بندی کردن  و مرتب سازی حجم بالای تیبل ها و فانکشن ها استفاده میشود، همچنین برای سطح دسترسی به یوزر ها می توان از اسکیما استفاده کرد

#### Domain
علاوه بر دیتا تایپ هایی که پیش فرض داره ، میتونیم تایپ های خود با شرط های خودمون رو بسازیم


###  Mvcc _multi version concurence control

یکی از ویژگی های دیتابیس است و این قابلیت رو میده که به صورت هم زمان چند درخواست خواندن از دیتابیس رو هندل میکنه و بدون وقفه یا تداخل ، به این صورت که ابتدا یک اسنپشات میگیرد و ترنزاکشن ها رو روی اون اسنپ شات هم اعمال میکنه و از روی اسنپ شات هم میخونه
، از هر سطر، ورژن های مختف می سازه و هر ترنزاکشن روی یک ورژن اعمال میشه و به بقیه کار نداره ،همچنین یک گاربج کالکتور داره که ورژن های قدیمی سطر ها رو پاک میکنه ، و هر ترنس اکشن یه یونیک آیدی داره که بهش (TXID) میگن

مثال_ اگر ترنس اکشنی را باز کنیم و در مدت طولانی آن را کامیت نکنیم ، در حقیقت از داده ها اسنپ شات گرفتیم و بعد از کامیت ، اسنپ شات پاک میشود

+  مزایا 

 باعث میشه بدون لاک کردن در موقع تغییر دیتا بیس ،چندین یوزر دیتا رو تغییر بدن

+ معایب 

 به همون اندازه که اسنپ شات میگیره ، مموری میخوره

###  Query proccessing

+ ۱ پارسینگ 

 وظیفه ی پارس کردن کوییری استرینگ رو داره و ابتدا گرامر رو چک می کنه ، در صورت درست بودن آن را به یک درخت تبدیل می کنه و 
اجزای آن رو واکشی میکنه

+   ترسفورمیشن rewriter 
 
در این مرحله با توجه به rules عملیات صورت میگیرد ، مثلا اگر دستور آپدیت باشد ، ابندا سطر جدید ساخته میشه و سطر قدیمی هیدن میشه و در صورت ترنساکشن ، وکیوم میشن

+ Planner 

وظیفه ی این قسمت ، انتخاب بهرین و سریع ترین راه برای اجرای کوییری است ، راه های مختلف را در نظر میگیرد و با توجه به cpu , io  و هزینه یکی را انتخاب میکند
همچنین یه بخش آماری برای براورد هزینه دارد که اطلاعاتی مانند سایز تیبلو .. دارد که آمار دیتابیس را در آن ذخیره میکند

+ Executor

در حقیقت این کارای اصلی  crud رو انجام میده 


## VACUUM مطالعه شود analize . Explain


### Psql
+ Interactive command line client - 
برای اتصال به پستگرس

+  آیا psycopgy یا pgx از psql استفاده  می کنن؟؟
خیر ، از postgresql protocol  استفاده میکند

+ pg_ctl , pg_ctlcluster - 
با این ابزار می توان سرویس پستگرس را مدیریت کرد مانند استاپ یا استارت کردن ، اما متداول نیست و همه با systemctl کار میکنن

+ pg_ctlcluster -  برای دستورات کامندی برای کلاستر ها می باشد

### concept

+ Data Definition Language_ ddl

تغییر در ساختار دیتا بیس مانند create , drop  در جدول یا دیتا بیس یا constraint ها و مثال عملی ، همون کوییری هستش ، مثلا یه کویری که ران می کنیم ، یه ddl هست

+ Dml _ Data Manipulation Language

Crud بر روی سطر های جداول است
دستورات و کوییری ها ، ۴نوع: insert, delete, update, select

+ COPY TABLE

می توان با این روش ، داده را مثلا از روی فایل با فرمت csv خواند و بر روی تیبل ریخت ، همچنین برعکس هم میشه

+ Varchar , char

اگر سایز varchar را بزرگ درنظر گرفت اما سایز داده ها کوچک بود ، به اندازه سایز داده ، حافظه اشغال میشود

#### Uncommon types 

پستگرس علاوه بر تایپ هایی که میشناسیم ، دیتا 
تایپ های دیگری هم داره:

+ Network Address Types

Cidr , Inet آی پی ورژن ۴ یا ۶ و سابنت مسک را ذخیره میکنه
 Macaddr مک آدرس سخت افزار ها را ذخیره میکنه

+ Geometric Types

Point طول و عرض رو ذخیره میکنه ، مختصات نقطه
Line  دو تا نقطه رو ذخیره میکنه ، نقطه شروع و پایان
Polygon چند نقطه رو ذخیره میکنه

####  Common Table Expressions _ CTE
می توان یک کوییری موقت ایجاد کرد و بر روی آن تیبل موقت ، کوییری زد، در حقیقت اگر بخواهیم از نتیجه ی یک کوییری مانند سلکت ایتفاده کنیم ، و یا نتایج یک تیبل سلکت را به صورت موقت ذخیره کنیم

+ Recursive CTE

می توان تعداد زیاد جدول موقت ایجاد کرد و از جدول ایجاد شده، استفاده کرد و دوباره جدول جدید ساخت

####  subquery
سابکوییری ، کوییری تو در تو درون یک کوییری است. همچنین می توان، برای مقدار دستورات زیر استفاده کرد:
SELECT, FROM, WHERE, HAVING

 انواع:
+ Scalar Subqueries 

 سابکوییری یک مقدار بر میگردونه ، مانند avg , sum ، می تونه یا عدد باشه یا استرینگ  ) درکل یه کالمن(

        WHERE employees.salary > (SELECT AVG(salary)

+ Row Subqueries 

یه سطر با چندین کالمن بر میگردونه

        WHERE (order_id, total) = (SELECT order_id, total FROM order)

+ Column Subqueries 

چندین سطر فقط یک کالمن

بیشتر برای موارد استفاده میشهIN, ALL, ANY.  


        WHERE price IN (SELECT MAX(price) FROM 

+ Table Subqueries 

 یک تیبل بیشتر برای استفاده از from

        FROM (SELECT customer_id, SUM(total) as

####  مقایسه cte vs subqueru
Cte ها یک جدول موقت هستند در حالی که subquery یک سلکت است و می توان خروجی یک سلکت ، عدد ، استرینگ ، سطر ، ستون یا تیبل باشد 
همچنین 


### Write ahead logging
مزایای استفاده از wal 
+ برای اینکه داده ها را کامل داشته باشیم تا در صورتی که crash کرد یا از دیسک پرید داشته باشیم
+ راهی ساده برای لاگ خوانی ترن اکشن ها
+ می توانیم داده های قدیمی که تغییر کردند را نگهداری کنیم
+ با استفاده از این فایل ها می توانیم در صورت کرش ، دادخ را بخوانیم و برگردانیم
+ می توان استریم رپلیکیشن انجام داد

#### Lateral join

اگر بخواهیم جویین بین یک تیبل و یک ساب کوییری بزنیم ، این مفهوم استفاده میکنیم