

        sudo lsof -t -i:8000    //مشاهده PID روی پورت

        ls -l /proc/<PID>/cwd //یافتن دایرکتوری کامند برای اجرا       /usr/local/projects/data-aggrigator/cmd/server

        ls -l /proc/<PID>/exe //یافتن دایرکتوری فایل اگزکیوت        /tmp/go-build2963007683/b001/exe/main

ip route __ مشاهده دیفالت گیت وی

nslookup __  مشاهده اینکه DNS درست کار می کند ، پیدا کردن IP هاست نیم ها


        
sudo kill -9 $(sudo lsof -t -i:8000)



scp root@1.2.3.4:/root/pcfilename.rar ./

مشاهده پرت های باز tcp udp 

    sudo apt install net-tools
    netstat -tulpn

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
