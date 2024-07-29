package data_consumer_handlers

import (
	"fmt"
	"stockanalyzer/internal/container"
	data_consumer_repository "stockanalyzer/internal/data_consumer/repository"

	"github.com/labstack/echo/v4"
)

// GeneralSearch is a function that returns a handler function that returns a JSON response with the message "General Search"
func GeneralSearch(container *container.Container, dataConsumerRepository *data_consumer_repository.DataConsumerRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := c.Request().Context()
		container.Cache.Set(ctx, "test_key", 2)
		value, err := container.Cache.Get(ctx, "test_key")
		if err != nil {
			return c.JSON(500, err.Error())
		}
		fmt.Printf("Value: %s\n", value)
		return c.JSON(200, "General Search")
	}
}
