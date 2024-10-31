package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

type CollectionPreview struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
}

type CollectionDetail struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Cover       string   `json:"cover"`
	Images      []string `json:"images"`
}

type SongPreview struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Cover       string `json:"cover"`
}

type SongDetail struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Cover       string   `json:"cover"`
	Images      []string `json:"images"`
}

type ShewProtocol struct {
	CollectionList   string `json:"collectionList"`
	CollectionDetail string `json:"collectionDetail"`

	SongDetail string `json:"songDetail"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	log.Fatal(app.Listen(":3000"))
}
