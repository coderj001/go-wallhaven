package app

import (
	"fmt"
	"net/url"

	"github.com/coderj001/go-wallheven/src"
)

type error interface {
	Error() string
}

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
	Query       string
	CTage       string
	PTage       string
	Sorting     string
	Color       string
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

	if p.Query != "" {
		paramStr = fmt.Sprintf("%s&q=%s", paramStr, p.Query)
	}
	if p.Page != 0 {
		paramStr = fmt.Sprintf("%s&page=%d", paramStr, p.Page)
	}
	if p.CTage != "" {
		paramStr = fmt.Sprintf("%s&categories=%s", paramStr, p.CTage)
	}
	if p.PTage != "" {
		paramStr = fmt.Sprintf("%s&purity=%s", paramStr, p.PTage)
	}
	if p.Sorting != "" {
		paramStr = fmt.Sprintf("%s&sorting=%s", paramStr, p.Sorting)
	}
	if p.Color != "" {
		paramStr = fmt.Sprintf("%s&colors=%s", paramStr, p.Color)
	}

	return paramStr
}

func (p Param) getFullURL() string {
	return src.BASE_CONFIG.GetURL(p.String())
}

func GetFullURL(page int, categories string, purity string, sorting string, color string, query string) (string, error) {
	params := Param{
		Page:        page,
		CTage:       src.CTAGS[categories],
		PTage:       src.PTAGS[purity],
		Sorting:     src.SORTING[sorting],
		Color:       color,
		Query:       url.QueryEscape(query),
		Resolutions: []Resolution{},
	}
	return params.getFullURL(), nil
}
