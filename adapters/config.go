package adapters

import "os"

type Config struct {
	Host     string
	Username string
	Password string
	DBName   string
	Port     string
}

func GetEnvironments() (*Config, error) {
	c := &Config{
		Host:     os.Getenv("HOST"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		Port:     os.Getenv("POSTGRES_PORT"),
	}

	return c, nil
}
