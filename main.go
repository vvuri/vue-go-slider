package main

import (
	"context"
	fiber "github.com/gofiber/fiber/v2"
	"github.com/lpernett/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type Slide struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Public bool               `json:"public"`
	Title  string             `json:"title"`
	Body   string             `json:"body"`
}

var collection *mongo.Collection

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	MONGODB_URI := os.Getenv("MONGODB_URI")

	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("MongoDB connected")

	collection = client.Database("vue-go-slider").Collection("slides")

	log.Println("Server started on :" + PORT)
	app := fiber.New()

	app.Get("/api/slides", getSlides)
	app.Post("/api/slides", createSlide)
	//app.Patch("/api/slides/:id", updateSlides)
	//app.Delete("/api/slides/:id", deleteSlides)

	log.Fatal(app.Listen(":" + PORT))
}

func getSlides(c *fiber.Ctx) error {
	var slides []Slide

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var slide Slide
		if err = cursor.Decode(&slide); err != nil {
			return err
		}
		slides = append(slides, slide)
	}

	return c.Status(200).JSON(slides)
}

func createSlide(c *fiber.Ctx) error {
	slide := new(Slide)

	if err := c.BodyParser(slide); err != nil {
		return err
	}

	if slide.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Body is empty"})
	}

	insertResult, err := collection.InsertOne(context.Background(), slide)
	if err != nil {
		return err
	}

	slide.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(slide)
}

/*
func updateSlides(c *fiber.Ctx) error {
	id := c.Params("id")

	for i, slide := range slides {
		if fmt.Sprint(slide.ID) == id {
			slides[i].Public = true
			return c.Status(200).JSON(slides[i])
		}
	}

	return c.Status(404).JSON(fiber.Map{"error": "Slide not found"})
}

func deleteSlides(c *fiber.Ctx) error {
	id := c.Params("id")

	for i, slide := range slides {
		if fmt.Sprint(slide.ID) == id {
			slides = append(slides[:i], slides[i+1:]...)
			return c.Status(200).JSON(fiber.Map{"success": "true"})
		}
	}

	return c.Status(404).JSON(fiber.Map{"error": "Slide not found"})
}
*/
