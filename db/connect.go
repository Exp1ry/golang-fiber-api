package db

import (
	"context"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDB() (*mongo.Database, context.CancelFunc, error) {

err := godotenv.Load()
if err!= nil {
	panic(err)
}

mongoUri := os.Getenv("MONGO_URI")

	ctx,cancel:= context.WithTimeout(context.Background(), 10*time.Second)
	if err!= nil {
		panic(err)
	}

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUri))
	if err!= nil {
		panic(err)
	}

	db := client.Database("test")

	return db, cancel, nil
}