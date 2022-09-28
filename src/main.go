package main

import "github.com/bachelor-service/di"

func main() {
	app, _ := di.NewApp()

	app.Start()
}
