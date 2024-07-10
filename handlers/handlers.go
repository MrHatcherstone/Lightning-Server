package handlers

import (
	"fmt"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Template struct {
	Template *template.Template
}

type Handlers struct {
	DevicesHandler
	SettingsHandler
}

func New() *Handlers {
	return &Handlers{
		DevicesHandler:  &devicesHandler{},
		SettingsHandler: &settingsHandler{},
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Template.ExecuteTemplate(w, name, data)
}

func newTemplate() *Template {
	return &Template{
		Template: template.Must(template.ParseGlob("views/*.html")),
	}
}

func SetDefault(e *echo.Echo) {
	e.Renderer = newTemplate()
	e.GET("/healthcheck", HealthCheckHandler)
}

var data DataPage

func SetApi(e *echo.Echo, h *Handlers) {
	//tmp data
	fmt.Print("HELLO WORLD", data)
	data = getSettingsPageData()

	e.Static("/static", "static")
	e.GET("/", func(c echo.Context) error {
		return h.GetPageSettings(c)
	})

	g := e.Group("/api/v1")
	g.POST("/devices/init", h.InitDevice)
	g.GET("/devices", h.GetDevices)
	g.POST("/settings/save", func(c echo.Context) error {
		return h.SaveSettings(c)
	})
}

func Echo() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	return e
}
