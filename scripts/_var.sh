#!/bin/bash

set -e
readonly BASEDIR=$PWD
readonly BINARY=hookme
readonly DIST=dist
readonly SHELL=/bin/bash
readonly LDFLAGS=$(govvv -flags)
readonly BUILD_FILE=cmd/hookme/main.go