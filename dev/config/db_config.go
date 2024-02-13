package config

import (
	"context"
	"fmt"
	"log"
	"net/url"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoDBClient() (*mongo.Client, error) {
	var ctx context.Context
	cfg := Configuration.DB
	optionsStr := fmt.Sprintf("ssl=false&%s=%s", "authSource", url.QueryEscape(cfg.AuthSource))

	// Set the MongoDB client options
	clientOptions := options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d/?%s",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, optionsStr))

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	var result bson.M
	if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Decode(&result); err != nil {
		return nil, err
	}

	return client, nil
}

func CloseDatabaseConnection(c *mongo.Client) {
	if err := c.Disconnect(context.Background()); err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully disconnected from MongoDB.")
}
