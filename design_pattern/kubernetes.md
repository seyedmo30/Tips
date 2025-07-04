# Kubernetes Concepts

### Cluster

مجموعه ای ماشین ها - مجازی یا واقعی - می گویند . که شامل یک نود اصلی - مستر - می باشد و یک یا چند ورکر دارد

### nodes

به هر ماشین تو کلاستر کردن یه نود می گن و نود اصلی مستر هست و به وسیله ی Kubelet شناسایی می شن و با مستر ارتباط می گیرن

هر ماشین رو یه نود میگن ، که اصولن یک نود مستر و چندین ورکر داریم

هر نود نسبت به ریسورسی که داره تو خودش pod داره ، بین ۱۰ تا ۱۰۰

دو نوع نود داریم - مستر - ورکر

### pod

resource group هست

گروهی از کانتینر ها ، شبکه و ستوریج مشترک و یک پاد ، آی پی واحد دارند

در واقع کوچک ترین واحد دپلوی . یک اینستنس که میتوان یک یا چند کانتینر را شامل شود که ریسورس بین آنها مشترک است

- داخل یه وروکر نود میشه چندین پاد بسازیم

- درون هر پاد میشه چندین کانتینر بیاریم بالا

- هر پاد آی پی خودش رو داره ، و درونش هر کانتینر پورت خودش رو داره

- bestpractice اینه که هر پاد ، یه کانتینر توش باشه و اینجوری اصل srp هم برقراره و مفهوم

میکروسرویس رو هم پیاده میکنه ، ولی گاهی میشه بیش از یه کانتینر در پاد ایجاد کرد ، مثلن زمانی که ابزار مانیتور یا لاگ یا پروکسی رو ست میکنیم ، همچنین گاهی کانتینر هایی برای اینیت باید بزاریم ، مانند زمانی که می خوایم مایگریشن بزنیم

مانند docker file که میشه یه ایمیج ساخت ، با استفاده از `defination file` هم میشه یه پاد ساخت و بهش گفت که درونش یک یا چند کانتینر از چه ایمیج هایی ساخته شود

**kind**

مشخص میکنیم که نوع آن پاد هست یا سرویس

بعد از ایجاد فایل دفینیشن ، با دستور زیر ، اون پاد را میسازیم

`kubectl apply -f file.yml`

یا

`kubectl create -f file.yml`

همچنین برای دیدن توضیحات بیشتر پاد و لاگ های اون event, از دستور زیر استفاده میکنیم

`kubectl describe pod nginx1`

برای مشاهده پاد جدید ساخته شده دستور زیر

`get pod -n ns1l1`

### Namespace

یک کلاستر مجازی است ، در حقیقت برای تقسیم بندی منابع بر روی یک کلاستر ، می توان از این مفهوم استفاده کرد

با استفاده از این میتوان ریسورس رو گروه بندی یا ایزوله کرد ، بعد از ساخت یه کلاستر به صورت پیشفرض ۳ نیم اسپیس ساخته میشن : دیفالت ، پابلیک و سیستم نیم اسپیس

`kubectl get namespaces`

برای پروژه های معمولی نیازی به تعریف نیم اسپیس های زیاد نداریم ، همون دیفالت کافیه ، همچنین تو هر نیم اسپیس ، باید اسم پاد ها یونیک باشه

برای گرفتن تمام نیم اسپیس ها

`kubectl get all`

فکر میکنم ابتدایی ترین دستورات در ctl با نیم اسپیس ها باشه در دستور بالا بدون گفتن ns یا pod پیشفرض نیم اسپیس میده

همچنین برای ساخت pod در درون ns ها ، از دستور زیر استفاده میشه

`kubectl get svc -n <name>`

همچنین برای ایجاد یک پاد در درون یک نیم اسپیس

`kubectl run nginx --image=nginx -n ns1`

### kubeconfig

روی نود مستر قرار داره و برای مدیریت نود ها هست و موارد زیر توش هست و اطلاعات رو توش میریزیم

