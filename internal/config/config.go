package config


import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)


type Config struct {
	Port string
	Database string
	Username string
	Password string
}

func Load() (*Config, error) {
	_=godotenv.Load()

	port := os.Getenv("POSTGRES_PORT")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	db := os.Getenv("POSTGRES_DB")
	user_name:=os.Getenv("MUSIC_USERNAME")
	pass:=os.Getenv("MUSIC_PASSWORD")

	if port == "" || user == "" || password == "" || db == "" {
		return nil, fmt.Errorf("missing required database environment variables")
	}

	if user_name == "" || pass == "" {
	return nil, fmt.Errorf("missing music server credentials")
}

	dbURL := fmt.Sprintf(
		"postgres://%s:%s@localhost:%s/%s",
		user,
		password,
		port,
		db,
	)

	return &Config{Port: "8080", Database: dbURL, Username: user_name, Password: pass}, nil
	
}