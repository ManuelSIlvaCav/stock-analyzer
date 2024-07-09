package internal

import (
	"stockanalyzer/internal/container"
	"stockanalyzer/internal/data_consumer"
	"stockanalyzer/internal/router"

	"go.uber.org/fx"
)

type ModulesContainer struct {
}

func NewModulesContainer(container *container.Container, router *router.Router,
	dataConsumerModule *data_consumer.DataConsumerModule) *ModulesContainer {
	return &ModulesContainer{}
}

var Modules = fx.Options(
	fx.Provide(NewModulesContainer),
	fx.Provide(data_consumer.NewDataConsumerModule),
)
