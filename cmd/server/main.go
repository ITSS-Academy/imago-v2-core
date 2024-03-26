package main

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"fmt"
	"github.com/itss-academy/imago/core/domain/Report"

	"github.com/itss-academy/imago/core/domain/auth"
	"github.com/itss-academy/imago/core/domain/comment"
	"github.com/itss-academy/imago/core/domain/profile"
	"github.com/itss-academy/imago/core/domain/post"
	authPkgDelivery "github.com/itss-academy/imago/core/internal/auth/delivery"
	authPkgInterop "github.com/itss-academy/imago/core/internal/auth/interop"
	authPkgRepo "github.com/itss-academy/imago/core/internal/auth/repo"
	authPkgUcase "github.com/itss-academy/imago/core/internal/auth/ucase"
	reportPkgDelivery "github.com/itss-academy/imago/core/internal/report/delivery"
	reportPkgInterop "github.com/itss-academy/imago/core/internal/report/interop"
	reportPkgRepo "github.com/itss-academy/imago/core/internal/report/repo"
	reportPkgUcase "github.com/itss-academy/imago/core/internal/report/usecase"

	profilePkgDelivery "github.com/itss-academy/imago/core/internal/profile/delivery"
	profilePkgInterop "github.com/itss-academy/imago/core/internal/profile/interop"
	profilePkgRepo "github.com/itss-academy/imago/core/internal/profile/repo"
	profilePkgUcase "github.com/itss-academy/imago/core/internal/profile/ucase"

	commentPkgDelivery "github.com/itss-academy/imago/core/internal/comment/delivery"
	commentPkgInterop "github.com/itss-academy/imago/core/internal/comment/interop"
	commentPkgRepo "github.com/itss-academy/imago/core/internal/comment/repo"
	commentPkgUcase "github.com/itss-academy/imago/core/internal/comment/ucase"
	"log"

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

	var reportRepo Report.ReportRepository
	var reportUsecase Report.ReportUseCase
	var reportInterop Report.ReportInterop
	var profileRepo profile.ProfileRepository
	var profileUsecase profile.ProfileUseCase
	var profileInterop profile.ProfileInterop
	var postRepo post.PostRepository
	var postUsecase post.PostUseCase
	var postInterop post.PostInterop

	postRepo = postPkgRepo.NewPostRepository(db)
	postUsecase = postPkgUcase.NewPostUseCase(postRepo)
	postInterop = postPkgInterop.NewPostBaseInterop(postUsecase, authUsecase)

	authRepo = authPkgRepo.NewAuthRepository(db)
	authUsecase = authPkgUcase.NewAuthUseCase(authRepo, authClient)
	authInterop = authPkgInterop.NewAuthInterop(authUsecase)

	reportRepo = reportPkgRepo.NewReportRepository(db)
	reportUsecase = reportPkgUcase.NewReportUseCase(reportRepo)
	reportInterop = reportPkgInterop.NewReportInterop(reportUsecase, authUsecase)
	profileRepo = profilePkgRepo.NewProfileRepository(db)
	profileUsecase = profilePkgUcase.NewProfileUseCase(profileRepo)
	profileInterop = profilePkgInterop.NewProfileInterop(profileUsecase, authUsecase)

	var commentRepo comment.CommentRepository
	var commentUsecase comment.CommentUseCase
	var commentInterop comment.CommentInterop

	commentRepo = commentPkgRepo.NewCommentRepository(db)
	commentUsecase = commentPkgUcase.NewCommentUseCase(commentRepo)
	commentInterop = commentPkgInterop.NewCommentInterop(commentUsecase, authUsecase)

	// add routes

	authApi := e.Group("/v2/auth")
	profileApi := e.Group("/v2/profile")
	commentApi := e.Group("/v2/comment")
	authPkgDelivery.NewAuthHttpDelivery(authApi, authInterop)
	profilePkgDelivery.NewProfileHttpDelivery(profileApi, profileInterop)
	commentPkgDelivery.NewCommentHttpDelivery(commentApi, commentInterop)


	postApi := e.Group("/v2/post")
	postPkgDelivery.NewPostHttpDelivery(postApi, postInterop)

	reportApi := e.Group("/v2/report")
	reportPkgDelivery.NewReportHttpDeliver(reportApi, reportInterop)


	// start server
	_ = e.Start(fmt.Sprintf("%s:%s", viper.GetString("server.host"), viper.GetString("server.port")))
}
