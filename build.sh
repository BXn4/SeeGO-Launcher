#!/bin/bash
source .env

wails3 task common:update:build-assets
sed -i "s/^var OA string$/var TEBEXACCOUNT string = \"${OA}\"/" internal/services/api.go
wails3 package GOOS=windows
sed -i "s/^var OA string = \"${OA}\"$/var TEBEXACCOUNT string/" internal/services/api.go
