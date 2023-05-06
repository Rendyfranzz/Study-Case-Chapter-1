package controller

import (
	"crud/entities"
	"crud/helper"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetNib(conn *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keyToFind := chi.URLParam(r, "id")
		var nib entities.Nib
		collection := conn.Database("oss").Collection("nib")

		err := collection.FindOne(r.Context(), bson.M{"nib": keyToFind}).Decode(&nib)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Map{"error": err.Error()})
			return
		}

		helper.ResponseJSON(w, http.StatusOK, nib)
	}
}
func RegisterNib(conn *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var nib entities.Nib

		if err := json.NewDecoder(r.Body).Decode(&nib); err != nil {
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}
		defer r.Body.Close()
		nib.ID = xid.New().String()
		collection := conn.Database("oss").Collection("nib")
		_, err := collection.InsertOne(r.Context(), &nib)
		if err != nil {
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
		response := map[string]string{"message": "success"}
		helper.ResponseJSON(w, http.StatusOK, response)
	}
}
