# cf https://hub.docker.com/_/alpine
FROM alpine:3.9

LABEL description="Docker image for auto-hound" \
      maintainer="sbz@6dev.net" \
      repository="https://github.com/sbz/auto-hound"

ENV GOPATH /go

COPY ./auto-hound.sh /auto-hound.sh
COPY ./auto-hound /auto-hound

WORKDIR /

RUN apk add --update go git libc-dev \
    && rm -rf /var/cache/apk/*
RUN go get github.com/etsy/hound/cmds/houndd
RUN go get github.com/etsy/hound/cmds/hound

EXPOSE 6080

ENTRYPOINT ["/auto-hound.sh"]
