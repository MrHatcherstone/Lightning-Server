package services

import "github.com/labstack/echo/v4"

type (
	DevicesService interface {
		InitDevice(c echo.Context) error
	}
)
