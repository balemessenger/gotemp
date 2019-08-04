#!/usr/bin/env bash

if  [[ "$1" == "" ]]; then
	go test -count=1 ./test/

else
    go test -count=1 -v -run $1 ./test/
fi
