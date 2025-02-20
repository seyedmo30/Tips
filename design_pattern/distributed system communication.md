
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

ابزار هایی مانند کافکا یا ربیت می توان استفاده کرد ، پترن هایی مثل Publish-Subscribe  CQRS - Event-Streaming - 

**Key Features** : Event producers, event consumers, event channels, event processing

**Challenges سختی هایی که داره** Complexity - Consistency - Latency - Debugging and Monitoring - Security

