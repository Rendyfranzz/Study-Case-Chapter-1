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
	r.Get("/nib/{id}", controller.GetNib(conn))
	r.Post("/regisnib", controller.RegisterNib(conn))
	r.Post("/addnews", controller.AddNews(conn))
	r.Get("/getnews", controller.GetNews(conn))
	r.Group(func(r chi.Router) {
		r.Use(authmid.JWTMiddleware)
		r.Post("/logout", controller.Logout(conn))
		r.Get("/test", Index)
		r.Post("/pengajuan", controller.Pengajuan(conn))
		r.Put("/updatepengajuan/{nik}", controller.EditPengajuan(conn))
	})
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}
