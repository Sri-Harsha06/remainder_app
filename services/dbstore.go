package services

import (
	"context"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var once sync.Once

func GetInstance() *mongo.Client {
	once.Do(func() {
		ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
		clientOptions := options.Client().ApplyURI("mongodb+srv://harsha:harsha@cluster0.xnvctix.mongodb.net/?retryWrites=true&w=majority")
		client, _ = mongo.Connect(ctx, clientOptions)
	})
	return client
}
