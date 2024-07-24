package data_consumer_providers

import (
	"context"
	data_consumer_models "stockanalyzer/internal/data_consumer/models"
)

type FinancialDataProvider interface {
	GetIncomeStatement(ctx context.Context, symbol string, original bool) (data_consumer_models.IncomeStatement, error)
	GeneralSearch(ctx context.Context, name string) (data_consumer_models.GeneralSearch, error)
}
