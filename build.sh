#!/usr/bin/env sh
set -e

if [ "$1" == "--help" -o "$1" == "-h" ]; then
    echo -e "\033[32m
  Build Applicaton.
  ----------------------------
  Usage:
    - ./build.sh [os (linux | darwin)]    Compile the package for the given platform. default darwin
    - ./build.sh --help                   Show this help
  Help:
    - Get more help see : https://dev.m-io.cn/help/build?t=gin
    "
    exit 1
fi

OS="linux"
if [ "$1" != "" ]; then
    OS=$1
fi

echo "[$(date '+%Y-%m-%d %H:%M:%S')] Start build for $OS platform."

go mod tidy

export CGO_ENABLED=0
export GOOS=$OS
export GOARCH=amd64

v="web-srv.v.`date +"%Y%m%d%H%M"`"
echo "[$(date '+%Y-%m-%d %H:%M:%S')] Building application with name: $v"

cmd="go build -tags=jsoniter -a -installsuffix cgo -ldflags '-s -w' -o dist/$v ."
echo "[$(date '+%Y-%m-%d %H:%M:%S')] Building command: $cmd"

go build -tags=jsoniter -a -installsuffix cgo -ldflags '-s -w' -o dist/$v .

if [ "$?" == "0" ]; then
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] Successful."
    # shellcheck disable=SC2225
    cp -f "dist/$v" "dist/srv"
    echo "[$(date '+%Y-%m-%d %H:%M:%S')] Target: dist/$v -> dist/srv"
else
    echo echo "[$(date '+%Y-%m-%d %H:%M:%S')] Fail."
fi
