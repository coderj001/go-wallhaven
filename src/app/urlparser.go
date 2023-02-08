package app

import (
	"fmt"
	"net/url"

	"github.com/coderj001/go-wallheven/src"
)

type error interface {
	Error() string
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

func (p *Param) String() string {
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

func (p *Param) getFullURL() string {
	return src.BASE_CONFIG.GetURL(p.String())
}

func GetFullURL(
	page int,
	categories, purity, sorting, color, query string,
) string {
	params := Param{
		Page:        page,
		CTage:       src.CTAGS[categories],
		PTage:       src.PTAGS[purity],
		Sorting:     src.SORTING[sorting],
		Color:       color,
		Query:       url.QueryEscape(query),
		Resolutions: []Resolution{},
	}
	return params.getFullURL()
}
