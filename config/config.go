package config

import (
	"errors"
	"os"
)

type EnvVariables struct {
	MONGO_URI     string
	DATABASE_NAME string
}

func LoadEnv() (*EnvVariables, error) {
	mongoURI := os.Getenv("MONGO_URI")
	databaseName := os.Getenv("DATABASE_NAME")
	if mongoURI == "" {
		return nil, errors.New("missing required environment variable: MONGO_URI")
	}
	if databaseName == "" {
		return nil, errors.New("missing required environment variable: DATABASE_NAME")
	}
	return &EnvVariables{
		MONGO_URI:     mongoURI,
		DATABASE_NAME: databaseName,
	}, nil
}
