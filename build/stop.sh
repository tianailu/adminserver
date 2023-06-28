#!/bin/bash
set -e
docker stop tal_1
docker stop master
docker stop node1
docker stop node2
docker rm tal_1
docker rm master
docker rm node1
docker rm node2
