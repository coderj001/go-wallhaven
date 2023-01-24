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

// Purity tags
var PTAGS = map[string]string{
	"sfw":     "100",
	"sketchy": "010",
	"nsfw":    "001",
	"ws":      "110",
	"wn":      "101",
	"sn":      "011",
	"all":     "111",
}

// Categories tags
var CTAGS = map[string]string{
	"all":     "111",
	"anime":   "010",
	"general": "100",
	"people":  "001",
	"ga":      "110",
	"gp":      "101",
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	BASE_CONFIG.APIURL = "https://wallhaven.cc/api/v1/search"
	BASE_CONFIG.APIKEY = os.Getenv("API_KEY")
}

func (c Config) getAPIKey() string {
	var url string
	if c.APIKEY != "" {
		url = c.APIURL + "?apikey=" + c.APIKEY
	} else {
		url = c.APIURL + "?"
	}
	return url
}

func (c Config) GET_URL(query string) string {
	return c.getAPIKey() + "&" + query

}
