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



`sudo apt install openfortivpn`

`sudo openfortivpn --trusted-cert 724e45fe738be47c769e09798847f7c1e07c0026a6ba7bab58cd641ef341f723`

```
host = 5.160.235.50
port = 443
username = smo.hosseini
password = a
```

***********


# shecan
178.22.122.100 ,185.51.200.2

# ssh
mostafa@10.21.10.8


git clone https://git.maani.app:glpat-TyUZNaMiucjkFzMGnDGc@git.maani.app/maani/backend/direct-debit.git

# chargoon (windows server -remmina) 

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


#### mysql dev read_only

فقط با php myadmin میشه رفت توش

https://mysql.dev.maani.app/

user: report

pass: hjsatyc656afbxhgdFsd56

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

### راه راحت لاگین در SWAGGER

ابتدا با کد ملی و پسورد و ورژن ۲ لاگیم می کنیم

**https://api.dev.maani.app/api/v1/auth/login**

```
curl -X 'POST' \
  'https://api.dev.maani.app/api/v1/auth/login' \
  -H 'accept: application/json' \
  -H 'device-id: bd336d9e-4fc9-4b90-fbd3-36d9e4fc9b90' \
  -H 'Content-Type: application/json' \
  -d '{
  "locationLat": 35.725285,
  "locationLng": 51.442619,
  "version": "2.0.0",
  "channel": "Google Play",
  "manufacturer": "Samsung",
  "system": 12,
  "androidApiLevel": 12,
  "model": "SM-G998B",
  "resolution": "1600x720",
  "network": "4G",
  "carrier": "Irancell",
  "loginBy": "Credentials",
  "loginFrom": "Direct",
  "username": "0019207670",
  "password": "Salam123"
}'
```
سپستوکن رو کپی می کنیم تو authorizations

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

+ **مدیریت استپ ها**

در مانی به دلیل استفاده از اپ موبایل آپدیت کردن خیلی هزینه داشت و پروداکت تمایل به فورس آپدیت نداشت ، و کاری که کردند در موبایل مراحل یک اکشن  در دیتابیس بوده و از بک اند مراحل پرسیده میشود ، با این تریک دیگه نیاز نیست تغییرات جزعی مانند مراحل و استپ ها  در صورت تغییر نیاز به آپدیت اپ باشد



+ **B2b**

نکاتی که باید هنگام کار با این پروژه ها توجه شود

همواره باید مستندات پرووایدر خوانده شود 

Source of truth باید مشخص شود و اما نباید کلید اصلی بر اساس پروایدر طراحی شود 

در صورتی که شاید از چند پروایدر برای نیازمندی واحد استفاده میکنیم ، بهتر است نام متد ها و جداول دیتابیس رو جنرال نام گذاری کنیم

بهتر است هر چه زود تر sandbox  و راه ارتباطی با پروایدر رو پیدا کرد

همیشه باید سنارو هایی رو دید که پروایدر به صورت موقت یا دایم از دسترس خارج شود

+ به نظرم باید مشخص شه زمانی که از پروایدر خدمات می گیریم ، کجا هزینه پرداخت میشه ، مثلن اگر تنها **draft**  ایجاد می کنیم آیا هزینه ای دارد یا زمانی که  **complete**  میشه باید هزینه بدیم

+ توی دیتابیس ساختار منظمی به ذهنم رسید که تمام شناسه هایی که درخواست من توشه رو با نام  **TransactionId**  میشه یافت همچنین شناسه ای که بعد ثبت  پروایدر برای پیگیری به من می ده رو با نام  **ReferenceId**  میشه یافت

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



### داستان sync و مشکلات آن

گاهی اوقات اپ درخواست ایجاد پرونده رو میزنه و میگه hasDigitalSignature یعنی داره یا نه ، گاهی برنامه تازه نصب شده و در دیتابیس جدید است و مراحل happy path  میره جلو در این صورت در دیتابیس و پروایدر ساخته میشه

اما گاهی طرف قبلن امضا داشته ولی clear cache  کرده

حال دوباره hasDigitalSignature چک میشه ، اگه false  باشه ابتدا sync  صدا زده میشه سپس register ، بر این باوریم هرگاه register  از سمت app صدا زده میشه آنگاه hasDigitalSignature نبوده تو app  پس باید تمامی امضا ها باطل شود و از نو ساخته باشه

مشکل فعلی sync هم اینه که اگر بد موقع صدا زده شه ، امضا ها رو باطل میکنه


**راه حل به ذهن من میرسه**

به جای sync یک api میدیم که باطل میکنه ، حال کلاینت میتونه با استفاده از شناسه یوزر بخواد که امضا ها رو باطل کن ، اینجوری مشکل سینک بودن بر طرف میشه و اگر خواست register  کنه با خطا مواجه میشه ، چون درخواست کننده میدونه  که باید دوباره register  کنه طبق **hasDigitalSignature** ابتدا revoke  کنه سپس register  کنه

**درون جدول  امضای دیجیتال چند bool داره که وضعیت های مختلفی رو میتونه داشته باشه**

`isRevoked - isActive - isRegistered`

ما می خواییم به جای این همه ستون ، یه ستون status  بزاریم ، نکته ای که باید توجه کنیم ، اینه که در امضای دیجیتال ، شاید برای تکمیل امضا چندین استپ باید گذرونده  بشه ، هر استپ ، وضعیت منحصر به فرد خودشو داره پس این دو وضعیت باید ادغام شود :

`isRevoked - isRegistered`

اما به نظرم بولین isActive باید بمونه زیرا ربطی به روند خطی نداره



# on_boarding


## actor

### user

وام گیرنده
### admin

ادمین مانی

### company


کمپانی مثل ایران خودرو  که برای کارکنانش طرح تعریف میکنه  مثلن ۵۰ میلیون ۶ ماهه

### merchant

فروشند مثل اسنوا یا دیجیکالا

### fund_provider

تامین مالی بانک ملی یا بلو بانک

### offline_fund_provider(manual_fund_provider)

لیزینگی ها مثل رایا

### legal_user

نمایندگی های فروشنده مثل نماینده های اسنوا

