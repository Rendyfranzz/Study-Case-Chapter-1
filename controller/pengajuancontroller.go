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

func Pengajuan(conn *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var pengajuan entities.Pengajuan

		if err := json.NewDecoder(r.Body).Decode(&pengajuan); err != nil {
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}
		defer r.Body.Close()
		pengajuan.ID = xid.New().String()

		collection := conn.Database("oss").Collection("pengajuan")
		_, err := collection.InsertOne(r.Context(), &pengajuan)
		if err != nil {
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
		response := map[string]string{"message": "success"}
		helper.ResponseJSON(w, http.StatusOK, response)
	}
}

func EditPengajuan(conn *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keyToFind := chi.URLParam(r, "nik")
		var pengajuan entities.Pengajuan

		if err := json.NewDecoder(r.Body).Decode(&pengajuan); err != nil {
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}
		defer r.Body.Close()
		collection := conn.Database("oss").Collection("pengajuan")
		edited := bson.M{"npwp": pengajuan.Npwp, "no_bpjs": pengajuan.No_bpjs}
		_, err := collection.UpdateOne(r.Context(), bson.M{"nik": keyToFind}, bson.M{"$set": edited})
		if err != nil {
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
		response := map[string]string{"message": "success"}
		helper.ResponseJSON(w, http.StatusOK, response)
	}
}
