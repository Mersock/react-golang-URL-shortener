package config

type Properties struct {
	Port     string `env:"MY_APP_PORT" env-default:"8080"`
	DBHost   string `env:"DB_HOST" env-default:"mongo"`
	DBPort   string `env:"DB_PORT" env-default:"27017"`
	DBName   string `env:"DB_NAME" env-default:"urlShorten"`
	DBUser   string `env:"DB_USER" env-default:"root"`
	DBPass   string `env:"DB_PASS" env-default:"123456"`
	DBColUrl string `env:"DB_Col_Url" env-default:"URL"`
}
