package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/coderj001/go-wallhaven/src"
)

func FetchAPI(url string) (src.SearchList, error) {
	resp, err := http.Get(url)
	if err != nil {
		return src.SearchList{
				Data: []src.ImageInfo{},
				Meta: src.MetaStruct{},
			},
			fmt.Errorf("error making HTTP request: %w", err)
	}
	defer resp.Body.Close()

	var searchList src.SearchList

	err = json.NewDecoder(resp.Body).Decode(&searchList)
	if err != nil {
		return src.SearchList{
				Data: []src.ImageInfo{},
				Meta: src.MetaStruct{},
			},
			fmt.Errorf("error making HTTP request: %w", err)
	}

	return searchList, nil
}
