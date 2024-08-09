package data_consumer_handlers

import (
	"stockanalyzer/internal/container"
	data_consumer_service "stockanalyzer/internal/data_consumer/service"

	"github.com/labstack/echo/v4"
)

func Report(container *container.Container, reportService *data_consumer_service.ReportService) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := reportService.GetReport("ULTA")

		if err != nil {
			return c.JSON(500, "Error getting income statement")
		}

		c.Response().Header().Set("Content-Type", echo.MIMEOctetStream)
		c.Response().Header().Set("Content-Disposition", "attachment; filename="+"WorkbookTest.xlsx")
		c.Response().Header().Set("Content-Transfer-Encoding", "binary")
		c.Response().Header().Set("Expires", "0")

		file.WriteTo(c.Response().Writer)

		c.Response().Flush()
		return nil
	}
}
