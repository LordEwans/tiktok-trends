package main

import (
	"log"

	tiktok_api "github.com/arthurkushman/tiktok-api"
)

func main() {
	tts := tiktok_api.NewTikTokService()

	resp, err := tts.Embed()
	if err != nil {
		log.Fatal(err)
	}
}
