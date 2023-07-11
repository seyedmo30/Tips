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

لیست ایندکس ها 

http://localhost:9200/_cat/indices


اینسرت تکی

curl -POST 'localhost:9200/exploit/_doc' -H 'Content-Type: application/json' -d'
{
        "timestamp": "2018-01-24 12:34:56",
        "message": "User logged in",
        "user_id": 4,
        "admin": false
} '

curl -X PUT "localhost:9200/exploit/_bulk?pretty" -H 'Content-Type: application/json' -d'
{ "create": { } }
{ "@timestamp": "2099-05-07T16:24:32.000Z", "event": { "original": "192.0.2.242 - - [07/May/2020:16:24:32 -0500] \"GET /images/hm_nbg.jpg HTTP/1.0\" 304 0" } }
{ "create": { } }
{ "@timestamp": "2099-05-08T16:25:42.000Z", "event": { "original": "192.0.2.255 - - [08/May/2099:16:25:42 +0000] \"GET /favicon.ico HTTP/1.0\" 200 3638" } }
'

مرتب بر اساس آیدی

http://192.168.13.248:9200/exploit/_search/?size=1000&pretty=true&sort=_id
