package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CloseConnFunc func(ctx context.Context) error

func Connect(uri string) (*mongo.Client, CloseConnFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, nil, err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return nil, nil, err
	}

	f := func(ctx context.Context) error {
		if err = client.Disconnect(ctx); err != nil {
			return err
		}

		return nil
	}

	return client, f, nil
}
