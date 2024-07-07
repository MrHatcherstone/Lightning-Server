package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	static  = "static"
	dynamic = "dynamic"
	rainbow = "rainbow"
)

type Lighting struct {
	ID   int
	Name string
}

type DeviceSettings struct {
	Brightness int
	Delay      int
	IsFinished bool
	RGB        []int
}

type DataPage struct {
	Settings      DeviceSettings
	LightingModes []*Lighting
}

type Template struct {
	Template *template.Template
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	data := newData()
	e.Renderer = newTemplate()
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", data)
	})

	e.POST("/init", initDevice)

	e.POST("/save", func(c echo.Context) error {
		brightness, err := strconv.Atoi(c.FormValue("brightness"))
		delay, err := strconv.Atoi(c.FormValue("delay"))
		isFinished := checkIsCheckboxChecked(c.FormValue("finished"))

		RGB := parseColorRgb(c.FormValue("red"), c.FormValue("green"), c.FormValue("blue"))

		fmt.Println("lighting", c.FormValue("lighting"))
		fmt.Println("END ===============")

		newSettings := newSettings(DeviceSettings{
			Brightness: brightness,
			IsFinished: isFinished,
			Delay:      delay,
			RGB:        RGB,
		})

		fmt.Println("parsed data is", newSettings)

		// handle error
		if err != nil {
		}

		data.Settings = newSettings
		return c.Render(200, "form", data)
	})

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":50064"))
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Template.ExecuteTemplate(w, name, data)
}

func newTemplate() *Template {
	return &Template{
		Template: template.Must(template.ParseGlob("views/*.html")),
	}
}

func initDevice(c echo.Context) error {
	body := c.Request().Body
	defer body.Close()

	var data map[string]interface{}
	err := json.NewDecoder(body).Decode(&data)
	if err != nil {
		return err
	}

	ip := data["ip"].(string)
	uuid := data["uuid"].(string)
	device := Device{
		ip:   ip,
		uuid: uuid,
	}
	writeDeviceToFile(device, "./data/device.txt")
	return c.String(http.StatusOK, "data:\n"+"ip:"+ip+", uuid:"+uuid)
}

func newSettings(settings DeviceSettings) DeviceSettings {
	return settings
}

func newData() DataPage {

	// Const
	LightingModes := []*Lighting{
		{1, static},
		{2, dynamic},
		{3, rainbow},
	}

	return DataPage{
		Settings: newSettings(DeviceSettings{
			IsFinished: false,
			Brightness: 10,
			Delay:      0,
			RGB:        []int{0, 0, 0},
		}),
		LightingModes: LightingModes,
	}
}

type Device struct {
	uuid string
	ip   string
}

func writeToFile(data string, filename string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}
func writeDeviceToFile(device Device, filename string) error {
	data := fmt.Sprintf("device: ip: %s, uuid: %s\n", device.ip, device.uuid)
	return writeToFile(data, filename)
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
