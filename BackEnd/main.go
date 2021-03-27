package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Mersock/react-golang-URL-shortener/BackEnd/config"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	c   *mongo.Client
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
		log.Fatalf("Unable to connect mongo $v", err)
	}
	db = c.Database(cfg.DBName)
	col = db.Collection(cfg.DBColUrl)
}

func main() {
	e := echo.New()
	e.POST("/api/item", handlers)
	e.Logger.Infof("Listen on $s:%s", cfg.DBHost, cfg.Port)
	e.Logger.Fatal(e.Start(fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)))
}
