#!/bin/bash -e

echo "building go ..."
GOOS=linux GOARCH=amd64 go build main.go
echo "build successful!"

echo "creating dist.zip"
zip dist.zip index.js main
