package main

import (
	fiber "github.com/gofiber/fiber/v2"
	"log"
)

type Slide struct {
	ID     int    `json:"id"`
	Public bool   `json:"public"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	log.Println("Server started on :4000")
	app := fiber.New()

	slides := []Slide{}

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Fiber started."})
	})

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

	log.Fatal(app.Listen(":4000"))

}
