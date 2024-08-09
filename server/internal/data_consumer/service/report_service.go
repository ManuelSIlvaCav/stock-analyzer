package data_consumer_service

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"
	"stockanalyzer/internal/container"
	data_consumer_repository "stockanalyzer/internal/data_consumer/repository"

	"github.com/xuri/excelize/v2"
)

type ReportService struct {
	dataRepository *data_consumer_repository.DataConsumerRepository
	container      *container.Container
}

func NewReportService(container *container.Container, dataConsumerRepository *data_consumer_repository.DataConsumerRepository) *ReportService {
	return &ReportService{
		dataRepository: dataConsumerRepository,
		container:      container,
	}
}

func (s *ReportService) GetReport(symbol string) (*excelize.File, error) {
	f := excelize.NewFile()

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// Create a new sheet.
	index, err := f.NewSheet("Sheet2")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	f.SetActiveSheet(index)

	s.AddIncomeSheet(context.Background(), f, symbol)

	return f, nil
}

func (s *ReportService) AddIncomeSheet(
	ctx context.Context,
	excelFile *excelize.File,
	symbol string) (*excelize.File, error) {

	data, err := s.dataRepository.GetIncomeStatement(ctx, symbol)

	if err != nil {
		return nil, err
	}

	//Process data and make a map year: data
	mappedData := map[string]map[string]interface{}{}

	years := []string{}

	for _, item := range data {
		mappedData[item.CalendarYear] = convertToMap(item)
		years = append(years, item.CalendarYear)
	}

	sort.Strings(years)

	fmt.Printf("Income statement retrieved for %s: %v\n", symbol, data)
	fmt.Printf("years: %v\n", years)
	// Create a new sheet.
	index, err := excelFile.NewSheet("Income Statement")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	excelFile.SetActiveSheet(index)

	//1. Set the heads
	//2. Set the first column of names
	//3. Set the rest of the columns of data
	//{"2024": {"Revenue": 100, "Cost of Revenue": 50, "Gross Profit": 50}, "2025": {"Revenue": 200, "Cost of Revenue": 100, "Gross Profit": 100}}
	headerIndex := 2

	for i, header := range getDataHeaders() {
		excelFile.SetCellValue("Income Statement", fmt.Sprintf("%s%d", "A", headerIndex+i), header)
	}

	for _, year := range years {
		yearData := mappedData[year]
		excelFile.SetCellValue("Income Statement", fmt.Sprintf("%s%d", "C", 0), year)

		for i, header := range getDataHeaders() {
			excelFile.SetCellValue("Income Statement", fmt.Sprintf("%s%d", "C", headerIndex+i), yearData[header])
		}
	}

	return excelFile, nil
}

func getDataHeaders() []string {
	columnHeaders := GetDataHeaders()
	return columnHeaders
}

func convertToMap(data interface{}) map[string]interface{} {
	var dataMap map[string]interface{}
	inrec, _ := json.Marshal(data)

	json.Unmarshal(inrec, &dataMap)
	fmt.Printf("dataMap: %v\n", dataMap)

	return dataMap
}
