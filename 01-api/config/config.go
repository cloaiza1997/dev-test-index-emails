package config

import "os"

type Config struct {
	Port      string
	ZincHost  string
	ZincUser  string
	ZincPass  string
	ZincIndex string
}

var ApiConfig = Config{
	Port:      os.Getenv("EMAIL_INDEX_API_PORT"),
	ZincHost:  os.Getenv("EMAIL_INDEX_ZS_HOST"),
	ZincUser:  os.Getenv("ZINC_FIRST_ADMIN_USER"),
	ZincPass:  os.Getenv("ZINC_FIRST_ADMIN_PASSWORD"),
	ZincIndex: "emails",
}
