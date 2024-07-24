package data_consumer_repository

import (
	"context"
	"fmt"
	"stockanalyzer/internal/container"
	"stockanalyzer/internal/container/postgres/stock_analyzer_pg"
	data_consumer_providers "stockanalyzer/internal/data_consumer/providers"
)

type DataConsumerRepository struct {
	provider  data_consumer_providers.FinancialDataProvider
	container *container.Container
}

func NewDataConsumerRepository(container *container.Container,
	provider data_consumer_providers.FinancialDataProvider) *DataConsumerRepository {
	ctx := context.Background()

	connection := container.PostgresSQL.GetConnection()

	queries := stock_analyzer_pg.New(connection)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		fmt.Println("Error listing authors:", err)
		return &DataConsumerRepository{
			provider: provider,
		}
	}
	fmt.Println("Authors:", authors)

	return &DataConsumerRepository{
		provider:  provider,
		container: container,
	}
}

func (r *DataConsumerRepository) GetIncomeStatement(ctx context.Context, symbol string) error {
	data, err := r.provider.GetIncomeStatement(ctx, symbol, true)
	if err != nil {
		return err
	}

	fmt.Println("Income statement retrieved", data)
	return nil
}

func (r *DataConsumerRepository) SearchName(ctx context.Context, name string) error {
	fmt.Println("Searching for name with name:", name)
	//We look at redis if we have the key
	cacheResult, err := r.container.Cache.Get(ctx, name)

	if len(cacheResult) == 0 || err != nil {
		data, err := r.provider.GeneralSearch(ctx, name)
		if err != nil {
			return err
		}
		fmt.Println("Data retrieved", data)

		//TODO save this data into elastic search or postgres

		//We store the key in redis to signal already fetched
		err = r.container.Cache.Set(ctx, name, data)
		if err != nil {
			return err
		}
		return nil
	}

	fmt.Println("Data retrieved from cache", cacheResult)

	return nil
}
