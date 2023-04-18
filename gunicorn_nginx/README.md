# gunicorn



ابتدا فایل gunicorn_conf.py را در دایرکتوری پروژه می سازیم ، سپس دستور پایین را می زنیم

gunicorn -c index/gunicorn_conf.py index.wsgi    

حال باید پروژه با ۳ ورکر ران شود

# nginx 

حال باید ابتدا فایل nginx.conf را باز نویسی کنیم ، دقت کنید در صورتی که یوزر،  www-data باشد باید تمامی دایرکتوری های static & media  را به این یوزر ، دسترسی دهیم

/etc/nginx/nginx.conf

سپس باید اطلاعات درون index_nginx.conf را در default کپی کنیم

/etc/nginx/sites-available/default

 و در نهایت 
 
 sudo systemctl restart nginx.service
 
 
