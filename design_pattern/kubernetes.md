## Kubernetes Concepts

+ Cluster

مجموعه ای ماشین ها - مجازی یا واقعی - می گویند . که شامل یک نود اصلی - مستر - می باشد و یک یا چند ورکر دارد 

+ nodes

به هر ماشین تو کلاستر کردن یه نود می گن و نود اصلی مستر هست و به وسیله ی Kubelet شناسایی می شن و با مستر ارتباط می گیرن

دو نوع نود داریم - مستر - ورکر

+ pod

گروهی از کانتینر ها ، شبکه و ستوریج مشترک و یک پاد ، آی پی  واحد دارند

در واقع کوچک ترین  واحد دپلوی . یک اینستنس  که میتوان یک یا چند کانتینر را شامل شود که ریسورس بین آنها مشترک است

می توان به زبون ساده گفت یک اینستنس از یک داکر کامپوس است یعنی مجموعه از سرویس و ابزار هاش

+ Namespace

یک کلاستر مجازی است ، در حقیقت برای تقسیم بندی منابع بر روی یک کلاستر ، می توان از این مفهوم استفاده کرد



+ kubectl

یک اینترفیس است برای ارتباط با کلاستر

+ etcd

یک متغییر گلوبال برای پیکر بندی هست و و اطلاعات همه ی نود ها را دارد

+ objects

به مجموع deployments, replica sets, services, pods از یک واحد در کلاستر میگن

+ + ReplicaSet

می توان تعداد تکرار پاد ها را با استفاده از لیبل مدیریت کرد

+ + Deployment

استقرار یک پاد یا رپلیکا را مدیریت می کند

 گروهی از Pods را مدیریت می‌کند و اطمینان می‌دهد که آنها در حالت دلخواه و مقیاس خودکار اجرا می‌شوند.


+ minikube

یه نسخه ی سبک برای پیاده سازی روی یه ماشین ، بعد از استارت کردن مینی کیوب، به صورت پیشفرض یه پاد ایجاد میشه

  ### kubectl command
 این دستور یک دولوپمنت می سازد
 
       kubectl run nginx --image=nginx

# commands

###  کلاستر تستی رو میسازه

‍`minikube start --driver=docker`



###  نود ها رو میشه دید

`kubectl get nodes`


###  دیپلویمنت ها رو میشه دید    

`kubectl get deployments`


###  سرویس ها رو میشه دید    

`kubectl get services`


**node port**

 یکی از سرویس های خوب برای تست که شبیه به `load balancer` در محیط تست عمل میکند 

### رپلیکا ست - ReplicaSet

با استفاده از این می توان تعداد پاد ها رو مشخص کرد ، حتی اگه کرش کنن یا پاک شن ، این ابزار به تعداد نگاه می کنه و تعداد پاد ایجاد می کنه

`kubectl get rs`


### Service

برای مدیریت درخواست ها و توضیع بین پاد ها شبیه load balancer

مهم ترین سرویس ها :

+ **ClusterIP** (default): Internal access within the cluster.
+ **NodePort**: Exposes a service on each node's IP and a static port.
+ **LoadBalancer**: Exposes the service to the internet using a cloud provider's load balancer.
+ **ExternalName**: Maps a service to an external DNS name.




### دیپلوینمت

با استفاده از این می تونیم نسخه ی برنامه ها رو مدیریت کنیم و `upgrade - down grade`  انجام بدیم

نحوه ی کارش اینجوریه بالا سر رپلیکا میشینه و می بینه اگه نسخه ای که بهش دستور دادن با رپلیکا ها یکی نبود ، اونا رو پاک می کنه

هر بار دیپلوی منت جدید بدیم توی لیست و تاریخچه ی رپلیکا ست ها ، یک نسخه افزوده میشه

`kubectl get rs`

با دستور بالا می بینیم که سطر اضافه میشه

جالبه که توی دستور پایین اطلاعات خاصی نیست :


`kubectl get deploy`



###  سرویس ها رو با مینیکیوب میشه دید    

`minikube service nginx`



###  پاد ها رو میشه دید

`kubectl get pods`



###  حذف پاد

توجه شود که این دستور در صورتی پاد رو پاک می کنه که رپلیکا ست دوباره این پاد رو نسازه

`kubectl delete pod test`

اگر دیدیم بعد از پاک کردن دوباره ایجاد شد ، باید بریم رپلیکا ست رو استاپ کنیم


###  کانفیگ

`cat ~/.kube/config`


توی این فایل ۳ بخش اصلی داره 

+  اولیش اطلاعات کلاستره مثل آدرس سرور و ورژن و پروایدرش (مثل مینی کیوب یا k3s )

+  یوزر و اطلاعاتی که هر یوزر داره

+  کانتکس که ارتباط بین کلاستر ویوزر هست  و این که هر یوزر به کدوم کلاستر ها دسترسی داره یا هر کلاستر چه نیم اسپیس هایی داره



#### tips

+ اگه با یوزر root  بخواهیم بریم دستورات رو اجرا کنیم احتمالن پیدا نکنه زیرا رو نمی ره از کانفیگ یوزر بخونه


###  گرفتن ایمیج

‍`kubectl get svc nginx`

###  اکسپوز کردن پورت پاد


`kubectl expose deployment nginx --type=NodePort --port=80 --name=nginx`


## روش مشاهده پاد مرحله به مرحله

 + پیدا کردن IP برای تست و درخواست

‍`minikube ip`


+ روش اگزپوز کردن یک پاد و پورتش


`kubectl expose pod pod-test --type=NodePort --port=8080`

+ بعد این باید سرویس رو کال کرد تا ببینیم به چه پورتی مپ شده

`kubectl get svc`

+  پاد ها رو میشه دید

`curl http://192.168.49.2:32593`

### متداول ترین راه ها برای دیباگ و یافت خطا

+ Check Pod Status: `kubectl get pods`
+ Describe Pod: `kubectl describe pod <pod-name> to check for events and errors.`
+ Logs: `kubectl logs <pod-name> for application logs.`
+ Exec into Pod: `kubectl exec -it <pod-name> -- /bin/bash to explore the container’s environment.`
+ YAML Validation: `Ensure the Pod specification file is correct.`



### چگونه ریسورس رو مدیریت می کنیم


```yml

apiVersion: v1  
kind: Pod  
metadata:  
  name: resource-limits-pod  
spec:  
  containers:  
  - name: nginx  
    image: nginx  
    resources:  
      requests:  
        memory: "64Mi"  
        cpu: "250m"  
      limits:  
        memory: "128Mi"  
        cpu: "500m" 

```