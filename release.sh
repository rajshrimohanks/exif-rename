#!/usr/bin/env bash
rm -rf release/

VER=$1
if [ -z "$1" ]; then
  VER="v0.0"
fi

echo "Building version: $VER"
export GOARCH=amd64

export GOOS=windows
go build -v -o release/exif-rename_${VER}_${GOOS}-${GOARCH}.exe .

export GOOS=linux
go build -v -o release/exif-rename_${VER}_${GOOS}-${GOARCH} .

export GOOS=darwin
go build -v -o release/exif-rename_${VER}_${GOOS}-${GOARCH} .
