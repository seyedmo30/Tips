
pg_dump -h 192.168.13.29 -U amir_sh  -Fc -f /var/backups/db/tootia-$(date +%Y-%m-%d).sql reports_csf








The most simple case is dumping and restoring on the same server:
$ pg_dump -h localhost -Fc test > /home/postgres/dump.sql

$ pg_restore -h localhost test < /home/postgres/dump.sql
 

Or with a plain text dump:
$ pg_dump -h localhost -f /home/postgres/dump.sql test

$ psql -h localhost -f /home/postgres/dump.sql test
 

Where this gets interesting is with multiple hosts. You can:
$ # dump a remote database to your local machine

$ pg_dump -h remotedb.mydomain.com -f /home/postgres/dump.sql test


$ # dump a local database and write to a remote machine

$ pg_dump -h remotedb.mydomain.com test | ssh postgres@remotedb.mydomain.com 'cat > dump.sql'


$ # dump a remote database and write to the same remote machine

$ pg_dump -h remotedb.mydomain.com test | ssh postgres@remotedb.mydomain.com 'cat > dump.sql'


$ # or a different remote machine

$ pg_dump -h remotedb1.mydomain.com test | ssh postgres@remotedb2.mydomain.com 'cat > dump.sql'
 

You also have similar restore options. I will use psql below but pg_restore works the same:
$ # dump a remote database and restore to your local machine

$ pg_dump -h remotedb.mydomain.com test1 | psql test2


$ # dump a local database and restore to a remote machine

$ pg_dump -h remotedb.mydomain.com test | ssh postgres@remotedb.mydomain.com 'psql test'


$ # dump a remote database and restore to the same remote machine

$ pg_dump -h remotedb.mydomain.com test1 | ssh postgres@remotedb.mydomain.com 'psql test2'


$ # or a different remote machine

$ pg_dump -h remotedb1.mydomain.com test | ssh postgres@remotedb2.mydomain.com 'psql test'
