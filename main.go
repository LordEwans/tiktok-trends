package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	arr := fetchData(c)
	fmt.Print(arr)
}

func fetchData(c *colly.Collector) []string {
	arr := make([]string, 0)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.OnHTML("strong", func(h *colly.HTMLElement) {
		class := h.Attr("class")

		if class == "tiktok-1p6dp51-StrongText ejg0rhn2" {
			arr = append(arr, h.Text)
		}
	})
	c.Visit("http://tiktok.com/explore")

	return arr
}
