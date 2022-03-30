package main

import (
	"fetch/application"
	"fetch/config"
)

func main() {

	conf := config.GetConfig()

	app := application.New(conf)

	app.Start()
}
