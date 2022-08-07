sudo docker run یا pull نام ایمیج

دانلود و اجرای ایمیج 

docker run -d nginx

اگر بخواهیم بدون تی ماکس برنامه دی اتچ کار کنه 

sudo docker images

لیست ایمیج ها 

sudo docker ps

تمام کانتر هایی که در حال اجرا هستند 

sudo docker ps -a

لاگ تمام کانتینر ها 

sudo docker rm <id>
  

docker rm $(docker ps -a -q)

  
  
پاک کردن کانتینر
docker stop id
  
ایستاپ کردن کانتینر 

docker start id
  
  
  
برای استارت کانتینر استاپ شده 

sudo docker rmi
  
  
پاک کردن ایمیج
sudo docker container prune

  
پاک کردن ایمیج های بدون تگ یا آویزان
  
docker image prune --filter="dangling=true"
  
  
تمامی کانتینر ها پاک میشن 

sudo docker network create mynetwork
  
ساخت شبکه 


  
  --rm اتماتیک بعد از باز شدن ریموو می شه

  
  -it می گه توش بمون 

  
  -p6379:6379 میگه از پورت داخل شبکه داکر ، رو پورت سیستم عامل هم نشون بده

  
  --name redis نام کانتینر تو داکر

  
  --network mynetwork

  
  -w /var/ مکان دایرکتوری در شروع
  
  
  -v /var/volume/sla/django:/data

  
  sudo docker run --rm  --name resis -p6379:6379 redis 


  sudo docker exec -it <id> <bash>
می ره دستور بش رو داخل کانتینر اجرا می کنه
آی تی داخل اون می مونه 


  docker login -u seyedmo30 -p 12345
لاگیدن در داکر هاب و گرفتن ایمیج تا ۲۰۰ تا


تو پروژه ها بهتره هر جا آی پی یا لوکال هاست بود
اسم کانتینری که دادیم رو بزاریم

  
  
  
  
  می تونیم بدون استفاده از داکر کامپوز ، داکر فایل رو بسازیم و از اون بیلد بگیریم
  
  
  docker build -t <<name>>:<<version>> <<directory>>
  
  
  docker build -t food:1.0 ~/test/
  
  
  
