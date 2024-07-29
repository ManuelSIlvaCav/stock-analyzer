package data_consumer_handlers

import (
	"net/http"
	"stockanalyzer/internal/container"
	data_consumer_repository "stockanalyzer/internal/data_consumer/repository"

	"github.com/labstack/echo/v4"
)

func IncomeStatement(container *container.Container, dataConsumerRepository *data_consumer_repository.DataConsumerRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		data, err := dataConsumerRepository.GetIncomeStatementAsReported(c.Request().Context(), "ULTA")

		if err != nil {
			return c.JSON(500, "Error getting income statement")
		}

		return c.JSON(http.StatusOK, &echo.Map{"data": data})
	}
}
