package main

import (
	"it-bni/config"
	"it-bni/middlewares"
	"it-bni/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	middlewares.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8080"))
}
