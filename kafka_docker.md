
# kafka

data pipelines - distributed event store -  Kafka is not a traditional message queue -  Kafka is a distributed messaging system -  RabbitMQ is a message broker, while Kafka is a distributed streaming platform. One of the primary differences between the two is that Kafka is pull-based, while RabbitMQ is push-based -Kafka offers much higher performance than message brokers like RabbitMQ - bus using a pub-sub model of stream-processing

# rabbitMQ

RabbitMQ is message-broker software - it allows for various protocol extensions via a plug-in architecture -  message broker between microservices

# kafka vs rabbit

RabbitMQ cannot replay events ربیت زمانی می تواند مفید باشد که برای پخش مجدد پیام ها در یک موضوع به این ویژگی نیاز ندارید
RabbitMQ برای برنامه هایی که معماری برنامه ناشناخته است بهتر است و با بیان مشکل و راه حل توسعه و تکامل می یابد. RabbitMQ در این شرایط در مقایسه با کافکا بسیار انعطاف پذیرتر و آسان تر است. اما زمانی که برنامه بالغ شد و نیاز به مقیاس‌پذیری، توان عملیاتی زیاد، قابلیت اطمینان، استحکام و قابلیت پخش مجدد پیام‌ها وجود داشت، RabbitMQ تبدیل به یک گلوگاه می‌شود و بهتر است به کافکا بروید.

messaging system : یک سیستم پیام رسان وظیفه انتقال دیتا ازبرنامه ای به برنامه دیگر را دارد  , دو نوع الگوی پیام رسان در دسترس است point-to-point و دیگری publish-subscribe

Point-to-point message system : پیام ها در صف هایی قرار می گیرند و یک یا چند مصرف کننده می تواند از پیام ها استفاده کنند اما تنها یک پیام خاص را یک مصرف کننده می تواند مصرف کند هنگامی که یک مصرف کننده یک پیام را در صف می خواند پیام از آن صف می تواند خارج شود

Publish-subscribe message system : در این سیستم پیام ها بر اساس topic یا موضوعات هستند، برخلاف point-to-point system زمانی که مصرف کنندگان از پیام ها استفاده می کنند پیام ها باقی می مانند پیام ها در یک یا چند topic هستند و مصرف کنندگان به همه ی پیام هایی که به topicهای مختلف است دسترسی دارند

topic : stream پیام هایی که متعلق به یک دسته خاصی هستند topic نامیده می شوند و دیتاها در topicها ذخیره می شود.Topicها به پارتیشن هایی تقسیم می شوند برای هر topic Kafka حداقل یک پارتیشن را در نظر می گیرد

پاتیشن : تاپیک  ممکن است پارتیشن های زیادی داشته باشد، بنابراین می تواند مقدار دلخواه داده را مدیریت کند

Partition offset : هر  پیام، یک شناسه منحصر به فردی دارد که آفست نامیده می شود

Broker : بروکر ها  سیستم ساده ای هستند که مسئول نگهداری داده های منتشر شده هستند. هر کارگزار ممکن است صفر یا بیشتر پارتیشن در هر تاپیک داشته باشد

Producers : تولیدکنندگان داده ها را برای کارگزاران کافکا ارسال می کنند
 . هر بار که یک تولید کننده پیامی را برای یک کارگزار منتشر می کند، کارگزار به سادگی پیام را به آخرین فایل بخش اضافه می کند .  در واقع، پیام به یک پارتیشن اضافه می شود. سازنده همچنین می تواند به پارتیشن مورد نظر خود پیام ارسال کند    .
Leader : لیدر نودی است که مسئول تمام خواندن و نوشتن برای پارتیشن داده شده است. هر پارتیشن یک سرور دارد که به عنوان رهبر عمل می کند.

Follower : گره‌ای که از دستورالعمل‌های رهبر پیروی می‌کند، فالوور نامیده می‌شود. اگر رهبر شکست بخورد، یکی از پیروان به طور خودکار رهبر جدید می شود. یک فالوور به عنوان مصرف کننده عادی عمل می کند، پیام ها را کانسوم می کند و  فروشگاه داده خود را به روز می کند









مشاهده لیست تاپیک ها  
sudo docker exec  -it 69078a9693f0 kafka-topics  --bootstrap-server localhost:9092 --list

نوشتن پیام در تاپیک            

sudo docker exec  -it 69078a9693f0 kafka-console-producer  --topic hafizium_178_exploit --bootstrap-server localhost:9092       

خواندن پیام در تاپیک

sudo docker exec  -it 69078a9693f0 kafka-console-consumer  --topic hafizium_178_exploit --bootstrap-server localhost:9092

ساخت تاپیک جدید                
 kafka-topics --create --topic quickstart-events --bootstrap-server localhost:9092         



حذف تاپیک

kafka-topics --delete --topic hafizium_178_exploit --bootstrap-server localhost:9092

تعداد مسیج در تاپیک

sudo docker exec -it 69078a9693f0 kafka-run-class kafka.tools.GetOffsetShell 
 --bootstrap-server localhost:9092 
  --topic hafizium_178_exploit 
  | awk -F  ":" '{sum += $3} END {print sum}'
  
  
