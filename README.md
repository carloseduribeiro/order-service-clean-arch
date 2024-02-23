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
