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

sudo docker rm  id
  

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

نکته مهم والیوم : توجه شود در صورتی که داکر به صورت عادی ران شود ، دایرکتوری درون کانتینر پاک شده و دایرکتوری تعریف شده سیستم جای آن می نشیند 
  
  sudo docker run --rm  --name resis -p6379:6379 redis 


  sudo docker exec -it id bash
می ره دستور بش رو داخل کانتینر اجرا می کنه
آی تی داخل اون می مونه 


docker run -ti project_web bash

در صورتی که کانتینر در حالت اگزیت است و یا هر بار ران می گیریم ، می خوابه ، می تونیم اینجوری رانش کنیم و ببینیم چشه؟


  docker login -u seyedmo30 -p 12345
لاگیدن در داکر هاب و گرفتن ایمیج تا ۲۰۰ تا


تو پروژه ها بهتره هر جا آی پی یا لوکال هاست بود
اسم کانتینری که دادیم رو بزاریم

  
  
  
  
  می تونیم بدون استفاده از داکر کامپوز ، داکر فایل رو بسازیم و از اون بیلد بگیریم
  
  
  docker build -t name:version directory
  
  
  docker build -t food:1.0 ~/test/
  
  
------------------------------------------------------------------------------------------------------

# postgres

برای ران کردن حتما باید پسورد داده شود

docker run -itd -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=salam -p 5432:5432 -v /data:/var/lib/postgresql/data --name postgresql postgres:15-alpine

سپس درونش اگزک می کنیم

sudo docker exec -it 98d7a843febd sh

به جای ورود متداول اینجوری می ریم توی پستگرس

PGPASSWORD=salam psql -U postgres
  

--------------------------------------------------------------------------------------------------------

# Dockerfile


FROM alpine:3.17

همیشه باید از یک ایمیج شروع کنیم به ساختن

FROM golang:1.20-rc-alpine as builder



RUN mkdir /app

دستورات رو اینجوری درون ایمیج میزنیم


WORKDIR /app

چنج دایرکتوری درون ایمیج


COPY . /app

کپی از سرور (آدرس اول ) به ایمیج (آدرس دوم) 


CMD /app/salam

اجرای دستور نهایی 


