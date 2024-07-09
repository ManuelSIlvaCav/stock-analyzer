package data_consumer

import (
	"stockanalyzer/internal/container"
	data_consumer_handlers "stockanalyzer/internal/data_consumer/handlers"
	data_consumer_providers "stockanalyzer/internal/data_consumer/providers"
	data_consumer_repository "stockanalyzer/internal/data_consumer/repository"
	r "stockanalyzer/internal/router"
)

type DataConsumerModule struct {
	DataConsumerRepository *data_consumer_repository.DataConsumerRepository
}

func NewDataConsumerModule(container *container.Container, router *r.Router) *DataConsumerModule {

	fmpProvider := data_consumer_providers.NewFMPDataProvider(container)

	dataConsumerRepository := data_consumer_repository.NewDataConsumerRepository(container, fmpProvider)

	routes := []r.Route{}

	routes = append(routes, router.BuildRoute("GET", "/income_statement", data_consumer_handlers.IncomeStatement(container, dataConsumerRepository)))

	router.SetRoutes("/financials", routes)

	return &DataConsumerModule{
		DataConsumerRepository: dataConsumerRepository,
	}
}
