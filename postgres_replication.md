\c db

wal_level = logical

listen_addressess = "*" '*'


pg_hba.conf

host    sampledb    replicauser     192.168.10.10/32   scram-sha-256


CREATE PUBLICATION bookpub FOR TABLE book;


CREATE SUBSCRIPTION booksub CONNECTION 'dbname=sampledb host=192.168.10.5 user=replicauser password=12345 port=5432' PUBLICATION bookpub



select * from pg_catalog.pg_publication;  

\dRp+




select * from pg_catalog.pg_subscription;

\dRs+



select * from pg_replication_slots ;



ALTER SUBSCRIPTION booksub DISABLE;
ALTER SUBSCRIPTION booksub SET (slot_name=NONE);
DROP SUBSCRIPTION booksub;


DROP publication bookpub ;
