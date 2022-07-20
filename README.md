# Golectron

Golectron is a boilerplate for creating production-ready apps with golang and electron.

## Getting Started 

Generate the server handlers:
```
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
./cmd/server/server.proto
```

Then run:
```
go run engine/cmd
```