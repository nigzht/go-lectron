#!/bin/sh

# generated files for go process.
protoc --go_out=. --go-grpc_out=. ./protos/*.proto

# generated files for node process.
cp ./protos/*.proto ./app/src/service/protos/
./app/node_modules/.bin/proto-loader-gen-types --longs=String --enums=String -defaults- --oneofs --grpcLib=@grpc/grpc-js --outDir=./app/src/service/protos/generated ./protos/*.proto