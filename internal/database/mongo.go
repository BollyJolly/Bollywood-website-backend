package database

import (
	"context"
	"log"
	"time"

	"github.com/amanhasnainy/bingo-backend/internal/config"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

var DB *mongo.Database

func ConnectMongo() {

	ctx, cancel := context.WithTimeout(
		context.Background(),
		10*time.Second,
	)

	defer cancel()

	client, err := mongo.Connect(
		options.Client().
		ApplyURI(config.GetEnv("MONGO_URI")),
	)
	
	err = client.Ping(ctx, nil)
	
	if err != nil {
		log.Fatal(err)
	}


	DB = client.Database("bolly_bingo")

	log.Println("✅ MongoDB connected")
}