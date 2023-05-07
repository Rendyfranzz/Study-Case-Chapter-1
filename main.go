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
	// conn, closeConnection, err := config.Connect("mongodb://app_user:app_password@localhost:27017/admin")
	conn, closeConnection, err := config.Connect("mongodb://localhost:27017/oss")
	if err != nil {
		log.Fatal(err)
	}

	defer closeConnection(context.Background())

	jwt := pkg.JWT{}

	userRepo := repo.NewUserRepo(conn)
	nibRepo := repo.NewNIBRepo(conn)
	submissionRepo := repo.NewSubmissionRepo(conn)
	newsRepo := repo.NewNewsRepo(conn)

	authService := service.NewAuthService(userRepo, jwt)
	nibService := service.NewNIBService(nibRepo)
	submissionService := service.NewSubmissionService(submissionRepo)
	newsService := service.NewNewsService(newsRepo)

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/api/login", handlers.Login(authService))
	e.POST("/api/register", handlers.Register(authService))
	e.GET("/api/news", handlers.GetNews(newsService))
	e.GET("/api/nib/:id", handlers.GetNIB(nibService))
	e.POST("/nib", handlers.CreateNIB(nibService))

	group := e.Group("/api", JWTMiddleware(jwt))
	group.POST("/logout", handlers.Logout())
	group.POST("/submission", handlers.CreateSubmissions(submissionService))
	group.PUT("/submission/:nik", handlers.EditSubmission(submissionService))

	e.Start(":8080")
}
