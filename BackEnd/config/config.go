package config

type Properties struct {
	Port   string `env:"MY_APP_PORT" env-default:"8080"`
	Host   string `env:"HOST" env-default:"localhost"`
	DBHost string `env:"DB_HOST" env-default:"localhost"`
	DBPort string `env:"DB_PORT" env-default:"27017"`
	DBName string `env:"DB_NAME" env-default:"url_shorten"`
	DBUser string `env:"DB_USER" env-default:"root"`
	DBPass string `env:"DB_PASS" env-default:"12345"`
}
