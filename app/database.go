package app

import (
	"context"
	"donations/helper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func NewSetupDB() *mongo.Collection {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)

	connectOptions := &options.ClientOptions{}
	connectOptions.ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(connectOptions)
	helper.PanicIfError(err)

	err = client.Connect(ctx)
	helper.PanicIfError(err)

	return client.Database("donations").Collection("donations")
}
