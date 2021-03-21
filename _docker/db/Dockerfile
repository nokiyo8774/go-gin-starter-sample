FROM mysql:5.7

COPY ./_docker/db/my.cnf /etc/mysql/conf.d
COPY ./db/migrate/01_create_schema.sql /docker-entrypoint-initdb.d
COPY ./db/migrate/02_seed.sql /docker-entrypoint-initdb.d