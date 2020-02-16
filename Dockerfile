FROM golang:latest
MAINTAINER "yeqiongzhou@whu.edu.cn"
WORKDIR /go/src/http-fake
ADD . /go/src/http-fake
RUN go build .
EXPOSE 2048
ENTRYPOINT ["./http-fake"]
