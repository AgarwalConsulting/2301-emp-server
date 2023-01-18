#!/usr/bin/env bash

docker build -t emp-server .

docker run -it --rm -e PORT=9000 -p 8000:9000 emp-server
