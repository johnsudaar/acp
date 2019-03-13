FROM golang:1.11
MAINTAINER jonathan.hurter@gmail.com

RUN go get github.com/cespare/reflex
EXPOSE 8042

ADD . /go/src/github.com/johnsudaar/acp
WORKDIR /go/src/github.com/johnsudaar/acp

RUN go build

CMD ["./acp"]
