python3 -m venv example_env


DATABASES = { 'default': { 'ENGINE': 'django.db.backends.sqlite3', 'NAME': os.path.join(BASE_DIR, 'db.sqlite3'), } }

python3 manage.py dumpdata dashboards.component > data.json

python manage.py loaddata data.json

import pdb; pdb.set_trace()

