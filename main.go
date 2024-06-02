package main

import (
	"fmt"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/lpernett/godotenv"
	"log"
	"os"
)

type Slide struct {
	ID     int    `json:"id"`
	Public bool   `json:"public"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")

	log.Println("Server started on :" + PORT)
	app := fiber.New()

	slides := []Slide{}

	app.Get("/api/slides", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(slides)
	})

	// Create Slide
	app.Post("/api/slides", func(c *fiber.Ctx) error {
		slide := &Slide{}

		if err := c.BodyParser(slide); err != nil {
			return err
		}

		if slide.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "Body is empty"})
		}

		slide.ID = len(slides) + 1
		slides = append(slides, *slide)

		return c.Status(201).JSON(slide)
	})

	// Update Slide
	app.Patch("/api/slides/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, slide := range slides {
			if fmt.Sprint(slide.ID) == id {
				slides[i].Public = true
				return c.Status(200).JSON(slides[i])
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "Slide not found"})
	})

	// Delete slide
	app.Delete("/api/slides/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		for i, slide := range slides {
			if fmt.Sprint(slide.ID) == id {
				slides = append(slides[:i], slides[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"success": "true"})
			}
		}

		return c.Status(404).JSON(fiber.Map{"error": "Slide not found"})
	})

	log.Fatal(app.Listen(":" + PORT))

}
