FROM golang:alpine AS build-env
# Working Directory in the container
WORKDIR $GOPATH/src/github.com/canaz/Kalypso_Go-gin-gonic_Laravel_Angular

# Copy everything inside the container
COPY . .

RUN echo "fs.inotify.max_user_watches=524288" | tee -a /etc/sysctl.conf | sysctl -p
RUN set -eux; \
    apk add --no-cache --virtual .build-deps \
    git gcc libc-dev; \
    go get -u github.com/kardianos/govendor; \
    go get -u github.com/pilu/fresh; \
    go get -u github.com/gin-gonic/gin; \    
    go get -u github.com/jinzhu/gorm; \
    go get -u github.com/go-redis/redis; \
    go get -u golang.org/x/crypto/bcrypt; \
    go get -u github.com/danilopolani/gocialite; \
    go get -u github.com/jinzhu/gorm/dialects/sqlite; \
    go get -u gopkg.in/go-playground/validator.v8; \
    go get -u github.com/gin-gonic/gin/binding; \
    go get -u github.com/gosimple/slug; \	
    go get -u github.com/dgrijalva/jwt-go;
ENV GOPATH /go
ENV GOROOT /usr/local/go