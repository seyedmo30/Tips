### Core Services

+ **API Gateway**

بهتر است تنها یک entry point داشته باشیم برای مدیریت موارد زیر : 

Handles routing, authentication, authorization, rate limiting, load balancing, logging, and monitoring.


+ **Authentication and Authorization**

مدیریت یوزر ها و دسترسی به کامپوننت ها


+ **Service Registry and Discovery**

کمتر مورد استفاده است و وظیفش اینه که بتونیم به سرویس ها دسترسی داشته باشیم و اونها رو با یه نام مشخص کنیم

در حقیقت این یه سرویس هست که بتونیم به جای یه آدرس ، یه نام برای سرویس داشته باشیم و با یه نام مشخص کنیم و در صورتی که بخوایم این سرویس رو به جای یه سرویس دیگه قرار دهیم می تونیم اون رو به جای یه سرویس دیگه قرار دهیم 


**Infrastructure Services**

یکم زیر ساختی تر :

+ **Configuration Management**

یه کانفیگ مرکزی مثل والت

+ **Monitoring and Logging**

Examples: Prometheus, Grafana, ELK

+ **Message Brokers/Event Buses**

asynchronous communication and event-driven architectures.

Examples: Apache Kafka, RabbitMQ

+ **Load Balancers**

Examples: NGINX, HAProxy

+ **Health Checking**

Examples: Kubernetes Liveness

+ **Container Orchestration**

مدیریت کامپوننت ها / کانتینر ها و سرویس ها مقیاس پذیری ، اتوماتیک دیپلوی

Examples: Kubernetes, Docker Swarm

+ **Service Mesh**

مدیریت  ارتباط سرویس ها با هم  مانند سیکیوریتی ، قابلیت مشاهده ی یک دیگر سرویس ها ،  و ترافیک بین آنها
