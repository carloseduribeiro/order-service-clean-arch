package main

import (
	"database/sql"
	"fmt"
	"io"
	"net"
	"net/http"

	graphqlhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/streadway/amqp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/go-sql-driver/mysql"

	"github.com/carloseduribeiro/order-service-clean-arch/configs"
	"github.com/carloseduribeiro/order-service-clean-arch/internal/domain/event/handler"
	"github.com/carloseduribeiro/order-service-clean-arch/internal/infra/graph"
	"github.com/carloseduribeiro/order-service-clean-arch/internal/infra/grpc/pb"
	"github.com/carloseduribeiro/order-service-clean-arch/internal/infra/grpc/service"
	"github.com/carloseduribeiro/order-service-clean-arch/internal/infra/web/webserver"
	"github.com/carloseduribeiro/order-service-clean-arch/pkg/events"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		config.DBUser, config.DBPassword, config.DBHost, config.DBPort, config.DBName,
	)
	db, err := sql.Open(config.DBDriver, dataSourceName)
	if err != nil {
		panic(err)
	}
	defer closeWithPanic(db)

	rabbitMQChannel := getRabbitMQChannel()
	eventDispatcher := events.NewEventDispatcher()
	err = eventDispatcher.Register("OrderCreated", &handler.OrderCreatedHandler{RabbitMQChannel: rabbitMQChannel})
	if err != nil {
		panic(err)
	}

	webServer := webserver.NewWebServer(config.WebServerPort)
	createOrderHttpHandler := initializeCreateOrderHttpHandler(db, eventDispatcher)
	if err = webServer.AddHandler(http.MethodPost, "/order", createOrderHttpHandler.Create); err != nil {
		panic(err)
	}
	listOrdersHttpHandler := initializeListOrderHttpHandler(db)
	if err = webServer.AddHandler(http.MethodGet, "/order", listOrdersHttpHandler.List); err != nil {
		panic(err)
	}
	fmt.Println("Starting web server on port", config.WebServerPort)
	go webServer.Start()

	grpcServer := grpc.NewServer()
	createOrderUseCase := initializeCreateOrderUseCase(db, eventDispatcher)
	listOrdersUseCase := initializeListOrdersUseCase(db)
	createOrderService := service.NewOrderService(*createOrderUseCase, *listOrdersUseCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", config.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", config.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	go grpcServer.Serve(lis)

	srv := graphqlhandler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CreateOrderUseCase: *createOrderUseCase,
	}}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	fmt.Println("Starting GraphQL server on port", config.GraphQLServerPort)
	http.ListenAndServe(":"+config.GraphQLServerPort, nil)
}

func getRabbitMQChannel() *amqp.Channel {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		panic(err)
	}
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	return ch
}

func closeWithPanic(c io.Closer) {
	if c == nil {
		panic("Closer is nil")
	}
	if err := c.Close(); err != nil {
		panic(err)
	}
}
