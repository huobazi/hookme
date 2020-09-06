#!/bin/bash

self_location_dir=$(
  cd $(dirname $0)
  pwd
)

. $self_location_dir/_var.sh

declare -a PLATFORMS=(
  "darwin/386"
  "darwin/amd64"
#  "darwin/arm"
#  "darwin/arm64"
  "freebsd/386"
  "freebsd/amd64"
  "freebsd/arm"
#  "freebsd/arm64"
  "linux/386"
  "linux/amd64"
  "linux/arm"
  "linux/arm64"
  "netbsd/386"
  "netbsd/amd64"
  "netbsd/arm"
  "netbsd/arm64"
  "openbsd/386"
  "openbsd/amd64"
  "openbsd/arm"
  "openbsd/arm64"
  "solaris/amd64"
  "windows/386"
  "windows/amd64"
  "windows/arm"
)

for PLATFORM in "${PLATFORMS[@]}"
do
    PLATFORM_SPLIT=(${PLATFORM//\// })
    GOOS=${PLATFORM_SPLIT[0]}
    GOARCH=${PLATFORM_SPLIT[1]}

    OUTPUT_NAME=$DIST/$BINARY-$GOOS-$GOARCH/$BINARY
    if [ $GOOS = "windows" ]; then
        OUTPUT_NAME+='.exe'
    fi
    echo "Building $GOOS-$GOARCH ..."
    mkdir -p $DIST/$BINARY-$GOOS-$GOARCH/
    GOOS=$GOOS GOARCH=$GOARCH go build -o $OUTPUT_NAME -ldflags "$LDFLAGS" $BUILD_FILE
    tar cz -C $DIST -f $DIST/$BINARY-$GOOS-$GOARCH.tar.gz $BINARY-$GOOS-$GOARCH
done

echo "Done."