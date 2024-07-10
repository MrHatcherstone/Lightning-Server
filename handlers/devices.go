package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type (
	DevicesHandler interface {
		InitDevice(c echo.Context) error
		GetDevices(c echo.Context) error
	}

	devicesHandler struct {
	}
)

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

func (h *devicesHandler) InitDevice(c echo.Context) error {
	body := c.Request().Body
	defer body.Close()

	var data map[string]interface{}
	err := json.NewDecoder(body).Decode(&data)
	if err != nil {
		return err
	}

	ip := data["ip"].(string)
	uuid := data["uuid"].(string)
	// device := Device{
	// 	ip:   ip,
	// 	uuid: uuid,
	// }
	// writeDeviceToFile(device, "./data/device.txt")
	return c.String(http.StatusOK, "data:\n"+"ip:"+ip+", uuid:"+uuid)
}

func (h *devicesHandler) GetDevices(c echo.Context) error {
	return c.String(http.StatusOK, "GetDevices")
}
