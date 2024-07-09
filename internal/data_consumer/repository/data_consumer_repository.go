package data_consumer_repository

import (
	"context"
	"fmt"
	"stockanalyzer/internal/container"
	"stockanalyzer/internal/container/postgres/stock_analyzer_pg"
	data_consumer_providers "stockanalyzer/internal/data_consumer/providers"
)

type DataConsumerRepository struct {
	provider data_consumer_providers.FinancialDataProvider
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
		provider: provider,
	}
}

func (r *DataConsumerRepository) GetIncomeStatement(ctx context.Context, symbol string) error {
	data, err := r.provider.GetIncomeStatement(ctx, symbol)
	if err != nil {
		return err
	}

	fmt.Println("Income statement retrieved", data)
	return nil
}
