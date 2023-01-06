package functions

import (
	"context"
	"errors"
	"remainder_app/dbiface"
	"remainder_app/model"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertData(collection dbiface.CollectionAPI, event model.Event) (*mongo.InsertOneResult, error) {
	if event.Name == "" {
		return &mongo.InsertOneResult{}, errors.New("no name Specified")
	} else if event.CreatedBy == "" {
		return &mongo.InsertOneResult{}, errors.New("created By Not Specified")
	} else if event.Event == "" {
		return &mongo.InsertOneResult{}, errors.New("event Name Not Specified")
	} else if event.Date < time.Now().Format("2006-01-02") {
		return &mongo.InsertOneResult{}, errors.New("an event Of the past,cannot insert")
	}
	return collection.InsertOne(context.Background(), event)
}

func FindDataById(collection dbiface.CollectionAPI, event model.Event) *mongo.SingleResult {
	res := collection.FindOne(context.Background(), event)
	return res
}

func GetData(collection dbiface.CollectionAPI, event model.Event) (*mongo.Cursor, error) {
	cursor, err := collection.Find(context.Background(), event)
	return cursor, err
}

func UpdateData(collection dbiface.CollectionAPI, event model.Event) (*mongo.UpdateResult, error) {
	filter := bson.D{{Key: "id", Value: event.Id}}
	result, err := collection.ReplaceOne(context.Background(), filter, event)
	return result, err
}

func DeleteData(collection dbiface.CollectionAPI, event model.Event) (*mongo.DeleteResult, error) {
	result, err := collection.DeleteOne(context.Background(), event)
	return result, err
}
