#!/usr/bin/bash

APP_NAME="njav-downloader"

# go clean -cache
sh -c "go build \
-tags=en \
-ldflags=\"\
-X 'main.Version=v1.0.1' \
-X 'main.AppName=$APP_NAME' \
-X 'main.BuildSrc=$PWD'\" \
-o ./dist/$APP_NAME main.go"