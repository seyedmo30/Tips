#dns remote
10.12.0.1

# git 
https://git.maani.app/
10.21.13.219

# vps windows
10.21.13.97

# mattermost
connect.maani.app

# vps
31.47.50.2

#  confluence
confluene.maani.app

# user
smo.hosseini
m.hosseini

## forticlient

+ servcer
  
5.160.235.50

+ username
smo.hosseini

+ password

***********

# pass
aE_c$dx9
SalaM#1234
 
# shecan
178.22.122.100 ,185.51.200.2

# ssh
mostafa@10.21.10.8


git clone https://git.maani.app:glpat-TyUZNaMiucjkFzMGnDGc@git.maani.app/maani/backend/direct-debit.git

# chargoon

Host : 10.21.13.97
user : m.hosseini
Pass : Aa123456

# env_dev

IP host : 10.21.15.99
Port
loan : 21602
admin : 21606
gateway : 21600
bridge : 21603
authorizer : 21601
zookeeper : 21681
pol : 21605
redis : 21679
mysql : 21636
kafka : 21692
phpmyadmin : 21680
mongo : 21617
notifier : 21604

# mysql

http://10.21.15.99:21180/

user : root

pass: bye

# vault
env.maani.app

token: hvs.CAESIBzHAH1Y-yo6_D-03rPqrWTX-BeQpxzKF9ouKukxsQDIGh4KHGh2cy5JSTJzaTNCdDdOR1lXVnNhZ2VqOHJnTjQ




## swagger

+ برای لاگین در dev  ابتدا تو phpmyadmin باید لاگین کنیم

```
// id خودمون رو پیدا می کنیم

http://10.21.15.99:21680/index.php?route=/sql&pos=0&db=authorizer&table=User

27ed348a-cbe2-4882-8c21-f970c05b87ca


// دیوایس ایدی پیدا می کنیم

http://10.21.15.99:21680/index.php?route=/sql&pos=0&db=authorizer&table=Device

SELECT * FROM `Device` WHERE userId = '27ed348a-cbe2-4882-8c21-f970c05b87ca';

bd336d9e-4fc9-4b90-fbd3-36d9e4fc9b90

```

در نهایت توی سوگر لاگین می کنیم


### forticlient

https://www.fortinet.com/support/product-downloads/linux

Ubuntu 22.04 LTS

Install gpg key

wget -O - https://repo.fortinet.com/repo/forticlient/7.4/ubuntu22/DEB-GPG-KEY | gpg --dearmor | sudo tee /usr/share/keyrings/repo.fortinet.com.gpg

Create /etc/apt/sources.list.d/repo.fortinet.com.list with the following content

deb [arch=amd64 signed-by=/usr/share/keyrings/repo.fortinet.com.gpg] https://repo.fortinet.com/repo/forticlient/7.4/ubuntu22/ stable non-free

Update package lists

sudo apt-get update

Install FortiClient

sudo apt install forticlient
 


### tips

+ **step-substep**

یکی از چیزای جالب مانی ، اینه که استپ های خیلی یخلی زیادی داره ،  و برای مدیریت این داستان  ، هر استپ یه شناسه uuid  داره و توی دیتابیس ترتیب اونا ذخیره شده

در هر مرحله استپ ها ، مدارک ازم برای وام رو تکمیل می کنن ، اپ یا گوشی استیت لسه و باید همیشه به سیستم در خواست بزنه و استپ رو بگیره 



+ **B2b**

نکاتی که باید هنگام کار با این پروژه ها توجه شود

همواره باید مستندات پرووایدر خوانده شود 

Source of truth باید مشخص شود و اما نباید کلید اصلی بر اساس پروایدر طراحی شود 

در صورتی که شاید از چند پروایدر برای نیازمندی واحد استفاده میکنیم ، بهتر است نام متد ها و جداول دیتابیس رو جنرال نام گذاری کنیم

بهتر است هر چه زود تر sandbox  و راه ارتباطی با پروایدر رو پیدا کرد

همیشه باید سنارو هایی رو دید که پروایدر به صورت موقت یا دایم از دسترس خارج شود
+ **config**

با اینکه منطقیه خیلی از دیتا هایی که کانفیگه رو توی env - vault - نگه داریم اما سامانه هایی که ادمین یا پشتیبان داره میشه بعضی از دیتا ها توی دیتابیس ذخیره مثلا : 

+ database socket , kafka socket , اینا رو توی env  میدیم ولی اطلاعات زیر از طریق دیتابیس خونده میشه

+ مبلغ سفته یا امضای دیجیتال که احتمالن هیچ وقت تغییر نمی کنه ، ولی شاید بعد یه سال قیمتا عوض شه

```
CREATE TABLE Configurations (  
    ConfigKey VARCHAR(255) PRIMARY KEY,  
    ConfigValue VARCHAR(255) NOT NULL,  
    Description TEXT,  
    LastUpdated TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,  
    UpdatedBy VARCHAR(100)  
);  
```

sss

‍ff

‍‍‍‍‍‍‍gg
