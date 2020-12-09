FROM golang:alpine AS build-env

RUN apk update && apk add git && go get gopkg.in/natefinch/lumberjack.v2

ENV GOROOT=/usr/local/go
ENV PATH=$GOPATH/bin:$GOROOT/bin:$PATH
ENV GOPATH=/home/usuario/Documentos/2nDaw/Servidor/Projects_go/
ENV PATH=$PATH:/home/usuario/Documentos/2nDaw/Servidor/Projects_go/bin



RUN echo "fs.inotify.max_user_watches=524288" | tee -a /etc/sysctl.conf | sysctl -p
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/kardianos/govendor
RUN go get -u github.com/pilu/fresh
RUN go get github.com/go-sql-driver/mysql
RUN  go get -u github.com/go-redis/redis
