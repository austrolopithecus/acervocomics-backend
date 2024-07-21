package main

import (
	"acervocomics/util"
	"acervocomics/view"

	"github.com/gofiber/fiber/v3"
)

func main() {
	err := util.InitDB()
	if err != nil {
		panic(err)
	}
	app := fiber.New()

	view.InitRoutes(app)

	_ = app.Listen("127.0.0.1:5556")
}
