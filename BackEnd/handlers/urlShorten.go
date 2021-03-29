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
	"go.mongodb.org/mongo-driver/bson"
)

var (
	v   = validator.New()
	cfg config.Properties
)

type (
	URL struct {
		OriginalUrl string    `json:"originalUrl" bson:"originalUrl" validate:"required,url"`
		UrlCode     string    `json:"urlCode" bson:"urlCode"`
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

func findOriginalUrl(ctx context.Context, collection dbiface.CollectionAPI, filter interface{}) string {
	var shortener URL
	err := collection.FindOne(ctx, filter).Decode(&shortener)
	if err != nil {
		log.Printf("Unable to find OriginalUrl :%v", err)
	}
	updateCounter := bson.M{
		"$set": bson.M{"counter": shortener.Counter + 1},
	}

	err = collection.FindOneAndUpdate(ctx, filter, updateCounter).Decode(&shortener)
	if err != nil {
		log.Printf("Unable to FindOneAndUpdate counter :%v", err)
	}

	return shortener.OriginalUrl
}

func listUrlShortens(ctx context.Context, collection dbiface.CollectionAPI) ([]URL, error) {
	var urlShortens []URL
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Unable to find listUrlShortens :%v", err)
		return nil, err
	}
	err = cursor.All(ctx, &urlShortens)
	if err != nil {
		log.Printf("Unable to read cursor listUrlShortens :%v", err)
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
		log.Printf("Unable to insert urlShorten %v", err)
		return err
	}
	return c.JSON(http.StatusCreated, res)
}

func (h *UrlHandler) RedirectShorten(c echo.Context) error {

	urlCode := c.Param("urlCode")
	originalUrl := findOriginalUrl(context.Background(), h.Col, bson.M{"urlCode": urlCode})
	if originalUrl == "" {
		return c.JSON(http.StatusNotFound, map[string]string{
			"message": "Not Found",
		})
	}
	return c.Redirect(http.StatusMovedPermanently, originalUrl)
}

func (h *UrlHandler) GetUrlShorten(c echo.Context) error {
	urlShorten, err := listUrlShortens(context.Background(), h.Col)
	if err != nil {
		log.Printf("Unable to get list urlShorten %v", err)
		return err
	}
	return c.JSON(http.StatusOK, urlShorten)
}
