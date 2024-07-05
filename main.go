package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	Template *template.Template
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Renderer = newTemplate()
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index", nil)
	})

	e.GET("/init", func(c echo.Context) error {
		return c.String(http.StatusOK, "GArri has big cock")
	})

	e.POST("/init", initDevice)

	e.POST("/save", saveSettings)

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
	fmt.Println("printed body:", data)
	return c.String(http.StatusOK, "data:\n"+"ip:"+ip+", uuid:"+uuid)
}

func saveSettings(c echo.Context) error {
	body := c.Request().Body
	fmt.Println(body)
	return nil
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
