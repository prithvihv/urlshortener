package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type tinyURL struct {
	FullURL    string
	TinyURLuid string
	Hits       int
	CreatedAt  int64
}

func updateHit(tinyurl string) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	filter := bson.M{"tinyurluid": tinyurl}
	update := bson.M{
		"$inc": bson.M{"hits": 1},
	}
	tinysCollection.FindOneAndUpdate(ctx, filter, update, &options.FindOneAndUpdateOptions{})
}

var (
	client          *mongo.Client
	mongoURL        = "mongodb://mongo.prv:27017"
	tinysCollection *mongo.Collection
)

func initDb() {

	// Initialize a new mongo client with options
	client, _ = mongo.NewClient(options.Client().ApplyURI(mongoURL))

	// Connect the mongo client to the MongoDB server
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := client.Connect(ctx)

	// Ping MongoDB
	ctx, _ = context.WithTimeout(context.Background(), 1*time.Second)
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Println("could not ping to mongo db service: %v\n", err)
		return
	}

	tinysCollection = client.Database("testing").Collection("tinyURLs")
	fmt.Println("connected to nosql database:", mongoURL)
}
