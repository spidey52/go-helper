package helper

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func FindOne[T any](collection *mongo.Collection, filter interface{}, opts ...*options.FindOneOptions) (*T, error) {
	var result T

	if filter == nil {
		filter = bson.M{}
	}

	err := collection.FindOne(context.Background(), filter, opts...).Decode(&result)

	if err == mongo.ErrNoDocuments {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &result, nil
}

func FindMany[T any](collection *mongo.Collection, filter interface{}, opts ...*options.FindOptions) ([]T, error) {
	var result []T = make([]T, 0)

	if filter == nil {
		filter = bson.M{}
	}

	cursor, err := collection.Find(context.Background(), filter, opts...)

	if err != nil {
		return nil, err
	}

	if err = cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	fmt.Println("Found", len(result), "documents")

	return result, nil
}

func FindById[T any](collection *mongo.Collection, id string) (*T, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	return FindOne[T](collection, bson.M{"_id": objectId})
}

func ToObjectId(id string) primitive.ObjectID {
	objectId, _ := primitive.ObjectIDFromHex(id)
	return objectId
}
