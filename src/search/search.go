package search

import (
	"fmt"

	"github.com/coderj001/go-wallheven/src"
)

type MetaStruct struct {
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	PerPage     int `json:"per_page"`
	Total       int `json:"total"`
}

type ThumbsStruct struct {
	Large    string `json:"large"`
	Original string `json:"original"`
	Small    string `json:"small"`
}

type ImageInfo struct {
	Id         string  `json:"id"`
	Url        string  `json:"url"`
	ShortUrl   string  `json:"short_url"`
	Views      int     `json:"views"`
	Favorites  int     `json:"favorites"`
	Source     string  `json:"source"`
	Purity     string  `json:"purity"`
	Category   string  `json:"category"`
	DimensionX int     `json:"dimension_x"`
	DimensionY int     `json:"dimension_y"`
	Resolution string  `json:"resolution"`
	Ratio      float64 `json:"ratio"`
	FileSize   int64   `json:"file_size"`
	FileType   string  `json:"file_type"`
	CreatedAt  string  `json:"created_at"`
	Colors     []string
	Path       string       `json:"path"`
	Thumbs     ThumbsStruct `json:"thumbs"`
}

type SearchList struct {
	Data []ImageInfo
	Meta MetaStruct
}

type Resolution struct {
	Width  int
	Height int
}

type Param struct {
	Page        int
	Categories  string
	Tag         string
	Resolutions []Resolution
}

func Resolutions(rs []Resolution) string {
	paramStr := ""

	for i, r := range rs {
		if i > 0 {
			paramStr = fmt.Sprintf("%s,%dx%d", paramStr, r.Width, r.Height)
		} else {
			paramStr = fmt.Sprintf("resolutions=%dx%d", r.Width, r.Height)
		}
	}
	return paramStr
}

func (p Param) String() string {
	paramStr := ""
	rsStr := Resolutions(p.Resolutions)

	if rsStr != "" {
		paramStr = rsStr
	}

	paramStr = fmt.Sprintf("%s&q=%s", paramStr, p.Tag)
	paramStr = fmt.Sprintf("%s&q=%d", paramStr, p.Page)

	url := src.BASE_CONFIG.GET_URL(paramStr)

	return url
}

func Search(p Param) (*SearchList, error) {}
