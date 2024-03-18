package main

import (
	"os"

	"github.com/MrSaveliy/vk-filmoteca/internal/handler"
	"github.com/MrSaveliy/vk-filmoteca/internal/repository"
	"github.com/MrSaveliy/vk-filmoteca/internal/service"
	filmoteca "github.com/MrSaveliy/vk-filmoteca/server"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	_ "github.com/swaggo/http-swagger/example/go-chi/docs"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

// @title vk-filmoteca
// @version 1.0
// @description API Server for Movies Application

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:1323/swagger/doc.json"), //The url pointing to API definition
	))

	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}
	// connect to db
	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed ti initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(filmoteca.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http Server: %s", err.Error())
	} else {
		logrus.Printf("Server is running on port %s", viper.GetString("port"))
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
