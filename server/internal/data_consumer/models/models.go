package data_consumer_models

type IncomeStatementYear struct {
	Date                           string  `json:"date"`
	Symbol                         string  `json:"symbol"`
	ReportedCurrency               string  `json:"reportedCurrency"`
	CIK                            string  `json:"cik"`
	FillingDate                    string  `json:"fillingDate"`
	CalendarYear                   string  `json:"calendarYear"`
	Period                         string  `json:"period"`
	Revenue                        int64   `json:"revenue"`
	CostOfRevenue                  int64   `json:"costOfRevenue"`
	GrossProfit                    int64   `json:"grossProfit"`
	GrossProfitRatio               float64 `json:"grossProfitRatio"`
	ResearchAndDev                 int64   `json:"researchAndDevelopmentExpenses"`
	GeneralAndAdmin                int64   `json:"generalAndAdministrativeExpenses"`
	SellingAndMkt                  int64   `json:"sellingAndMarketingExpenses"`
	SellingGeneralAndAdminExpenses int64   `json:"sellingGeneralAndAdministrativeExpenses"`
	OtherExpenses                  int64   `json:"otherExpenses"`
	OperatingExpenses              int64   `json:"operatingExpenses"`
	CostAndExpenses                int64   `json:"costAndExpenses"`
	InterestIncome                 int64   `json:"interestIncome"`
	InterestExpense                int64   `json:"interestExpense"`
	DepreciationAndAmortization    int64   `json:"depreciationAndAmortization"`
	Ebitda                         int64   `json:"ebitda"`
	Ebitdaratio                    float64 `json:"ebitdaratio"`
	OperatingIncome                int64   `json:"operatingIncome"`
	OperatingIncomeRatio           float64 `json:"operatingIncomeRatio"`
	TotalOtherIncomeExpensesNet    int64   `json:"totalOtherIncomeExpensesNet"`
	IncomeBeforeTax                int64   `json:"incomeBeforeTax"`
	IncomeBeforeTaxRatio           float64 `json:"incomeBeforeTaxRatio"`
	IncomeTaxExpense               int64   `json:"incomeTaxExpense"`
	NetIncome                      int64   `json:"netIncome"`
	NetIncomeRatio                 float64 `json:"netIncomeRatio"`
	Eps                            float64 `json:"eps"`
	Epsdiluted                     float64 `json:"epsdiluted"`
	WeightedAverageShsOut          int64   `json:"weightedAverageShsOut"`
	WeightedAverageShsOutDil       int64   `json:"weightedAverageShsOutDil"`
}

type GeneralSearch struct {
	Symbol            string
	Name              string
	Currency          string
	StockExchange     string
	ExchangeShortName string
}
