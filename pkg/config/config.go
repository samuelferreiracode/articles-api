package config

import "os"

type Config struct {
	MongoDBURI string
	Database   string
}

func LoadConfig() *Config {
	return &Config{
		MongoDBURI: os.Getenv("MONGODB_URI"),
		Database:   os.Getenv("DATABASE_NAME"),
	}
}
