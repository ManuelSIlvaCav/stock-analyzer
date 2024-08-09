package data_consumer_providers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"stockanalyzer/internal/container"
	data_consumer_models "stockanalyzer/internal/data_consumer/models"
)

type FMPDataProvider struct {
	apiKey string
}

func NewFMPDataProvider(container *container.Container) *FMPDataProvider {
	return &FMPDataProvider{
		apiKey: container.Config.ExternalAPIConfig.FmpApiKey,
	}
}

// This is the income statement formated by the provider not as reported.
func (f *FMPDataProvider) GetIncomeStatement(ctx context.Context, symbol string, original bool) ([]byte, error) {
	if original {
		return f.GetIncomeStatementAsReported(ctx, symbol, original)
	}
	return f.GetIncomeStatementEdit(ctx, symbol)
}

func (f *FMPDataProvider) GetIncomeStatementEdit(ctx context.Context, symbol string) ([]byte, error) {
	fmt.Printf("Getting income statement for %s\n", symbol)

	method := "GET"

	baseURL := "https://6662ab2f-82f9-445c-9de7-5d547eff2b12.mock.pstmn.io"

	url := fmt.Sprintf("%s/api/v3/income-statement/%s?period=annual&apikey=%s", baseURL, symbol, f.apiKey)

	req, err := http.NewRequest(method, url, nil)

	fmt.Println("url", url)

	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	defer res.Body.Close()

	fmt.Println("returning res")

	if res.StatusCode != 200 {
		fmt.Println("Error getting income statement", res.StatusCode)
		return nil, fmt.Errorf("error getting income statement")
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return body, nil
}

func (f *FMPDataProvider) GetIncomeStatementAsReported(ctx context.Context, symbol string, original bool) ([]byte, error) {
	fmt.Printf("Getting income statement for %s\n", symbol)

	method := "GET"

	baseURL := "https://6662ab2f-82f9-445c-9de7-5d547eff2b12.mock.pstmn.io"

	url := fmt.Sprintf("%s/api/v3/income-statement-as-reported/%s?period=annual&apikey=%s", baseURL, symbol, f.apiKey)

	req, err := http.NewRequest(method, url, nil)

	fmt.Println("url", url)

	if err != nil {
		fmt.Println("error", err)
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	defer res.Body.Close()

	fmt.Println("returning res")

	if res.StatusCode != 200 {
		fmt.Println("Error getting income statement", res.StatusCode)
		return nil, fmt.Errorf("error getting income statement")
	}

	body, err := io.ReadAll(res.Body)

	if err != nil {
		fmt.Print(err)
		return nil, err
	}

	return body, nil
}

func (f *FMPDataProvider) GeneralSearch(ctx context.Context, name string) (data_consumer_models.GeneralSearch, error) {
	return data_consumer_models.GeneralSearch{}, nil
}
