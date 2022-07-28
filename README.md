# Golectron

Golectron is a boilerplate for creating medium-to-large-sized production-ready apps with golang and electron.

# Features

- [x] RPC Communication(gRPC)
- [ ] App Security
- [ ] Build processes
- [ ] Autoupdater
- [ ] Code signing
- [x] Hot reloading

# Design

Go-lectron uses [gRPC](https://grpc.io/) to communicate between processes. It is designed to be extensible and easy to build in a monorepo format. It supports mac and windows out of the box. Linux users [click here.](https://i.kym-cdn.com/entries/icons/original/000/035/699/pepe.jpg)

## Getting Started

Firstly, you will need the protoc tool for generating definitions. You can find installation instructions [here.](https://grpc.io/docs/protoc-installation/)

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

## Development

pls run your github actions changes locally. its faster and doesn't cost money. You can do so with [act.](https://github.com/nektos/act)