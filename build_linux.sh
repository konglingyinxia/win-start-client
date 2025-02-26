
VERSION=0.0.7
BUILD_DATE=$(date +'%Y-%m-%d')
BUILD_TIME=$(date +'%H:%M:%S')
BUILD_GO_VERSION=$(go version)

wails build -platform linux  \
  -ldflags " \
        -X 'main.version=$VERSION' \
        -X 'main.buildTime=$BUILD_TIME' \
        -X 'main.buildDate=$BUILD_DATE' \
        -X 'main.buildGoVersion=${BUILD_GO_VERSION}'"
