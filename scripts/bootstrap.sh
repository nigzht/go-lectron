#!/bin/sh

DIR="${BASH_SOURCE%/*}"
if [[ ! -d "$DIR" ]]; then DIR="$PWD"; fi
. "$DIR/protos.sh"

# install dependencies.
cd app; yarn install; cd ..
cd service; go get all; cd ..

# install the protoc gen go binary.
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# initialize go workspace.
go work init ./service

# set the path for protoc to find the go binary.
export PATH="$PATH:$(go env GOPATH)/bin"

# generate files for proto.
"$DIR/protos.sh"