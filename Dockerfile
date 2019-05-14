FROM golang:alpine

ADD ./src /go/src/app
WORKDIR /go/src/app

ENV PORT=3001

RUN apk add git

RUN go get -u github.com/gorilla/mux
RUN go get github.com/sirupsen/logrus 
RUN go get github.com/pilu/fresh

CMD ["go", "run", "main.go"]