همچنین آدرس اون توی مستر نود :

`.kube/config`

### kubectl

یک اینترفیس است برای ارتباط با کلاستر

### kubelet

دستورات از نود مستر باید به نود ورکر برسند ، این kubelet هست که درستور از مستر میگیره و در نود خودشاجرا میکنه ، در حقیقت یه agent هست

### etcd

یک متغییر گلوبال برای پیکر بندی هست و و اطلاعات همه ی نود ها را دارد . پایگاه داده هست

### kube api server

هم ای پی آی ها رودی و خروجی رو مدیریت می کنه ، هم مانند سویچ رو نتورک ، اگه سرویس های دیگه بخوان با هم صحبت کنن ، مثل اسکجولر یا کنترلر از طریق این api server با هم صحبت می کنن ، توجه سود میشه کاستوم کرد که سرویس ها

هم مستقیم امکان پذیره

### kube scheduler

وظیفش اینکه تشخیص بده اگه یه پاد رو کدوم نود بیاد بالا با اطلاعاتی که داره

### kube conroller manager

هلت چک پاد ها رو میگیره و هر موقع یه پاد کرش کرد یا ترکید ، سریع می فهمه و به اسکجولر ، از طریق api server میگه

### objects

به مجموع deployments, replica sets, services, pods از یک واحد در کلاستر میگن

- - ReplicaSet

می توان تعداد تکرار پاد ها را با استفاده از لیبل مدیریت کرد

- - Deployment

استقرار یک پاد یا رپلیکا را مدیریت می کند

گروهی از Pods را مدیریت می‌کند و اطمینان می‌دهد که آنها در حالت دلخواه و مقیاس خودکار اجرا می‌شوند.

- minikube

یه نسخه ی سبک برای پیاده سازی روی یه ماشین ، بعد از استارت کردن مینی کیوب، به صورت پیشفرض یه پاد ایجاد میشه

### kubectl command

این دستور یک دولوپمنت می سازد

       kubectl run nginx --image=nginx

**node port**

یکی از سرویس های خوب برای تست که شبیه به `load balancer` در محیط تست عمل میکند

### رپلیکا ست - ReplicaSet

با استفاده از این می توان تعداد پاد ها رو مشخص کرد ، حتی اگه کرش کنن یا پاک شن ، این ابزار به تعداد نگاه می کنه و تعداد پاد ایجاد می کنه

برای ایجاد رپلیکا از development pod برای اسکیل آپ کردن

`kubectl get rs`

### services (network)

یک ریسورس است و در مبحث شبکه کمک میکند به این صورت:

وقتی ما پاد میسازیم یک ip داخلی بهش تخصیص داده میشه اما با کرش و ... هر بار یه پاد دیگه بیاد بالا ip تغییر میکنه ، برای حل این مشکل ما از ریسورس سرویس استفاده مکنیم ، سرویس مستقل از طول عمر ریسورس پاد ها یا دیپلویمنت هاست

به این صورت که هر سرویس یک نام مشخص یا ip

مشخص داره ، و می تونه ارتباط بگیره با یک یا چندین پاد ، دیگه در این صورت نیاز نیست مستقیم به پاد درخواست بزنیم چون یا ip پاد موقت است یا دوست داریم درخواست ها بین چند پاد پخش بشن

نود ها از طریق کارت شبکه فیزیکی خود با مستر ارتباط دارن ، و همه این ها توی یک range هستند، مستر ( احتمال زیاد ) وظیفه ی ip routing رو هم انجام میده یعنی یه تیبل تو خودش داره که میگه اگر هدف یه پاد با آی پی 10.21.10.6 هست پس next hub آی پی فیزیکی نود 18.54.96.64 هست

توجه کنید بعد از کانفیگ کیوب ، اگر ifconfig بزنیم و کارت شبکه هامون رو ببینیم احتمالا 10.42.0.1 اضافه شده باشه و میتونیم توی اون ماشین دسترسی داشته باشیم به این شبکه

