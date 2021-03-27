package handlers

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type UrlShorten struct {
	originalUrl string    `json:"original_url" bson:"original_url"`
	urlCode     string    `json:"url_code" bson:"url_code"`
	shortUrl    string    `json:"short_url" bson:"short_url"`
	createdAt   time.Time `json:"created_at" bson:"created_at"`
	updatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}

func CreateUrlShorten(c echo.Context) error {
	return c.JSON(http.StatusCreated, "create complete")
}
