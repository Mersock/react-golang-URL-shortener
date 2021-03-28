package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Mersock/react-golang-URL-shortener/BackEnd/config"
	"github.com/Mersock/react-golang-URL-shortener/BackEnd/handlers"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	db  *mongo.Database
	col *mongo.Collection
	cfg config.Properties
)

func init() {
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Configuration env cannot read %v", err)
	}
	connectURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort)
	c, err := mongo.Connect(context.Background(), options.Client().ApplyURI(connectURI))
	if err != nil {
		log.Fatalf("Unable to connect mongo %v", err)
	}
	db = c.Database(cfg.DBName)
	col = db.Collection(cfg.DBColUrl)
}

func main() {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	h := handlers.UrlHandler{Col: col}
	e.POST("/api/urlShorten", h.CreateUrlShorten)
	e.GET("/:urlCode", h.RedirectShorten)
	e.Logger.Infof("Listen on $s:%s", cfg.DBHost, cfg.Port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", cfg.Port)))
}
