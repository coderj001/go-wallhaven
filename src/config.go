package src

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIURL string
	APIKEY string
}

var config Config

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	config.APIURL = "https://wallhaven.cc/api/v1/search"
	config.APIKEY = os.Getenv("API_KEY")
}

func (c Config) get_url() string {
	var url string
	if c.APIKEY != "" {
		url = c.APIURL + "?apikey=" + c.APIKEY
	} else {
		url = c.APIURL + "?"
	}
	return url
}
