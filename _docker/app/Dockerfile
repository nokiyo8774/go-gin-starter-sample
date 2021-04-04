FROM golang:1.16.2

COPY ./_docker/app/docker-entrypoint.sh /usr/local/bin/
RUN chmod 744 /usr/local/bin/docker-entrypoint.sh

RUN mkdir -p /var/log/app

ENV SRC_DIR=$GOPATH/src/app
RUN mkdir -p $SRC_DIR
WORKDIR $SRC_DIR
ADD ./app $SRC_DIR
ENTRYPOINT ["docker-entrypoint.sh"]