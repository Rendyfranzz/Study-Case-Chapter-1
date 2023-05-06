package controller

import (
	"crud/entities"
	"crud/helper"
	"crud/model"
	"encoding/json"
	"net/http"

	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddNews(conn *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var news entities.News

		if err := json.NewDecoder(r.Body).Decode(&news); err != nil {
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}
		defer r.Body.Close()
		news.ID = xid.New().String()
		collection := conn.Database("oss").Collection("news")
		_, err := collection.InsertOne(r.Context(), &news)
		if err != nil {
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
		response := map[string]string{"message": "success"}
		helper.ResponseJSON(w, http.StatusOK, response)
	}
}

func GetNews(conn *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		collection := conn.Database("oss").Collection("news")
		rows, err := collection.Find(r.Context(), bson.M{})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(Map{"error finding": err.Error()})
			return
		}

		var news []entities.News

		for rows.Next(r.Context()) {
			var new entities.News

			err := rows.Decode(&new)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(Map{"error streaming": err.Error()})
				return
			}

			news = append(news, new)
		}

		json.NewEncoder(w).Encode(news)
	}
}

func GetNews2(conn *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fs, err := model.GetNews(conn, r.Context())
	}
}
