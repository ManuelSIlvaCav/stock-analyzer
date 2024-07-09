package data_consumer_providers

import (
	"context"
	data_consumer_models "stockanalyzer/internal/data_consumer/models"
)

type FinancialDataProvider interface {
	GetIncomeStatement(ctx context.Context, symbol string) (data_consumer_models.IncomeStatement, error)
}
