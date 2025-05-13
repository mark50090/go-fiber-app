package pkg

import (
	"context"
	"fmt"
	"go-fiber-app/configs"

	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func MongoBuilder(cfg *configs.Config) *mongo.Database {
	uri := cfg.Database.Connection
	options := options.Client().ApplyURI(uri).SetCompressors([]string{"zstd"})

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		log.Panic(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Mongo connected")
	return client.Database(cfg.Database.DbName)
}
