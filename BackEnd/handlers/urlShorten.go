package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/Mersock/react-golang-URL-shortener/BackEnd/dbiface"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	v = validator.New()
)

type URL struct {
	ID          primitive.ObjectID `json:"_id" bson:"_id"`
	OriginalUrl string             `json:"original_url" bson:"original_url" validate:"required,url"`
	UrlCode     string             `json:"url_code" bson:"url_code"`
	ShortUrl    string             `json:"short_url" bson:"short_url"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at" bson:"updated_at"`
}

type UrlHandler struct {
	Col dbiface.CollectionAPI
}

type UrlShortenValidator struct {
	validator *validator.Validate
}

func (v *UrlShortenValidator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func insertUrlShortens(ctx context.Context, urlShortens []URL, collection dbiface.CollectionAPI) ([]interface{}, error) {
	var insertedIds []interface{}
	for _, urlShorten := range urlShortens {
		urlShorten.ID = primitive.NewObjectID()
		insertID, err := collection.InsertOne(ctx, urlShorten)
		if err != nil {
			log.Printf("Unable to insert :%v", err)
			return nil, err
		}
		insertedIds = append(insertedIds, insertID.InsertedID)
	}
	return insertedIds, nil
}

func (h *UrlHandler) CreateUrlShorten(c echo.Context) error {
	var urlShortens []URL
	c.Echo().Validator = &UrlShortenValidator{validator: v}
	if err := c.Bind(&urlShortens); err != nil {
		log.Printf("Unable to bind :%v", err)
		return err
	}

	for _, urlShorten := range urlShortens {
		if err := c.Validate(urlShorten); err != nil {
			log.Printf("Unable to validate the urlShorten %+v %v", urlShorten, err)
			return err
		}

	}

	IDs, err := insertUrlShortens(context.Background(), urlShortens, h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, IDs)
}
