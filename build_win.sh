
#1、安装 mingw-w64
export CGO_ENABLED=1
export GOOS=windows
export GOARCH=amd64
export CC=x86_64-w64-mingw32-gcc
export CXX=x86_64-w64-mingw32-g++
VERSION=0.0.7
BUILD_DATE=$(date +'%Y-%m-%d')
BUILD_TIME=$(date +'%H:%M:%S')
BUILD_GO_VERSION=$(go version)

wails build -platform windows  \
    -ldflags " -s -w -H=windowsgui \
                 -X 'main.version=$VERSION' \
                 -X 'main.buildTime=$BUILD_TIME' \
                 -X 'main.buildDate=$BUILD_DATE' \
                 -X 'main.buildGoVersion=${BUILD_GO_VERSION}'" \
    -webview2 embed -upx
# 打包
