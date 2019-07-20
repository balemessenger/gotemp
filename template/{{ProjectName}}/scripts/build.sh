#!/usr/bin/env bash

VERSION=`${PWD}/scripts/version.sh`
TIME=$(date)

if [ -d build/release ]; then
    mkdir -p build/release
    mkdir -p build/debug
fi

if [ "$1" == "release" ]; then
    echo "Building in release mode"
    go build -o build/release/{{ProjectName}}_$VERSION -a -installsuffix cgo -ldflags="-X '{{ProjectName}}/version.BuildTime=$TIME' -X '{{ProjectName}}/version.BuildVersion=$VERSION' -s" main.go
else
    echo "Building in debug mode"
    go build -o build/debug/{{ProjectName}}_$VERSION -a -v -installsuffix cgo -ldflags="-X '{{ProjectName}}/version.BuildTime=$TIME' -X '{{ProjectName}}/version.BuildVersion=$VERSION' -s" main.go
fi