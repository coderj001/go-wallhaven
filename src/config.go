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

var BASE_CONFIG Config

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	BASE_CONFIG.APIURL = "https://wallhaven.cc/api/v1/search"
	BASE_CONFIG.APIKEY = os.Getenv("API_KEY")
}

func (c Config) getBaseUrl() string {
	var url string
	if c.APIKEY != "" {
		url = c.APIURL + "?apikey=" + c.APIKEY
	} else {
		url = c.APIURL + "?"
	}
	return url
}

func (c Config) GET_URL(query string) string {
	return c.getBaseUrl() + "&" + query

}
