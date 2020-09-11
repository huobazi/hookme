#!/bin/bash

set -e
readonly BASEDIR=$PWD
readonly BINARY=hookme
readonly DIST=dist
readonly SHELL=/bin/bash
readonly LDFLAGS=$(govvv -flags -pkg $(go list ./internal/constants))
readonly BUILD_FILE=main.go