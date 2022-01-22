package main

import (
	"it-bni/config"
	"it-bni/routes"
)

func main() {
	config.InitDB()
	e := routes.New()

	e.Logger.Fatal(e.Start(":8080"))
}
