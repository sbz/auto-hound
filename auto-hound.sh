#!/bin/sh

cd /
[ ! -f /config.json ] && /auto-hound > /config.json
/go/bin/houndd -conf /config.json
