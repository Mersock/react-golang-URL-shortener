package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Mersock/react-golang-URL-shortener/BackEnd/config"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	c   *mongo.Client
	db  *mongo.Database
	col *mongo.Collection
	cfg config.Properties
	h   UrlHandler
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

func Test(t *testing.T) {
	t.Run("test Create UrlShorten", func(t *testing.T) {
		body := `
		[{
			"original_url": "https://echo.labstack.com/guide/request/",
			"url_code": "sdfwe382",
			"short_url": "d82ewf"
		},
		{
			"original_url": "https://echo.labstack.com/guide/request/",
			"url_code": "23fewf",
			"short_url": "safsdf"
		},
		{
			"original_url": "https://echo.labstack.com/guide/request/",
			"url_code": "234ks",
			"short_url": "asdf23"
		}]
		`
		req := httptest.NewRequest("POST", "/api/item", strings.NewReader(body))
		res := httptest.NewRecorder()
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		e := echo.New()
		c := e.NewContext(req, res)
		h.Col = col
		err := h.CreateUrlShorten(c)
		assert.Nil(t, err)
	})
}
