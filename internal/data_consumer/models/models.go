package data_consumer_models

type IncomeStatement struct {
	Symbol           string
	ReportedCurrency string
	CalendarYear     string
	Period           string
}

type GeneralSearch struct {
	Symbol            string
	Name              string
	Currency          string
	StockExchange     string
	ExchangeShortName string
}
