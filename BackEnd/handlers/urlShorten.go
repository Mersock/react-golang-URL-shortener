package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/Mersock/react-golang-URL-shortener/BackEnd/dbiface"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var (
	v = validator.New()
)

type (
	URL struct {
		OriginalUrl string `json:"OriginalUrl" bson:"OriginalUrl" validate:"required,url"`
		UrlCode     string `json:"UrlCode" bson:"UrlCode"`
		ShortUrl    string `json:"ShortUrl" bson:"ShortUrl"`
	}

	UrlHandler struct {
		Col dbiface.CollectionAPI
	}

	UrlShortenValidator struct {
		validator *validator.Validate
	}
)

func (v *UrlShortenValidator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func insertUrlShortens(ctx context.Context, urlShortens URL, collection dbiface.CollectionAPI) (interface{}, error) {
	res, err := collection.InsertOne(context.Background(), urlShortens)
	if err != nil {
		log.Printf("Unable to insert :%v", err)
		return nil, err
	}
	return res, nil
}

func (h *UrlHandler) CreateUrlShorten(c echo.Context) error {
	var urlShortens URL
	c.Echo().Validator = &UrlShortenValidator{validator: v}
	if err := c.Bind(&urlShortens); err != nil {
		log.Printf("Unable to bind :%v", err)
		return err
	}

	if err := c.Validate(urlShortens); err != nil {
		log.Printf("Unable to validate the urlShorten %+v %v", urlShortens, err)
		return err
	}

	_, err := insertUrlShortens(context.Background(), urlShortens, h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, urlShortens)
}
