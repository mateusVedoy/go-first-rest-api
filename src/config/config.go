package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Init() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func ServeBASEURL() string {

	Init()

	var URLBASE string = os.Getenv("BASE_URL")
	var PORT string = os.Getenv("AVAILABLE_PORT")

	return URLBASE + ":" + PORT
}
