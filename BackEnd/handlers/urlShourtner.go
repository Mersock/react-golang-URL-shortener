package handlers

import (
	"time"
)

type urlShorten struct {
	originalUrl string    `json:"original_url" bson:"original_url"`
	urlCode     string    `json:"url_code" bson:"url_code"`
	shortUrl    string    `json:"short_url" bson:"short_url"`
	createdAt   time.Time `json:"created_at" bson:"created_at"`
	updatedAt   time.Time `json:"updated_at" bson:"updated_at"`
}
