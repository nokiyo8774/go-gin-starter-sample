FROM golang:latest

COPY ./_docker/app/docker-entrypoint.sh /usr/local/bin/
RUN chmod 744 /usr/local/bin/docker-entrypoint.sh

ENV SRC_DIR=$GOPATH/src/app
RUN mkdir -p $SRC_DIR
WORKDIR $SRC_DIR
ADD ./app $SRC_DIR
ENTRYPOINT ["docker-entrypoint.sh"]