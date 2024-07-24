package data_consumer_providers

import (
	"context"
	"encoding/json"
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

func (f *FMPDataProvider) GetIncomeStatement(ctx context.Context, symbol string, original bool) (data_consumer_models.IncomeStatement, error) {
	fmt.Printf("Getting income statement for %s\n", symbol)

	method := "GET"

	baseURL := "https://6662ab2f-82f9-445c-9de7-5d547eff2b12.mock.pstmn.io"

	url := fmt.Sprintf("%s/api/v3/income-statement/%s?period=annual&apikey=%s", baseURL, symbol, f.apiKey)

	req, err := http.NewRequest(method, url, nil)

	fmt.Println("url", url)

	if err != nil {
		fmt.Println("error", err)
		return data_consumer_models.IncomeStatement{}, err
	}

	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		fmt.Print(err)
		return data_consumer_models.IncomeStatement{}, err
	}

	defer res.Body.Close()

	fmt.Println("returning res")

	if res.StatusCode != 200 {
		fmt.Println("Error getting income statement", res.StatusCode)
		return data_consumer_models.IncomeStatement{}, fmt.Errorf("error getting income statement")
	}

	body, err := io.ReadAll(res.Body)

	//fmt.Printf("body: %s\n", body)

	if err != nil {
		fmt.Print(err)
		return data_consumer_models.IncomeStatement{}, err
	}

	fmt.Println("parsing body")
	var data []map[string]interface{}

	_ = json.Unmarshal(body, &data)

	//fmt.Print("jsonMap", data)

	return data_consumer_models.IncomeStatement{}, nil
}

func (f *FMPDataProvider) GeneralSearch(ctx context.Context, name string) (data_consumer_models.GeneralSearch, error) {
	return data_consumer_models.GeneralSearch{}, nil
}