نکته باحال : اگه مثلا nodeport بسازیم در حقیقت درونش cluster ip هم داریم یا اگه load balancer بسازیم ، قبلی ها رو توش داریم

انواع تایپ های سرویس

- cluster ip
  خیلی ساده ، تعریف عادیه سرویس رو داره

به این صورت که یه سرویس با تایپ کلاستر آی پی انتخاب میکنیم ، یه نام نتورک براش انتخاب میکنیم و اون دسته پاد ها با تگ و لیبل بهش میدیم ، از این به بعد اگه لیست end point های رو بگیریم ، میبینیم ip داخلی پاد هابه این نیم مپ شده

توجه شود این شبکه برای در خواست های داخلی است ، مثلن می خواییم چند تا سرویس که فقط درخواست های سرویس های داخلی رو انجام میدن ، ببینن ، اگه در خواست تو بخش نتورک به جای ip اسم این سرویس رو بدیم ، شبکه ریسالو میکنه و به یکی از پاد های شبکه میرسه ، اما بیرون از کلاستر ، در دسترس نیست

- node port

مانند clusterip هست با این ویژگی که یه پورت بهمون می ده و با اون و ip external node میشه از خارج از کلاستر ، دسترسی به اون میشه پیدا کرد

هر سرویس nodeport , یه پورت بهمون میده ، و این پورت درخواست رو تبدیل میکنه به نام شبکه

توجه کنید اسم nodeport ربطی به نود نداره و بهتره اسمش clusterPort میبود

- load balancer

تفاوت این با نودپورت در اینه که اون یه پورت میده ، اما در این تایپ ، یه اکسترنال ip میده ، اما بیشتر در محیط های cloud مانند aws و ... استفاده میشه ، اگر ما از این تایپ استفاده کنیم ezternal ip در وضعیت pending میمونه ، و باید از ابزار های third party استفاده کنیم تا بتونیم از این تایپ در لوکال استفاده کنیم

- external name

تا این جا هر چی گفتیم ، مواردی بوده که در خواست به سمت پاد یا کلاستر میومده ، با استفاده از این تایپ ، برای هر دامین یا آدرس بیرونی ، یه اسم در نظر میگیریم ،توجه شود بدون این تایپ هم

میشه به بیرون درخواست داد حالا این چه مزییتی میده ،اینکه دیگه با آدرس دقیق کار نداریم ، اگر آدرس موقت باشه هم میشه آدرس رو بروز کرد ، ولی نامی که در نظر گرفتیم ثابته و هزینه تغییر کمه

اگر از یه پاد چند رپلیکا اوردیم بالا و خواستیم با سرویس اونا رو با یک ip به دست بیاریم میشه دستور زیر رو استفاده کرد

میشه end point ها رو دید
`kubectl get ep`

### daemon set

گاهی ما نیاز داریم که روی هر نود ، یک سرویس بیاریم بالا ، توجه شود که تعداد رو مانند رپلیکا

مشخص نمی کنیم ، بلکه میگیم روی هر نود یکی بیاد بالا در این حال از دایمون ست استفاده میکنیم مثال:

یکی از سرویس های promateus که اطلاعات رو از نود جمع میکنه ، node exporter هست.
با استفاده از این kind می تونیم این سرویس رو بالا بیاریم

### دیپلوینمت

با استفاده از این می تونیم نسخه ی برنامه ها رو مدیریت کنیم و `upgrade - down grade` انجام بدیم

نحوه ی کارش اینجوریه بالا سر رپلیکا میشینه و می بینه اگه نسخه ای که بهش دستور دادن با رپلیکا ها یکی نبود ، اونا رو پاک می کنه

از ایمیج هایی که داریم ، کانتینر در پاد میسازه

هر بار دیپلوی منت جدید بدیم توی لیست و تاریخچه ی رپلیکا ست ها ، یک نسخه افزوده میشه

