package data_consumer_service

import (
	"fmt"
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

func (s *ReportService) GetReport() (*excelize.File, error) {
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

	return f, nil

}
