# order-service-clean-arch

- **English Version: [README.md](README.md)**

Este repositório contém um projeto desenvolvido durante o módulo Clean Architecture do curso GoExpert da FullCycle.

No arquivo [CleanArchitectureReviewGuide_ptBR.md](./assets/CleanArchitectureReviewGuide_ptBR), consta um resumo foi
feito para revisar os conceitos da Clean Architecture. Ele abrange os assuntos abordados no curso com mais informações
que julguei ser necessário adicionar.

## Descrição

order-service-clean-arch uma aplicação desenvolvida utilizando os princípios de design da Arquitura Limpa (Clean
Architecture). Ela expõe três serviços Web:

* **HTTP Server**: porta ```80```
* **gRPC**: porta ```50051```
* **graphQL**: porta ``4000``

### HTTP Server

A aplicação possui dois endpoints:

1. Criação de um novo pedido: ```POST /order```;
2. Listagem de pedidos: ```GET /order```.

Ambas as chamadas estão definidas na pasta [api](./api) na raiz do projeto.

### gRPC

Nós estamos utilizando a implementação [gRPC-go](https://pkg.go.dev/google.golang.org/grpc) do [gRPC](https://grpc.io/)
comunicação via RPC com [Protocol Buffers 3](https://protobuf.dev/programming-guides/proto3/).

Você pode utilizar o [Protocol Buffer Compiler](https://grpc.io/docs/protoc-installation/) para analisar e compilar o
arquivo ```.proto```, que contém as definições dos serviços e mensagens. Veja
o [Quick Start](https://grpc.io/docs/languages/go/quickstart/#prerequisites) guide para mais informações.

Se você quiser fazer alterações no projeto, você pode compilar o arquivo ```.proto``` com o seguinte comando:

```shell
protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto
```

Você pode instalar o **[Evans gRPC Client](https://github.com/ktr0731/evans)** para realizar as RPCs. Ele lista os
serviços disponibilizados pelo nosso servidor gRPC de forma amigável.

### GraphQL

## Configuração

## Construção e Execução

## Testes Unitários e Cobertura de Testes

