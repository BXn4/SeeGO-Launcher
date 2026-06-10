#!/bin/bash
source .env

GOOS=windows GOARCH=amd64 go build \
  -ldflags "-X 'seegolauncher/internal/services.OA=${OA}'" \
  -o ./bin/seego-launcher.exe
