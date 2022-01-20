package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type neo4jSettings struct {
	dbHost, dbUser, dbPass string
}

func parseEnvVars() (neo4jSettings, error) {
	// Load the .env file in the current directory
	godotenv.Load()

	var settings neo4jSettings

	if settings.dbHost = os.Getenv("DB_HOST"); settings.dbHost == "" {
		return neo4jSettings{}, fmt.Errorf("DB host must be set")
	}

	if settings.dbUser = os.Getenv("DB_USER"); settings.dbUser == "" {
		return neo4jSettings{}, fmt.Errorf("DB user be set")
	}

	if settings.dbPass = os.Getenv("DB_PASS"); settings.dbPass == "" {
		return neo4jSettings{}, fmt.Errorf("DB pass must be set")
	}

	return settings, nil
}
