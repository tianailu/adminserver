#!/bin/sh

docker save docker-hub.cloud.top/srp-av/edr-exp:1.0.0 | gzip > edr-exp_1.0.0.tar.gz
