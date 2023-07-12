تعداد داکیومنت های ایندکس

http://localhost:9200/exploit/_count

ساخت ایندکس جدید

PUT  http://192.168.13.248:9200/exploit

حذف ایندکس

DELETE  http://192.168.13.248:9200/exploit

گرفتن داکیومنت

http://192.168.13.248:9200/exploit/_doc/5

گرفتن داکیومنت بدون متا دیتا

http://192.168.13.248:9200/exploit/_source/5

گرفتن  مالتی داکیومنت ( مطالعه شود mget )

http://192.168.13.248:9200/exploit/_mget

لیست داکیومنت

http://192.168.13.196:9200/exploit/_search/?size=1000&pretty=true

لیست ایندکس ها

http://localhost:9200/_cat/indices?v=true

# mget

اگر بخوایم مشخص کنیم چه فیلد هایی می خواییم و چه فیلد هایی نمی خواییم :

GET /_mget

        {
          "docs": [
            {
              "_index": "my-index-000001",               مشخص می کنیم بر روی کدام ایندکس
              "_id": "1",                                 مشخص می کنیم بر روی کدام ایدی
              "_source": [ "field3", "field4" ],           مشخص می کنیم  کدام فیلد ها رو می خوایم
            },
            {
              "_index": "my-index-000001",
              "_id": "2",
                        "_source": {
                        "include": [ "user" ],                       مشخص می کنیم  کدام فیلد ها رو می خوایم
                        "exclude": [ "user.location" ]                مشخص می کنیم  کدام فیلد ها رو نمی خوایم
                        "stored_fields": [ "field1", "field2" ]         ترتیب
                        "routing": "key2"                               انتخاب شارد
                      }
            }
          ]
        }



# query

خیلی قوی تر از مالتی گت هست و داخل سرچ استفاده میشود


درون api search قرار دارد و بعد از کوییری ، می بایست مشخص کینم



        GET /_search
        {
          "query": {                                باید درون کوییری مشخص کنیم که match , term , range ,....
            "match": {                                   در متچ ، باید کله دقیق برابر مقدار باشد تا پیدا شود ، مقدار اول کلید و مقدار دوم متن مورد نظر هست
              "field1": {                                                        فیلدی که در مچ باید جستوجو درون آن انجام شود
                "query":
                      {
                        "query": "Programs Rating",                متن جستوجو
                        "operator": "and" ,                        باید تمام متن در خروجی باشد
                        "fuzziness": "AUTO",                        اجازه می دهد که به صورت فازی سرچ کنیم
                      }

              }
            }
          }
        }

نکته :
match -------
  در مچ باید کلمات دقیق باشند .

prefix -------
اگر ابتدای کلمه را می دانیم باید از prefix استفاده کنیم .

match_phrase -----
ترتیب عبارت هم مهم است .

fuzzy -----
غلط های املایی رو درست می کنه .









اگر بخواهیم درون یک کوییری چند شرط بزاریم ، باید از bool استفاده کنیم ، خود بول باید یکی از موارد must و filter و should و must_not را بگیرد .

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
