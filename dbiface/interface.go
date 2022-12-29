package dbiface

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CollectionAPI interface {
	InsertOne(ctx context.Context, document interface{},
		opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
	FindOne(ctx context.Context, filter interface{},
			opts ...*options.FindOneOptions) *mongo.SingleResult
	Find(ctx context.Context, filter interface{},
			opts ...*options.FindOptions) (*mongo.Cursor,error)
	ReplaceOne(ctx context.Context, filter interface{},
				replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) 
	DeleteOne(ctx context.Context, filter interface{},
				opts ...*options.DeleteOptions) (*mongo.DeleteResult, error)
}
