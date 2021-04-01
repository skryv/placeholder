#!/usr/bin/env bash

docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.16.2 go build -v

docker build -t dckr.skryv.com/hello-world .
