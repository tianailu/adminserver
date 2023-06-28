#!/bin/bash

go build -ldflags "-w -s" -race -v -o tal ../cmd/main/main.go

cp -R ../config /data/tal/
cp ./tal /data/tal/
chmod +x /data/tal/tal
docker build -t tal:1.0.0 .
