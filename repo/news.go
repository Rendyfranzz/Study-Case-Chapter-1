package repo

import (
	"context"

	"github.com/Rendyfranzz/Study-Case-Chapter-1/config"
	"github.com/Rendyfranzz/Study-Case-Chapter-1/model"
	"go.mongodb.org/mongo-driver/bson"
)

type NewsRepo struct {
	db *config.DB
}

func NewNewsRepo(db *config.DB) *NewsRepo {
	return &NewsRepo{db: db}
}

func (n *NewsRepo) GetNews(ctx context.Context) ([]model.News, error) {
	var news []model.News
	collection := n.db.Client.Database("oss").Collection("news")
	rows, err := collection.Find(ctx, bson.M{})

	for rows.Next(ctx) {
		var new model.News

		rows.Decode(&new)
		news = append(news, new)
	}
	return news, err
}