`kubectl get rs`

با دستور بالا می بینیم که سطر اضافه میشه

جالبه که توی دستور پایین اطلاعات خاصی نیست :

`kubectl get deploy`

### state full setp

این تضمین و میده هر استوریج برای اینستنس خاص است

وقتی از این ریسورس استفاده میکنیم ، به ازای هر اینستنس اعداد اینیتجر در انتهای نامش میزاره و به ترتیب اون ها رو میسازه و کامل مشخصه که رندوم نیست

توی این حالت میشه استوریج ها رو پرو ویژن کرد

### volumes

- persistant volume

استوریج هایی که خودمون پرو ویژن کردیم رو این جامی تونیم ببینیم

`kubectl get pv`

- persistanct volume claim

شبیه به جدول ریلیشنی در دیتابیس هست و میگه که هر اینسنتس sfs به کدوم pv ارتباط داره

#### نکات

- اگرخواستیم دیتابیس رو رپلیکا کنیم ، توجه شود حتما اسلیو ها باید پایین بیان و در نهایت مستر چون اگر برعکس باشه ، اسایو ها به محض این که ببینن مستر نیست الکشن میکنن و یکی میشه مستر

و به همین ترتیب ادامه دار میشه

- در شاردینگ دیپلویمنت به صورت دیفالت ممیفهمه که هر استوریج یا ولیوم برای کدوم شاردینگ هست ، به عبارت دیگه اگه بخواهیم چند تا اینستنس بیاریم بالا و شارد کنیم ، باید مشخص کنیم هر استوریج مال کدوم اینستنس هست زمان بالا اومدن

### hpa \_ horizental pods autoscale

یه ریسورس هست و هدفش اینه با توجه به حجم درخواست ها ، تعداد پاد ها رو تنظیم کنه ، از یه سری متریک میتونه استفاده کنه و معروف ترینس متریک cpu هست

# commands

## commons


**و چند دستور کاربردی برای کار با محیطی که devops در اختیار قرار میده**

### پاد ها رو میشه دید

`kubectl get pods`


###  مشاهده لاگ

ابتدا نیازه لیست پاد ها رو دریافت کنیم

`kubectl logs -f pod/mostafa-backend-tenancy-digitalsignature-76d465f8bd-2s8lz`


### دیپلویمنت ها رو میشه دید

`kubectl get deployments.apps`

### پایین اوردن یه سرویس در محیط

`kubectl scale deployment mostafa-backend-tenancy-upg --replicas=0`

### کلاستر تستی رو میسازه

‍`minikube start --driver=docker`

### نود ها رو میشه دید

`kubectl get nodes`

### سرویس ها رو میشه دید

`kubectl get services`

### سرویس ها رو با مینیکیوب میشه دید

`minikube service nginx`

### حذف پاد

توجه شود که این دستور در صورتی پاد رو پاک می کنه که رپلیکا ست دوباره این پاد رو نسازه

`kubectl delete pod test`

اگر دیدیم بعد از پاک کردن دوباره ایجاد شد ، باید بریم رپلیکا ست رو استاپ کنیم

### کانفیگ

`cat ~/.kube/config`

توی این فایل ۳ بخش اصلی داره

- اولیش اطلاعات کلاستره مثل آدرس سرور و ورژن و پروایدرش (مثل مینی کیوب یا k3s )

- یوزر و اطلاعاتی که هر یوزر داره

- کانتکس که ارتباط بین کلاستر ویوزر هست و این که هر یوزر به کدوم کلاستر ها دسترسی داره یا هر کلاستر چه نیم اسپیس هایی داره

#### tips

- اگه با یوزر root بخواهیم بریم دستورات رو اجرا کنیم احتمالن پیدا نکنه زیرا رو نمی ره از کانفیگ یوزر بخونه

### گرفتن ایمیج

‍`kubectl get svc nginx`

### اکسپوز کردن پورت پاد

