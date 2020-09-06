#!/bin/bash
set -e

self_location_dir=$(cd `dirname $0`; pwd)

. $self_location_dir/_var.sh

echo "Build hookme with ldflags: $LDFLAGS"
go build  -o $DIST/$BINARY -ldflags "$LDFLAGS" $BUILD_FILE