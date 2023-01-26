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

// Base Config
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

var SORTING = map[string]string{
	"relevance":  "relevance",
	"random":     "random",
	"date_added": "date_added",
	"favorites":  "favorites",
	"views":      "views",
	"toplist":    "toplist",
	"hot":        "hot",
}

type MetaStruct struct {
	CurrentPage int    `json:"current_page"`
	LastPage    int    `json:"last_page"`
	PerPage     string `json:"per_page"`
	Total       int    `json:"total"`
}

type ThumbsStruct struct {
	Large    string `json:"large"`
	Original string `json:"original"`
	Small    string `json:"small"`
}

type ImageInfo struct {
	Id         string `json:"id"`
	Url        string `json:"url"`
	ShortUrl   string `json:"short_url"`
	Views      int    `json:"views"`
	Favorites  int    `json:"favorites"`
	Source     string `json:"source"`
	Purity     string `json:"purity"`
	Category   string `json:"category"`
	DimensionX int    `json:"dimension_x"`
	DimensionY int    `json:"dimension_y"`
	Resolution string `json:"resolution"`
	Ratio      string `json:"ratio"`
	FileSize   int64  `json:"file_size"`
	FileType   string `json:"file_type"`
	CreatedAt  string `json:"created_at"`
	Colors     []string
	Path       string       `json:"path"`
	Thumbs     ThumbsStruct `json:"thumbs"`
}

type SearchList struct {
	Data []ImageInfo
	Meta MetaStruct
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

func (c Config) GetURL(query string) string {
	return c.getAPIKey() + query

}
