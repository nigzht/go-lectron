#!/bin/sh

DIR="${BASH_SOURCE%/*}"
if [[ ! -d "$DIR" ]]; then DIR="$PWD"; fi
. "$DIR/protos.sh"

# install dependencies.
cd app; yarn install; cd ..
cd service; go get all; cd ..

# initialize go workspace.
go work init ./service

# set the path for protoc.
export PATH="$PATH:$(go env GOWORK)/bin"

# generate files for proto.
"$DIR/protos.sh"