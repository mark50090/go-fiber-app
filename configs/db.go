package configs

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var mongoClient *mongo.Client

func ConnectMongo() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set in .env")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("MongoDB ping error:", err)
	}

	// เปลี่ยน mongoClient เป็น MongoClient เพื่อให้มันใช้งานได้ในส่วนอื่นๆ
	MongoClient = client
	log.Println("✅ Connected to MongoDB!")
}

func GetMongoClient() *mongo.Client {
	if MongoClient == nil {
		ConnectMongo()
	}
	return MongoClient
}
