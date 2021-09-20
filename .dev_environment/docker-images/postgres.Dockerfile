FROM postgres

COPY /.dev_environment/docker-images/install-extensions.sql /docker-entrypoint-initdb.d/
