
### ارتباط سرویس ها با هم 

**Synchronous Communication** 

+ **HTTP (RESTful APIs)** خیلی راحت و سریعه ، خیلی جاها استفاده می کنه ، استیت لس هست و خیلی خانا هست و ولی ایرادش شاید تاخیر و سنگینیش باشه

+ **gRPC** سرعت بالا ،  استفاده از http2  ،  و استفاده برای زمانی که سرعت بالا و تاخیر کم میخواییم ، استریم هم میتونیم باهاش کنیم ، خوانایی و پیاده سازی و نگه داریش سخته

**Asynchronous Communication**

+ **Messaging Systems (Message Brokers)** مانند RabbitMQ, Apache Kafka ، پابلیشر ها پیام ها رو به بروکر ها میدن و میشه یک یا چند کانسور استفاده کنن ، سختیش مدیریت و مانیتور کردن داده و تست هست


+ **Event-Driven Communication** از کافکا استفاده میشه ، پابلیشر ایونت ها رو میزاره و چندین کانسومر ورمیدارن ، در این حالت ایده آل زمانیه که چندین سرویس کانوم کنن و تفاوتش با مورد بالایی ، اینه که شاید چندین سرویس منتظر باشه ، بدیش هم اینه که دیباگ و تست و مانیتور سخته و آدنپوتسی و اوردرینگ رو باید دستی توش هندل کنیم

**Hybrid Communication**

+ **WebSockets** چنل های ارتباطی ایجاد میکنه ، و تنها یه tcp connection  رو برای تعامل می خواد  ، برای آپدیت ریل تایم و لایو داشبود و چت خوبه ، بدیش اینه که نیاز به زیر ساخت مخصوص داره

+ **Remote Procedure Call (RPC)** اجازه میده یه فانکشن توی یه سرویس ، یه فانکشن توی یه سرویس دیگه رو کال کنه




### Event-driven architecture

یک الگوی طراحی است که اجزای آن رویداد محور ارتباط با هم دارن و این ارتباط شامل generation, detection, and consumption of events است

در این الگو ، کامپوننت ها لوس کاپل هستند و اسکیل و انعطاف پذیری بالا می باشد


**Key Features** : Event producers, event consumers, event channels, event processing

**Challenges سختی هایی که داره** Complexity - Consistency - Latency - Debugging and Monitoring - Security

# interview

اگر request  باید real-time  پاسخ داده بشه قطعا باید از sync ها مثل grpc , rest استفاده کنیم اما گاهی پاسخ شاید long-running-prcess  باشه در این صورت روش های زیر هستند

### زمانی که نیاز به ارتباط سینک داریم ، ایا منطقی است از مسیج بروکر ها یا صف استفاده کرد؟

منطقی این است که در سیستم ارتباطاتی که ریکویست ریسپانسی هستند ، از سینک استفاده کنیم like grpc  اما در حالاتی میشه از کافکا استفاده کرد

+ زمانی که می خوایم لوس کاپل کنیم سرویس ها رو 

+ زمانی که حجم داده خیلی زیاده مثلن از باکت استفاده کنیم 
+ زمانی که معماری های ایونت دیریون مد نظر داریم
+ میشه یه جورایی به جای لود بالانسر هم استفاده کرد به این معنی که اگر چندین پاد یه تاپیک رو کانسوم کنند ، یه جورایی این نقش رو بازی میکنه

+ اما میشهگفت زمانی که سرعت پاسخ مهم است ، واقعن هیچ توجیهی ندارد



## client-server (request-response)
فرض کنیم پردازش های طولانی داریم - long running process -  در این حالت زمانی که حجم دخواست ه کم باشه ، یا بخواهیم real-time پاسخ بدیم یا بخواهیم پیجیدگی سیستم رو کم کنیم می تونیم از این روش استفاده کنیم

باید توی سرور یه صف بزاریم و درخواست رو از کلاینت بگیریم ، اما کلاینت باید polling  کنه و وضعیت درخواستش رو ببینه

همچنین برای حل idempotency and tracing  باید کلاینت یه request_id یونیک جنریت کنه و می تونه توی x-request header  بزاره که gateway سرور اون رو کش کنه و از idempotency مطمعن باشه و همچنین اگه پاسخ کلاینت تایم اوت خورد ، راحت retry  کنه و ref_number  رو دریافت کنه 

