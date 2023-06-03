
# kafka

data pipelines - distributed event store -  Kafka is not a traditional message queue -  Kafka is a distributed messaging system -  RabbitMQ is a message broker, while Kafka is a distributed streaming platform. One of the primary differences between the two is that Kafka is pull-based, while RabbitMQ is push-based -Kafka offers much higher performance than message brokers like RabbitMQ - bus using a pub-sub model of stream-processing

# rabbitMQ

RabbitMQ is message-broker software - it allows for various protocol extensions via a plug-in architecture -  message broker between microservices

# kafka vs rabbit

RabbitMQ cannot replay events ربیت زمانی می تواند مفید باشد که برای پخش مجدد پیام ها در یک موضوع به این ویژگی نیاز ندارید
RabbitMQ برای برنامه هایی که معماری برنامه ناشناخته است بهتر است و با بیان مشکل و راه حل توسعه و تکامل می یابد. RabbitMQ در این شرایط در مقایسه با کافکا بسیار انعطاف پذیرتر و آسان تر است. اما زمانی که برنامه بالغ شد و نیاز به مقیاس‌پذیری، توان عملیاتی زیاد، قابلیت اطمینان، استحکام و قابلیت پخش مجدد پیام‌ها وجود داشت، RabbitMQ تبدیل به یک گلوگاه می‌شود و بهتر است به کافکا بروید.





مشاهده لیست تاپیک ها  
sudo docker exec  -it 69078a9693f0 kafka-topics  --bootstrap-server localhost:9092 --list

نوشتن پیام در تاپیک            

sudo docker exec  -it 69078a9693f0 kafka-console-producer  --topic hafizium_178_exploit --bootstrap-server localhost:9092       


ساخت تاپیک جدید                
 kafka-topics --create --topic quickstart-events --bootstrap-server localhost:9092         
