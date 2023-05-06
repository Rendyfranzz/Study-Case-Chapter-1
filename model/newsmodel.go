package model

import (
	"context"
	"crud/entities"

	"go.mongodb.org/mongo-driver/mongo"
)

type news interface {
	GetNews(conn *mongo.Client, ctx context.Context) []entities.News
	AddNews(conn *mongo.Client, news entities.News) bool
}

type News struct {
	ID        string
	Judul     string
	Deskripsi string
	Tanggal   string
}

func (n *News) GetNews(conn *mongo.Client, ctx context.Context) []entities.News {

	return nil
}
