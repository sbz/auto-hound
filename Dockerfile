FROM alpine:latest
MAINTAINER sbz@6dev.net

ENV GOPATH /go

COPY ./auto-hound.sh /auto-hound.sh
COPY ./auto-hound /auto-hound

RUN apk update
RUN apk add go git libc-dev
RUN cd /
RUN go get github.com/etsy/hound/cmds/houndd
RUN go get github.com/etsy/hound/cmds/hound

EXPOSE 6080

ENTRYPOINT ["/auto-hound.sh"]
