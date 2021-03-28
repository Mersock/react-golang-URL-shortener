package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Mersock/react-golang-URL-shortener/BackEnd/dbiface"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UrlShorten struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	originalUrl string             `json:"original_url" bson:"original_url"`
	urlCode     string             `json:"url_code" bson:"url_code"`
	shortUrl    string             `json:"short_url" bson:"short_url"`
	createdAt   time.Time          `json:"created_at" bson:"created_at"`
	updatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type UrlHandler struct {
	Col dbiface.CollectionAPI
}

func insertUrlShortens(ctx context.Context, urlShortens []UrlShorten, collection dbiface.CollectionAPI) ([]interface{}, error) {
	var insertedIds []interface{}
	for _, urlShorten := range urlShortens {
		urlShorten.ID = primitive.NewObjectID()
		insertID, err := collection.InsertOne(ctx, urlShorten)
		if err != nil {
			log.Printf("Unable to insert :%v", err)
			return nil, err
		}
		fmt.Println("insertID", insertID)
		insertedIds = append(insertedIds, insertID.InsertedID)
	}
	return insertedIds, nil
}

func (h *UrlHandler) CreateUrlShorten(c echo.Context) error {
	var urlShortens []UrlShorten
	fmt.Println(urlShortens)
	if err := c.Bind(&urlShortens); err != nil {
		log.Printf("Unable to bind :%v", err)
		return err
	}
	IDs, err := insertUrlShortens(context.Background(), urlShortens, h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, IDs)
}
