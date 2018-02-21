#!/bin/sh

cd /
[ ! -f /config.json ] && go run /glr.go > /config.json
/go/bin/houndd -conf /config.json