حالا کلاینت باید poll  کنه تا آخرین وضعیت رو ببینه همچنین سرور می تونه کمی به کلاینت حال بده و server-send-event(sse) بزاره و نوتیف به کلاینت بده ، اما لاجیک ریترای پیچیده میشه - توجه شود اول و آخر باید کلاینت polling  داشته باشه ، اگه بخواهیم به جاش callback - notif  و یا webhook  رو بیاریم کلی لاجیک اضافی سمت  سرور میاد که همون قدر ریسورس می خواد مثلا اگه کلاینت داون باشه تا کی باید ریترای کنه؟

مزایاش راحتی و سادگی ، است اما بدیش اینه قابل اسکیل نیست ، polling به شدت سر بازه و اگه کلاینتا زیاد شن یا پاسخ دیر آماده شه خیلی سربار ریسورس داریم و شاید کلاینت دیر polling کنه و دیر نتیجه رو ببینه و tight coupling خفن داره




###  تفاوت کیو و پاب ساب

در کیو ، تنها یه صفه ، می توان بهصورت غیر همزمان مسیج گذاشته شود و بعد از خواندن از صف ، از صف حذف میشه


اما در پاب ساب ، حذف نمیشه بلکه در مدت معلوم می مونه شاید چندین سابس کراب در زمان های مختلف بخونن

با استفاده از گروپ آی دی ، میشه نیاز های صف در کافکا بر آورده کرد

مسیج ها به صورت راند رابین در پارتیشن ها میرن و به صورت پارالل به کانسیومر ها می رن

توی ربیت ، بروکر می ره به کانسیومر میگه داده ی جدید داری ، اما در کافکا ، خود کانسیومر باید بره از بروکر بپرسه پیام جدید چی داری



تحقیق بشه در مورد ack و کامیت در کافکا

اینگار  down time  داریم

تحقیق در مورد 0 , 1 all

در نهایت در کافکا بر اساس زمان مسیج حذف میشه اما در ربیت بعد از اینکه به کانسیومر ها داده شد

## message queues or event-driven architecture

این روش تمام معایب و مزایای روش بالایی رو بر عکس داره

راحت میشه اسکیل کرد - نیازی به polling نیست - در نتیجه ریسورس پولینگ رو نداره - همچنین decoupling پیاده شده - همچنین چون پیام های توی صف ها عمر طولانی دارن ، در صورت پایین اومدن سرور یا کلاینت مقاوم ترن

اما به شدت پیچیده هستن به طوری که تمام retry - idempotency - faolt telorance باید لاجیکش پیاده بشه - به نسبت روش های سینک سرعتش کمتره - و دیباگ خیلی سخت تره

**implement**

اما حالا بیایم پترن هایی رو ببینیم که retry - idempotency- tracing - fault telorance رو ببینیم


#### **Kafka Retry Pattern**

+ Consumer-Level Retries

کانسومر بعد از دریافت پیام ، باید اون رو پراسس کنه و اگه نتونست حلقه بزاره و دوباره تلاش کنه 

+ Dead Letter Queue (DLQ)

اگه نتونست پردازش کنه ،  اون رو به صفی می فرسته که بشه بعدا دیباگ کرد و بشه از حلقه تکرار خارج شد

+ Exponential Backoff with Jitter

میشه به جای این که پشت سر هم تلاش کرد ، تو بازه هایی که هر بار طولانی تر میشه تلاش کرد

#### **Kafka Idempotency**

اگر استراتجی **at-least-once delivery** رو داشته باشیم باید مطمعن شویم که تکراری نباشه

+ unique request ID

می تونیم از شناسه درخواست استفاده کنیم و اون ها رو توی **Redis** بریزیم و چک کنیم

+ Transactional Consumer Approach (exactly-once semantics)


اینگار قابلیتی هست که من بلد نیستم و میشه از قابلیت کافکا به نام **transactional.id** استفاده کرد

و زمانی که اطمینان داریم کامیت کنیم

همچنین از**upsert** دیتابیسی برای جلوگیری از تکراری بودن استفاده کنیم

+ استفاده از version  پیام

به این صورت که هر مسیج علاوه برای request-id برای تکراری نبودن ، ورژن هم داشته باشه که اگر ورژن قدیمی تر اومد ، اون رو آپدیت کنیم


#### **Tracing in Kafka**

تقریبا شبیه به حالت client-server sync  هست و نکته خاصی نداره

#### **fault-tolerant strategy**

در صورتی که کافکا کرش کند اتفاقاتی که احتمالا رخ بدهد

+ در صورتی که رپلیکا ها هم کرش  کنند احتمالا **پیامها** از بین بروند

