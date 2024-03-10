# network

        sudo lsof -t -i:8000    //مشاهده PID روی پورت

        ls -l /proc/<PID>/cwd //یافتن دایرکتوری کامند برای اجرا       /usr/local/projects/data-aggrigator/cmd/server

        ls -l /proc/<PID>/exe //یافتن دایرکتوری فایل اگزکیوت        /tmp/go-build2963007683/b001/exe/main

        ip route // مشاهده دیفالت گیت وی

        nslookup //  مشاهده اینکه DNS درست کار می کند ، پیدا کردن IP هاست نیم ها

+ مشاهده پرت های باز tcp udp 
        
            sudo apt install net-tools
            netstat -tulpn
  
+ با دستور ifconfig می توان اطلاعات کلی ماشین مانند default gateway , netmask ,inet  را دید همچنین شبکه اصلی ماشین enp4s0 را دید همچنین شبکه های محلی مانند داکر هم قابل رویت است

                ifconfig
        
sudo kill -9 $(sudo lsof -t -i:8000)



scp root@1.2.3.4:/root/pcfilename.rar ./

مشاهده سایز فولدر 

    sudo du -hs volumes

پیدا کردن گنده ترین فایل ها

    find -type f -exec du -Sh {} + | sort -rh | head -n 5

### cloud-init 
برای اجرای دستورات در سیستم عامل مانند تعریف یوزر ، ساخت کرونتب و اجرای بش اسکریپت و تمامی دستوراتی که با دسترسی سودو نیاز به اجرا هست . می توان گفت پس از اینیت کردن سیستم عامل ، با استفاده از cloud-init می توان دستوراتی در سیستم عامل اجرا کرد

    cloud-init status

#### env 
در صورتی که مقدار env را در ترمینال export کنیم ، بعد از ریلود شدن ، پاک می شن و برای ماندگاری باید در bash ذخیره کینم و در پایان، بش را source کنیم ، همچنین می توان برای استفاده همه ی یوزر ها ، در آدرس زیر ذخیره کرد

    /etc/environment



#### Package managers
ابزاری برای نصب برنامه و پیش نیاز ها و مدیریت آن ،برای دبیان apt هست و برای فدورا yum

#### service manager
برای استارت و استاپ شدن سرویس ها مثل systemd
, همچنین برای اینکه دستی سرویسی رو استاپ یا ریاستارت کنیم از دستور systemctl  استفاده میکنیم
همچنین با دستور enable , desable مشخص میکنیم که بعد از بوت شدت ، اتوماتیک استارت شود

#### Ansible
می تونه همزمان به کلی یرور وصل شه و یه دستور رو اگزکیوت کنه ، مثلا بجای اینکه ssh بزنیم به تک تک سرورا و شکن ست کنیم ....

#### Terraform
کانفیگ ها ی اولیه برای اینکه سرور بیاد بالا

#### limit docker
می تونیم بر بروی کانتینر ها لیمیت بزاریم

--cpus _ تعداد هسته هایی که در اختیار قرار میدیم رو مشخص می کنیم

CPU CFS scheduler period مطالعه شود

1000 milisecond در اختیار میزاریم

CPU CFS quota on the container مطالعه شود

Cpu realtime scheduler vs cpu cfs scheduler

#### shell
+ Bash

 خیلی از فیچر هارو نداره ، اتوکامپلیت نداره 
بعد از باز کردن بش ، بدون اینکه کی لاگین کردن از .bashrc می خونه و اصطلاحا می گن interactive non-login
  
+ Zsh

کلی فانکشن داره ، اتو کامپلیت خوبی داره ، میشه 
اکتنشن اضافه کرد بهش ،  
چیزایی که تو bashrc داریم:

Export path=
آدرس یه فایل باینری

#### Env
مقادیری در لینوکس به صورت کی ولیو ، اطلاعات زیر درونش است :

• System and user paths

• API keys and credentials

• Software configurations

• Default settings for shell and commands

+ Printenv / env 

دستور بالا تمام مقادیر را نمایش میدهد

ذخیره به صورت موقت : 
اکسپورت در ترمینال
ذخیره به صورت پایدار :

nano /etc/environment 
و اجرا
source /etc/environment 
البته در آدرس زیر هم میشه
`~/.profile` or `~/.bashrc` 

می تونیم توی دایرکتوری زیر هم چند فایل بسازیم 
/etc/profile.d/*.sh
