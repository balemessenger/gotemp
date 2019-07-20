#!/usr/bin/env bash

if  [[ "$1" == "" ]]; then
	go test -count=1 ./pkg/
	go test -count=1 ./internal/
	go test -count=1 ./api/

else
    go test -count=1 -v -run $1 ./pkg/
	go test -count=1 -v -run $1 ./internal/
	go test -count=1 -v -run $1 ./api/
fi
