package repo

import (
	"context"

	"github.com/Rendyfranzz/Study-Case-Chapter-1/config"
	"github.com/Rendyfranzz/Study-Case-Chapter-1/model"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
)

type NIBRepo struct {
	db *config.DB
}

func NewNIBRepo(db *config.DB) *NIBRepo {
	return &NIBRepo{db: db}
}

func (n *NIBRepo) GetBy(ctx context.Context, nibCode string) (model.NIB, error) {
	var nib model.NIB

	collection := n.db.Client.Database("oss").Collection("nib")

	err := collection.FindOne(ctx, bson.M{"nib": nibCode}).Decode(&nib)
	return nib, err
}

func (n *NIBRepo) Insert(ctx context.Context, nib model.NIB) error {
	nib.ID = xid.New().String()
	collection := n.db.Client.Database("oss").Collection("nib")
	_, err := collection.InsertOne(ctx, &nib)

	return err
}
