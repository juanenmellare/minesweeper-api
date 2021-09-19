FROM postgres

COPY /install-extensions.sql /docker-entrypoint-initdb.d/
