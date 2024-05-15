# NETWORK

## nat

اگر یک شبکه ی پرایویت داشته باشیم و همه ی آن ها از طریق یک IP static به اینترنت وصل شوند ، در این صورت از نت استفاده کردیم .


به منظور ایجاد یک NAT (Network Address Translation) برای ترافیک شبکه، می‌توان از دستور `iptables` در لینوکس استفاده کرد. این دستور به شما امکان می‌دهد تا
 قوانین ترافیک شبکه را تعریف و مدیریت کنید.
 برای پیکربندی NAT، می‌توان از دستورات زیر استفاده کرد:
 + فعال کردن NAT
 iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
 + ترافیک ورودی به پورت 80 را به پورت 8000 هدایت کن
 iptables -t nat -A PREROUTING -i eth0 -p tcp --dport 80 -j REDIRECT --to-ports 8000

