package controller

import (
	"crud/config"
	"crud/helper"
	"crud/model"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/rs/xid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type Map map[string]interface{}

func Register(conn *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusBadRequest, response)
			return
		}
		defer r.Body.Close()
		user.ID = xid.New().String()
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hashPassword)
		collection := conn.Database("oss").Collection("users")
		_, err := collection.InsertOne(r.Context(), &user)
		if err != nil {
			response := map[string]string{"message": err.Error()}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
		response := map[string]string{"message": "success"}
		helper.ResponseJSON(w, http.StatusOK, response)
	}
}

func Login(conn *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user model.User

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(Map{"error": err.Error()})
			return
		}
		defer r.Body.Close()
		var temp model.User
		collection := conn.Database("oss").Collection("users")

		err := collection.FindOne(r.Context(), bson.M{"name": user.Name}).Decode(&temp)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Map{"error": err.Error()})
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(temp.Password), []byte(user.Password)); err != nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(Map{"error": err.Error()})
			return
		}
		expTime := time.Now().Add(time.Minute * 2)
		claims := &config.Claims{
			User: user.Name,
			RegisteredClaims: jwt.RegisteredClaims{
				ExpiresAt: jwt.NewNumericDate(expTime),
			},
		}
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenString, err := token.SignedString(config.JWT_KEY)
		if err != nil {
			// If there is an error in creating the JWT return an internal server error
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Path:     "/",
			Value:    tokenString,
			Expires:  expTime,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		})
		response := map[string]string{"message": "login berhasil"}
		helper.ResponseJSON(w, http.StatusOK, response)
	}

}

func Logout(conn *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// hapus token yang ada di cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Path:     "/",
			Value:    "",
			HttpOnly: true,
			MaxAge:   -1,
		})

		response := map[string]string{"message": "logout berhasil"}
		helper.ResponseJSON(w, http.StatusOK, response)
	}

}
func Test(conn *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// hapus token yang ada di cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Path:     "/",
			Value:    "",
			HttpOnly: true,
			MaxAge:   -1,
		})

		response := map[string]string{"message": "logout berhasil"}
		helper.ResponseJSON(w, http.StatusOK, response)
	}

}
