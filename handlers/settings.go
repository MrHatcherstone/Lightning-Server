package handlers

import (
	"fmt"
	"lightningServer/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type (
	SettingsHandler interface {
		GetPageSettings(c echo.Context) error
		SaveSettings(c echo.Context) error
	}

	settingsHandler struct {
	}
)

type DataPage struct {
	Settings      models.DeviceSettings
	LightingModes []*models.Lighting
}

const (
	static  = "static"
	dynamic = "dynamic"
	rainbow = "rainbow"
)

func getSettingsPageData() DataPage {
	LightingModes := []*models.Lighting{
		{ID: 1, Name: static},
		{ID: 2, Name: dynamic},
		{ID: 3, Name: rainbow},
	}

	return DataPage{
		Settings: models.DeviceSettings{
			Brightness: 100,
			Delay:      100,
			IsFinished: false,
			RGB:        []int{0, 0, 0},
		},
		LightingModes: LightingModes,
	}
}

func (h *settingsHandler) GetPageSettings(c echo.Context) error {
	return c.Render(http.StatusOK, "index", data)
}

func (h *settingsHandler) SaveSettings(c echo.Context) error {

	// tmp need to move to service
	brightness, err := strconv.Atoi(c.FormValue("brightness"))
	delay, err := strconv.Atoi(c.FormValue("delay"))
	isFinished := checkIsCheckboxChecked(c.FormValue("finished"))

	RGB := parseColorRgb(c.FormValue("red"), c.FormValue("green"), c.FormValue("blue"))

	fmt.Println("lighting", c.FormValue("lighting"))
	fmt.Println("END ===============")

	newSettings := models.DeviceSettings{
		Brightness: brightness,
		IsFinished: isFinished,
		Delay:      delay,
		RGB:        RGB,
	}

	// handle error
	if err != nil {
	}

	data.Settings = newSettings

	return c.Render(http.StatusOK, "form", data)
}

func checkIsCheckboxChecked(value interface{}) bool {
	if value == nil {
		return false
	}
	strValue, ok := value.(string)
	if !ok {
		return false
	}
	if strValue == "on" || strValue == "true" {
		return true
	}
	return false
}

func parseColorRgb(r string, g string, b string) []int {
	red, err := strconv.Atoi(r)
	green, err := strconv.Atoi(g)
	blue, err := strconv.Atoi(b)
	if err != nil {
	}

	return []int{red, green, blue}
}
