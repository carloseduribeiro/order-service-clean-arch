//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"
	"github.com/google/wire"

	"github.com/carloseduribeiro/order-service-clean-arch/internal/application/usecase"
	"github.com/carloseduribeiro/order-service-clean-arch/internal/domain/entity"
	"github.com/carloseduribeiro/order-service-clean-arch/internal/domain/event"
	"github.com/carloseduribeiro/order-service-clean-arch/internal/infra/database"
	"github.com/carloseduribeiro/order-service-clean-arch/internal/infra/web"
	"github.com/carloseduribeiro/order-service-clean-arch/pkg/events"
)

var setOrderRepositoryDependency = wire.NewSet(
	database.NewOrderRepository,
	wire.Bind(new(entity.OrderRepositoryInterface), new(*database.OrderRepository)),
)

var setEventDispatcherDependency = wire.NewSet(
	events.NewEventDispatcher,
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
	wire.Bind(new(events.EventDispatcherInterface), new(*events.EventDispatcher)),
)

var setOrderCreatedEvent = wire.NewSet(
	event.NewOrderCreated,
	wire.Bind(new(events.EventInterface), new(*event.OrderCreated)),
)

func initializeCreateOrderUseCase(
	db *sql.DB, eventDispatcher events.EventDispatcherInterface,
) *usecase.CreateOrderUseCase {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		usecase.NewCreateOrderUseCase,
	)
	return &usecase.CreateOrderUseCase{}
}

func initializeCreateOrderHttpHandler(
	db *sql.DB, eventDispatcher events.EventDispatcherInterface,
) *web.CreateOrderHttpHandler {
	wire.Build(
		setOrderRepositoryDependency,
		setOrderCreatedEvent,
		web.NewWebOrderHandler,
	)
	return &web.CreateOrderHttpHandler{}
}

func initializeListOrdersUseCase(sb *sql.DB) *usecase.ListOrdersUseCase {
	wire.Build(setOrderRepositoryDependency, usecase.NewListOrdersUseCase)
	return &usecase.ListOrdersUseCase{}
}

func initializeListOrderHttpHandler(db *sql.DB) *web.ListOrdersHandler {
	wire.Build(setOrderRepositoryDependency, web.NewListOrdersHandler)
	return &web.ListOrdersHandler{}
}
