sudo kill -9 $(sudo lsof -t -i:8000)



python3 -m venv example_env




DATABASES = {
    'default': {
        'ENGINE': 'django.db.backends.sqlite3',
        'NAME': os.path.join(BASE_DIR, 'db.sqlite3'),
    }
}



python3 manage.py dumpdata dashboards.component > data.json

python manage.py loaddata data.json

scp root@1.2.3.4:/root/pcfilename.rar ./


////

sudo -u postgres psql

create database sla ;

create user admin with encrypted password '0231asdf';

grant all privileges on database sla to admin;

import pdb; pdb.set_trace()
