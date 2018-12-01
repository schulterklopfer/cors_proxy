#!/bin/bash

GOOS=linux   GOARCH=amd64       go build -o dist/cors_proxy.linux-amd64   -ldflags="-s -w"
GOOS=linux   GOARCH=arm GOARM=5 go build -o dist/cors_proxy.linux-arm5    -ldflags="-s -w"
GOOS=linux   GOARCH=arm GOARM=6 go build -o dist/cors_proxy.linux-arm6    -ldflags="-s -w"
GOOS=linux   GOARCH=arm GOARM=7 go build -o dist/cors_proxy.linux-arm7    -ldflags="-s -w"
GOOS=linux   GOARCH=arm64       go build -o dist/cors_proxy.linux-arm8    -ldflags="-s -w"
GOOS=darwin  GOARCH=amd64       go build -o dist/cors_proxy.darwin-amd64  -ldflags="-s -w"
GOOS=windows GOARCH=amd64       go build -o dist/cors_proxy.windows-amd64 -ldflags="-s -w"
GOOS=windows GOARCH=amd64       go build -o dist/cors_proxy.windows-amd64 -ldflags="-s -w"
