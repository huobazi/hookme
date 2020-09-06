#!/bin/bash
set -e

self_location_dir=$(cd `dirname $0`; pwd)

. $self_location_dir/_var.sh

echo "Clean..."
  rm -rf $DIST