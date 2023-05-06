package main

import (
	"context"
	"log"

	"github.com/Rendyfranzz/Study-Case-Chapter-1/config"
	"github.com/Rendyfranzz/Study-Case-Chapter-1/handlers"
	"github.com/Rendyfranzz/Study-Case-Chapter-1/pkg"
	"github.com/Rendyfranzz/Study-Case-Chapter-1/repo"
	"github.com/Rendyfranzz/Study-Case-Chapter-1/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	conn, closeConnection, err := config.Connect("mongodb://app_user:app_password@localhost:27017/admin")
	if err != nil {
		log.Fatal(err)
	}

	defer closeConnection(context.Background())

	jwt := pkg.JWT{}

	userRepo := repo.NewUserRepo(conn)
	nibRepo := repo.NewNIBRepo(conn)

	authService := service.NewAuthService(userRepo, jwt)
	nibService := service.NewNIBService(nibRepo)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/api/login", handlers.Login(authService))
	e.POST("/api/register", handlers.Register(authService))

	group := e.Group("/api", JWTMiddleware(jwt))
	group.POST("/logout", handlers.Logout())
	group.GET("/nib/:id", handlers.GetNIB(nibService))
	group.POST("/nib", handlers.CreateNIB(nibService))

	e.Start(":6000")
}
