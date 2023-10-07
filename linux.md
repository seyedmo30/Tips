sudo kill -9 $(sudo lsof -t -i:8000)

scp root@1.2.3.4:/root/pcfilename.rar ./

مشاهده پرت های باز tcp udp 

    sudo apt install net-tools
    netstat -tulpn

مشاهده سایز فولدر 

    sudo du -hs volumes

پیدا کردن گنده ترین فایل ها

    find -type f -exec du -Sh {} + | sort -rh | head -n 5

برای اجرای دستورات در سیستم عامل مانند تعریف یوزر ، ساخت کرونتب و اجرای بش اسکریپت و تمامی دستوراتی که با دسترسی سودو نیاز به اجرا هست . می توان گفت پس از اینیت کردن سیستم عامل ، با استفاده از cloud-init می توان دستوراتی در سیستم عامل اجرا کرد

