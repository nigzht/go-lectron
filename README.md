# Golectron

Golectron is a boilerplate for creating production-ready apps with golang and electron. 

# Features

- [ ] Secure Communication
- [ ] App Security
- [ ] Build processes
- [ ] Autoupdater
- [ ] Code signing


# Design

Golectron is only meant for single backend service implementations. We try to keep the tools used as agnostic as possible, although some decisions are made -- namely, communication between the go and node process uses gRPC, electron-forge is used for building and the node process uses typescript. The frontend implementation is left up to the user but there are forks that use react and next.

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