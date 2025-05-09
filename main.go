package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"

	"go-fiber-app/configs"
	"go-fiber-app/router"
)

func main() {
	godotenv.Load()

	configs.ConnectMongo()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app := fiber.New()

	router.SetupRoutes(app)

	log.Printf("🚀 Server starting on port %s...\n", port)
	log.Fatal(app.Listen(":" + port))
}
