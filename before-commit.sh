#!/bin/bash

go fmt ./...
golint ./...
go vet ./...
go get github.com/gorilla/mux
go test ./... -v -cover
export ANDROID_HOME=$HOME"/android-sdk"
gomobile bind --target=android .
gomobile bind --target=ios .
