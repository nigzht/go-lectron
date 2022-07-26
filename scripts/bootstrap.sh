# install dependencies.
cd app; yarn install; cd ..
cd service; go get all; cd ..

# initialize go workspace.
go work init ./service

# set the path for protoc.
export PATH="$PATH:$(go env GOWORK)/bin"

# generate proto files for node and go processes.
protoc --go_out=. \
--go-grpc_out=. \
./proto/fmt.proto

./app/node_modules/.bin/proto-loader-gen-types --longs=String --enums=String -defaults- --oneofs --grpcLib=@grpc/grpc-js --outDir=./app/src/service/proto/ ./proto/*.proto