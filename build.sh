#!/bin/bash
source .env

wails3 task common:update:build-assets
sed -i "s/^var OA string$/var OA string = \"${OA}\"/" internal/services/api.go
wails3 package GOOS=windows
sed -i "s/^var OA string = \"${OA}\"$/var OA string/" internal/services/api.go
