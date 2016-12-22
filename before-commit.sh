#!/bin/bash

gofmt -s -w .
golint ./...
go vet ./...
go get github.com/gorilla/mux
go test ./... -v -cover -race
export ANDROID_HOME=$HOME"/android-sdk"
gomobile bind --target=android .
gomobile bind --target=ios .
