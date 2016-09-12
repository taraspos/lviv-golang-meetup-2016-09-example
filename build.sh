#!/bin/bash
export GO15VENDOREXPERIMENT=1 \
  && export GOPATH="$(pwd)" \
  && go get github.com/Masterminds/glide \
  && mkdir -p $GOPATH/bin \
  && mkdir -p $GOPATH/src/app \
  && mv *.go $GOPATH/src/app/ \
  && mv vendor/ $GOPATH/src/app/ \
  && cd $GOPATH/src/app/ \
  && go build -a -o $GOPATH/bin/application
