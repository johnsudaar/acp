FROM golang:1.14
MAINTAINER jonathan.hurter@gmail.com

RUN go get github.com/cespare/reflex
EXPOSE 8042

ADD . /go/src/github.com/johnsudaar/acp
WORKDIR /go/src/github.com/johnsudaar/acp

ENV GOFLAGS -mod=vendor

RUN go build

CMD ["./acp", "start"]
