package app

import (
	"encoding/json"
	"net/http"

	"github.com/coderj001/go-wallheven/src"
)

func FetchAPI(url string) (src.SearchList, error) {
	resp, err := http.Get(url)
	if err != nil {
		return src.SearchList{}, err
	}
	defer resp.Body.Close()

	var searchList src.SearchList

	err = json.NewDecoder(resp.Body).Decode(&searchList)
	if err != nil {
		return src.SearchList{}, err
	}

	return searchList, nil
}
