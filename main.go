package main

import (
	"lightningServer/handlers"
	"os"
)

func main() {
	e := handlers.Echo()
	h := handlers.New()

	handlers.SetDefault(e)
	handlers.SetApi(e, h)

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "50064"
	}

	e.Logger.Fatal(e.Start(":" + PORT))
}
