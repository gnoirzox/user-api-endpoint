FROM mysql:5.7

EXPOSE 3306

VOLUME ["/var/lib/mysql"]

ADD tables.sql /docker-entrypoint-initdb.d
