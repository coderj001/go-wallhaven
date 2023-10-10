//nolint:tagliatelle,stylecheck
// This line is ignoring lint checks because it is necessary for my use case.

package src

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	APIURL   string
	APIKEY   string
	DIR      string
	USERNAME string
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
	var (
		apikey   string
		dir      string
		username string
	)
	if err := godotenv.Load(); err == nil {
		log.Println("Using config file: .env")
		apikey = os.Getenv("API_KEY")
		dir = os.Getenv("DIR")
		username = os.Getenv("USER_NAME")
	} else {
		viper.SetConfigFile(fmt.Sprintf("%s/.go-wallhaven", os.Getenv("HOME")))
		viper.SetConfigType("json")
		if err := viper.ReadInConfig(); err != nil {
			log.Printf("Error reading config file, %s\n", err)
		}
		apikey = viper.GetString("API_KEY")
		dir = viper.GetString("DIR")
		username = viper.GetString("USER_NAME")
	}

	BASE_CONFIG.APIURL = "https://wallhaven.cc/api/v1/search"
	BASE_CONFIG.APIKEY = apikey
	BASE_CONFIG.USERNAME = username
	if dir == "" {
		BASE_CONFIG.DIR = "./tmp"
	} else {
		BASE_CONFIG.DIR = dir
	}
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

func (c Config) getUserName() (string, error) {
	if c.USERNAME != "" {
		return c.USERNAME, nil
	} else {
		return "", errors.New("Username not found")
	}
}

func (c Config) GetURL(query string) string {
	return c.getAPIKey() + query
}
