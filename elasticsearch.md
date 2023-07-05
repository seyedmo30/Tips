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

اینسرت تکی

curl -POST 'localhost:9200/exploit/_doc' -H 'Content-Type: application/json' -d'
{
        "timestamp": "2018-01-24 12:34:56",
        "message": "User logged in",
        "user_id": 4,
        "admin": false
} '

