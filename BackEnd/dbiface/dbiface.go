package dbiface

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	CollectionAPI interface {
		InsertOne(ctx context.Context, document interface{},
			opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
		Find(ctx context.Context, filter interface{},
			opts ...*options.FindOptions) (*mongo.Cursor, error)
		FindOne(ctx context.Context, filter interface{},
			opts ...*options.FindOneOptions) *mongo.SingleResult
		FindOneAndUpdate(ctx context.Context, filter interface{},
			update interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult
	}
)
