# Nemesis

## Commands

go run engine/cmd

nemesis server --port=6969

protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
./cmd/server/server.proto