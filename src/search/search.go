package search

import (
	"fmt"

	"github.com/coderj001/go-wallheven/src"
)

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
