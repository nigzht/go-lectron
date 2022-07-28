# Golectron

Golectron is a boilerplate for creating medium-to-large-sized production-ready apps with golang and electron.

# Features

- [ ] RPC Communication(gRPC)
- [ ] App Security
- [ ] Build processes
- [ ] Autoupdater
- [ ] Code signing
- [ ] Hot reloading

# Design

Go-lectron uses [gRPC](https://grpc.io/) to communicate between processes. It is designed to be extensible and easy to build in a monorepo format.

## Getting Started

Firstly, you will need the protoc tool for generating definitions. You can find installation instructions [here.](https://grpc.io/docs/protoc-installation/)

You will also need to install the protobufjs-cli tool for converting proto files to json:
```
npm install -g protobufjs-cli --save-prefix=~ 
```

Next run the bootstrap script.
```
./scripts/bootstrap.sh
```

All done! you can now run the project by either running the app or the standalone service:
```
# the app, which will launch the built service
cd service/cmd; go build -o ../../build; cd ..
cd app; yarn start

# or the service

go run service/cmd
```

