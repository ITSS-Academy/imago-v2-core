package main

import (
	"context"
	"firebase.google.com/go/v4"
	"fmt"
	"github.com/itss-academy/imago/core/domain/auth"
	"github.com/itss-academy/imago/core/domain/post"
	authPkgDelivery "github.com/itss-academy/imago/core/internal/auth/delivery"
	authPkgInterop "github.com/itss-academy/imago/core/internal/auth/interop"
	authPkgRepo "github.com/itss-academy/imago/core/internal/auth/repo"
	authPkgUcase "github.com/itss-academy/imago/core/internal/auth/ucase"
	postPkgDelivery "github.com/itss-academy/imago/core/internal/post/delivery"
	postPkgInterop "github.com/itss-academy/imago/core/internal/post/interop"
	postPkgRepo "github.com/itss-academy/imago/core/internal/post/repo"
	postPkgUcase "github.com/itss-academy/imago/core/internal/post/ucase"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"google.golang.org/api/option"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	viper.SetConfigName("current") // name of config file (without extension)
	viper.SetConfigType("yaml")    // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")       // optionally look for config in the working directory
	err := viper.ReadInConfig()    // Find and read the config file
	if err != nil {                // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	// Firebase init
	opt := option.WithCredentialsFile(viper.GetString("firebase.credential"))
	firebaseApp, err := firebase.NewApp(context.Background(), &firebase.Config{
		ProjectID: viper.GetString("firebase.projectid"),
	}, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	authClient, err := firebaseApp.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	// database init
	db, err := gorm.Open(postgres.Open(viper.GetString("database.dsn")), &gorm.Config{})
	if err != nil {
		log.Fatalf("error initializing database: %v\n", err)
	}

	e := echo.New()
	// add middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// dependency injection by manual
	var authRepo auth.AuthRepository
	var authUsecase auth.AuthUseCase
	var authInterop auth.AuthInterop

	var postRepo post.PostRepository
	var postUsecase post.PostUseCase
	var postInterop post.PostInterop

	postRepo = postPkgRepo.NewPostRepository(db)
	postUsecase = postPkgUcase.NewPostUseCase(postRepo)
	postInterop = postPkgInterop.NewPostBaseInterop(postUsecase, authUsecase)

	authRepo = authPkgRepo.NewAuthRepository(db)
	authUsecase = authPkgUcase.NewAuthUseCase(authRepo, authClient)
	authInterop = authPkgInterop.NewAuthInterop(authUsecase)

	// add routes

	authApi := e.Group("/v2/auth")
	authPkgDelivery.NewAuthHttpDelivery(authApi, authInterop)

	postApi := e.Group("/v2/post")
	postPkgDelivery.NewPostHttpDelivery(postApi, postInterop)

	// start server
	_ = e.Start(fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port")))
}
