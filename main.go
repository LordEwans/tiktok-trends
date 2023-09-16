package main

import (
	"github.com/LordEwan/tiktok-trends/internal/fetch"
	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	c := colly.NewCollector()

	arr := fetch.GetData(c)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(arr)
	})

	app.Listen(":0000")
}