`kubectl expose deployment nginx --type=NodePort --port=80 --name=nginx`

## روش مشاهده پاد مرحله به مرحله

- پیدا کردن IP برای تست و درخواست

‍`minikube ip`

- روش اگزپوز کردن یک پاد و پورتش

`kubectl expose pod pod-test --type=NodePort --port=8080`

- بعد این باید سرویس رو کال کرد تا ببینیم به چه پورتی مپ شده

`kubectl get svc`

- پاد ها رو میشه دید

`curl http://192.168.49.2:32593`

### متداول ترین راه ها برای دیباگ و یافت خطا

- Check Pod Status: `kubectl get pods`
- Describe Pod: `kubectl describe pod <pod-name> to check for events and errors.`
- Logs: `kubectl logs <pod-name> for application logs.`
- Exec into Pod: `kubectl exec -it <pod-name> -- /bin/bash to explore the container’s environment.`
- YAML Validation: `Ensure the Pod specification file is correct.`

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

# tips

- باید کدی میزنیم ، کلود نیتیو باشه cloud native مثال یعنی این که کد عادی شاید حالیش نشه که دیتابیس ها چند کلاستر میشن ، شاید جای نام نتورک ، آی پی بگیرن و شاید کد یا سرویس آمادگی این که چند تا پاد بیاد بالا رو نداشته باشن ، ما باید جوری کد بزنیم که سرویسامون کلود نیتیو باشن

- اگه ۳ تا نود داشته باشیم و بخوایم یه پاد رو با ۳ رپلیکاست بالا بیاریم ، به صورت دیفالت روی هر نود یه پاد میاد بالا ، ولی اگه از anti affinity rule استفاده کنیم احتمالا تضمین نکنه که پخش کنه ، همچنین اگه یه نود ریسورسش محدود باشه اون رو حذف میکنه

- همیشه در نظر داشته باشیم که اگر یه node داشته باشیم و چندین پاد مشابه ، بهتر از این است که در یک نود ، یه پاد باشه ، در نگاه اول شاید بگیم

وقتی resource ثابت هست چرا چندین پاد روی یه نود بیاریم بالا ، دلیلش موارد زیر
هست

- ریسورس ها برای کار موازی ایزوله هستند

- اگر بنا به درخواست ، یک پاد کرش کرد ، پاد های دیگه بالا هستند

- اینجوری خودمون رو ملزم میکنیم که حداقل چند پاد داریم در نتیجه معماری io خیلی تمییزه و آمادگی scale up رو داریم

اما معایبش دقیقا برعکس موارد بالا هست

- چون چندین پاد داریم ، overhead داریم و کوبرنتیس ، شبکه و لود بالانسر نیاز داره برای این

- نیاز به کانفیگ های پیچیده تر برا دقیق تر برای لیمیت گذاری هستیم

- نیاز به معماری های پیچیده برای جلو گیری از خطرات data resource share , data consistancy و manage session هست ،

# terms

- ckAd

مدرکی که کوبرنتیز به عنوان دولوپر میده

- workload

طریقه ی دیپلوی( اجرای) پاد ها

- cri

پلتفورمی هست که کوبر بسترشو فراهم کرده تا شرکت بتونن ابزاری برای مدیریت بسازن مانند docker , containerd

- cni

مانند بالایی هست ، یه اینترفیس که بستر فراهم میکنه ابزار ها برای نتورک کار کنن

- sef

شبیه به minio است ولی جانع تر ، علاوه بر ابجکت ، چیزای دیگه هاذخیره میکنن

- ترافورم

با این می تونیم سرور ها رو مدیریت کنیم، مثلکانفیگ کردن سیستم عامل

- انسیبل

بعد از پیاده سازی با ترافورم ، میشه پکیج ها و کانفیگ های پکیج ها رو با این نصب کرد

- kubeAdm \_ spraiy

حالا با استفاده از این ابزار روی ماشین ها ، کوبر مستر و کوبر اسلیو میاره بالا
