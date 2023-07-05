تعداد داکیومنت های ایندکس 

http://localhost:9200/exploit/_count

ساخت ایندکس جدید

PUT  http://192.168.13.248:9200/exploit

حذف ایندکس

DELETE  http://192.168.13.248:9200/exploit

گرفتن داکیومنت

http://192.168.13.248:9200/exploit/_doc/5

لیست داکیومنت

http://192.168.13.196:9200/exploit/_search/?size=1000&pretty=true
