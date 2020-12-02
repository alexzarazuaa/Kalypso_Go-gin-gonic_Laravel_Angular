FROM golang:1.10

# Working Directory in the container
WORKDIR $GOPATH/src/github.com/canaz/go

# Copy everything inside the container
COPY . .

RUN go get -d -v ./...


RUN go install -v ./...


EXPOSE 3000


CMD ["fresh"] 