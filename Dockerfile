FROM golang:alpine

ADD ./dive-into-go /go/src/github.com/dykyi-roman/dive-into-go
WORKDIR /go/src/github.com/dykyi-roman/dive-into-go

ENV PORT=3001

RUN apk add git

RUN go get -u github.com/gorilla/mux
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/sirupsen/logrus 
RUN go get github.com/pilu/fresh
RUN go get -u github.com/go-redis/redis
RUN go get -u github.com/derekparker/delve/cmd/dlv

CMD ["dlv", "debug", "go", "run", "main.go"]
