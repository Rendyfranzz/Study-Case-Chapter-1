package main

import (
	"context"
	"crud/authmid"
	"crud/controller"
	"crud/db"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	conn, closeConnection, err := db.Connect("mongodb://localhost:27017/oss")
	if err != nil {
		log.Fatal(err)
	}

	defer closeConnection(context.Background())

	log.Println("Connection ready!!!")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Post("/register", controller.Register(conn))
	r.Post("/login", controller.Login(conn))
	r.Group(func(r chi.Router) {
		r.Use(authmid.JWTMiddleware)
		r.Post("/logout", controller.Logout(conn))
		r.Get("/test", Index)
	})
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
