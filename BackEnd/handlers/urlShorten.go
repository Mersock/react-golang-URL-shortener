package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Mersock/react-golang-URL-shortener/BackEnd/config"
	"github.com/Mersock/react-golang-URL-shortener/BackEnd/dbiface"
	"github.com/Mersock/react-golang-URL-shortener/BackEnd/helper"
	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
)

var (
	v   = validator.New()
	cfg config.Properties
)

type (
	URL struct {
		OriginalUrl string    `json:"originalUrl" bson:"OriginalUrl" validate:"required,url"`
		UrlCode     string    `json:"urlCode" bson:"UrlCode"`
		ShortUrl    string    `json:"shortUrl" bson:"ShortUrl"`
		Expires     time.Time `json:"expires" bson:"expires"`
		Counter     int       `json:"counter" bson:"counter"`
	}

	UrlHandler struct {
		Col dbiface.CollectionAPI
	}

	UrlShortenValidator struct {
		validator *validator.Validate
	}
)

func init() {
	if err := cleanenv.ReadEnv(&cfg); err != nil {
		log.Fatalf("Configuration env cannot read %v", err)
	}
}

func (v *UrlShortenValidator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func insertUrlShortens(ctx context.Context, urlShortens URL, collection dbiface.CollectionAPI) (interface{}, error) {
	t := time.Now()
	expires := t.Add(time.Hour)
	urlShortens.Expires = expires

	strCode := helper.RandURLCode(8, 1, 1)
	urlShortens.UrlCode = strCode

	shortUrl := fmt.Sprintf("%s://%s/%s", cfg.URLSchema, cfg.URLPrefix, strCode)
	urlShortens.ShortUrl = shortUrl

	urlShortens.Counter = 0

	_, err := collection.InsertOne(context.Background(), urlShortens)
	if err != nil {
		log.Printf("Unable to insert :%v", err)
		return nil, err
	}

	return urlShortens, nil
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

	res, err := insertUrlShortens(context.Background(), urlShortens, h.Col)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, res)
}