+ احتمالا **آفست** کانسیومر اشتباه شود

+ **پرودیوسر** نتونه پیام بزاره

+ و در نهایت **بک لاگ** پر شه و با حجم زیادی پیام در صف مواجه شویم

**راه حل ها**

+ از رپلیکا استفاده کنیم

+ حداقل از ۳ **Kafka brokers** استفاده کنیم

+ پرودیسر **enable.idempotence = true** رو فعال کنه

+ کانسیومر هم **enable.auto.commit = false** رو فعال کنه 

+ آخرین راه هم اینه مسیج ها رو قبل پرودیوس کش کنیم توی **MirrorMaker**  مثل ردیس


### حالا سناریو سیستم دیزاین می چینیم

+ کلاینت یه مسیج با req_id آماده می کنه . توجه شود اینجا نیازی به retry **نیست**

+ سپس داخل تاپیک **orders-topic** پرودیوس می کنه و مطمعن میشه که سالم پرودیوس شده

+ سرور از تاپیک **orders-topic** کانسوم می کنه 

+ چک می کنه ببینه توی **ردیس** تکراری نباشه

+ تلاش می کنه که پردازش کنه و در صورت موفقیت آمیز بودن نتیجه رو ذخیره کنه

+ اما اگر موفق نشد ، مسیج رو میندازه تو تاپیک **orders-retry-topic** 

+ اگه تعداد **retry** ها به سقف رسید اون رو میندازه توی تاپیک **orders-dlq**

+ با لاجیک مشخص مسیج ها از تاپیک های **orders-retry-topic** و **orders-dlq** تلاش میشه دوباره پردازش بشه

+ و در نهایت شبیه به همین سناریو در برگشت و پاسخ باید پیاده سازی بشه




### سناریو های پیچیده

+ فرض کنیم توی سیستم ، یه کیف پول داریم ، همچنین باید چند سرویس کار هایی انجام دهند و مبلغ از کیف کسر میشه ، اما آخرین سرویس به خطا می خوره ، حال باید چکار کرد؟

جواب :

ابتدا باید در نظر داشت سرویس کیف پول باید برای هر تراکنش ، وضعیت داشته باشد ، pend , complete , fail 

همچنین باید یه تیبل دیگه هم داشت که لاگ 

های ساگا رو هندل کنه 

به این صورت که هر استیت تغییر میکنه ، باید یه سطر به ایت تیبل اضافه بشه


+ اگر یه سرویس دیر بخواد پاسخ بده ، چه ارتباطی باید داشته باشن، و این که چه نکاتی باید توجه بشه؟

ارتباط غیر همزمانه ، برای آگاهی درخواست کنند ه باید از polling , webhook  استفاده بشه ، بعد درخواست به کاربر صفحه ی در حال پردازش نشون میدیم ، بعد پاسخ بهتره با پیامک یا ایمیل اطلاع بدیم

برای جلوگیری از race condition  باید از صف استفاده کنیم


باید خطا ها رو هندل کنیم ، مانند retry ,  همچنین باید تضمین کنیم که پیام به دست سرویس میرسد

+ آیا کافکا تضمین میدهد که پیام میرسه به دست کانسیومر؟

At Least Once

به این معنا که کافکا تضمین میکند حداقل یه بار تاپیک ها به کانسیومر بدهد اما شاید بیش از یک بار پیام بدهد

Exactly Once

در این حالت تضمین میکند که تنها یک بار بدهد

با این وجود باید مدیریت خطا رو در کد هم در نظر بگیریم

dead-letter queues یا retries 

برای بازیابی پیام‌های از دست رفته هنگام کانسیوم کردن

همچنین در صورتی که از سرویس جوابی دریافت نکردیم میتون از retry  استفاده کرد فقط توجه داشته باشید که idempotent را رعایت کنیم

نتیجه گیری

Kafka به خودی خود نمی‌تواند تضمین قطعی دهد که پیام حتماً به مصرف‌کننده (سرویس مرچنت) می‌رسد، زیرا این بستگی به تنظیمات 

شما و نحوه پیاده‌سازی مصرف‌کننده دارد. اما با استفاده از تنظیمات مناسب مانند Exactly Once Delivery و Ack Mechanisms، می‌توانید احتمال رسیدن قطعی پیام‌ها را به طور قابل توجهی افزایش دهید و همچنین با مکانیزم‌هایی مانند retry و dead-letter queue می‌توانید از پیام‌های از دست رفته جلوگیری کنید
