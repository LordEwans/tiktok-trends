package main

import (
	"os"

	"github.com/LordEwan/tiktok-trends/internal/fetch"
	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":0000"
	} else {
		port = ":" + port
	}

	return port
}

func main() {
	app := fiber.New()
	c := colly.NewCollector()

	arr := fetch.GetData(c)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(arr)
	})

	app.Listen(getPort())
}
