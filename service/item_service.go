package service

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"

	"go-fiber-app/configs"
	"go-fiber-app/models"
)

func CreateItem(c *fiber.Ctx) error {
	var item models.Item
	if err := c.BodyParser(&item); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if item.Name == "" || item.Price <= 0 {
		return c.Status(400).JSON(fiber.Map{"error": "Name and Price are required and must be valid"})
	}

	collection := configs.MongoClient.Database("myappdb").Collection("items")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, bson.M{
		"name":       item.Name,
		"price":      item.Price,
		"created_at": time.Now(),
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": fmt.Sprintf("Failed to insert item: %v", err)})
	}

	return c.Status(201).JSON(fiber.Map{"message": "Item created successfully", "item_id": res.InsertedID, "item": item})
}
