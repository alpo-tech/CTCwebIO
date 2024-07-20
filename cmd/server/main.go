package main

import (
	"CRCwebIO/internal/app"
)

func main() {
	application := app.NewApp()
	application.InitializeRoutes()
	application.Run(":8088")
}
