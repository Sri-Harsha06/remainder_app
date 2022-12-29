package functions

import(
	"remainder_app/dbiface"
	"context"
	"remainder_app/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertData(collection dbiface.CollectionAPI, event model.Event) (*mongo.InsertOneResult, error) {
	res, err := collection.InsertOne(context.Background(), event)
	return res, err
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
