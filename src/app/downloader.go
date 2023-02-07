package app

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/coderj001/go-wallheven/src"
)

const numGoroutines = src.BASE_CONFIG.NUMGOROUTINES
const dir = src.BASE_CONFIG.DIR

func DownloadImgs(searchList src.SearchList) error {
	var wg sync.WaitGroup
	wg.Add(len(searchList.Data))
	for i := 0; i < numGoroutines; i++ {
		go func(i int){
			for j := i; i<len(searchList.Data); j+=numGoroutines{
				return nil
			}
		}
	}
	return nil
}

func downloadImage(imgInfo src.ImageInfo, wg *sync.WaitGroup){
	defer wg.Done()
	
	resp, err := http.Get(imgInfo.Path)
	if err != nil {
		log.Printf("Error while downloading (%s): %s\n", imgInfo.Url, err)
		return
	}
	defer resp.Body.Close()
	
	// check the dir is exists or not
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist){
		err:= os.Mkdir(dir, os.ModePerm)
		if err != nil {
			log.Printf("Error while creating dir (%s): %s\n", dir, err)
			return
		}
	}

	// create the file to store download content
	
	filename := strings.Join([]string{dir, strings.Join([]string{imgInfo.Id, filepath.Ext(imgInfo.Path)}, "")}, "/")
	fileOut, err := os.Create(filename)
	if err != nil {
		log.Printf("Unable to create file (%s): %s\n", imgInfo.Id, err)
		return
	}
	defer fileOut.Close()
	
	_, err = io.Copy(fileOut, resp.Body)
	if err != nil {
		log.Printf("Error while writing content to file: %s\n", err)
		return
	}
	
	fmt.Printf("File downloaded successfully: %s\n", imgInfo.ShortUrl)
}