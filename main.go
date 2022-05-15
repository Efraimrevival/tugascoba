package main

import (
	"eraport/route"
	"eraport/config"
)

func main() {
	config.InitDB()
	e := route.New()
	e.Logger.Fatal(e.Start(":8080"))
}