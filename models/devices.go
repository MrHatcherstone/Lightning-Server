package models

type Device struct {
	ID   int
	Name string
}

type DeviceSettings struct {
	Brightness int
	Delay      int
	IsFinished bool
	RGB        []int
}

type Lighting struct {
	ID   int
	Name string
}

const (
	static  = "static"
	dynamic = "dynamic"
	rainbow = "rainbow"
)
