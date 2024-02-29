# order-service-clean-arch

* **Portuguese Version: [README_ptBR.md](README_ptBR.md)**

This repository contains a project developed during the Clean Architecture module of the GoExpert course by FullCycle.

## How to Execute de project

```shell
# Run the database and RabbitMQ server in docker:
docker compose up -d

# Run the application:
cd cmd/orderSystem && go run main.go wire_gen.go
```

## At this project

### HTTP Server

### gRPC

We are using [gRPC-go](https://pkg.go.dev/google.golang.org/grpc) implementation of [gRPC](https://grpc.io/) for
communication by RPC with [Protocol Buffers 3](https://protobuf.dev/programming-guides/proto3/).

You can use the [Protocol Buffer Compiler](https://grpc.io/docs/protoc-installation/) to parse and compile
the ```.proto``` file, witch contain service and message definitions. See
gRPC [Quick Start](https://grpc.io/docs/languages/go/quickstart/#prerequisites) guide for more information.

Parsing and compiling ```.proto``` file:

```shell
protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto
```

### GraphQL

### google-wire

We are using the [dependency injection](https://stackoverflow.com/questions/130794/what-is-dependency-injection) design
principle. In practice, that means we pass in whatever each component needs. This style of design lends itself to
writing easily tested code and makes it easy to swap out one dependency with another.

One downside to dependency injection is the need for so many initialization steps. Wire make's the process of
initializing our components smoother.

See Wire's docs: https://github.com/google/wire/tree/main

The file ```cmd/orderSystem/wire.go``` contains the initializer definitions to generate Wire's code.
