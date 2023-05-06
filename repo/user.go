package repo

import (
	"context"

	"github.com/Rendyfranzz/Study-Case-Chapter-1/config"
	"github.com/Rendyfranzz/Study-Case-Chapter-1/model"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRepo struct {
	db *config.DB
}

func NewUserRepo(db *config.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) GetBy(ctx context.Context, email string) (model.User, error) {
	var user model.User

	collection := u.db.Client.Database("oss").Collection("users")

	err := collection.FindOne(ctx, bson.M{"email": email}).Decode(&user)
	return user, err
}

func (u *UserRepo) Insert(ctx context.Context, user model.User) error {
	user.ID = xid.New().String()
	collection := u.db.Client.Database("oss").Collection("users")
	_, err := collection.InsertOne(ctx, &user)

	return err
}
