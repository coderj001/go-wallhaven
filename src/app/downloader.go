package app

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/coderj001/go-wallheven/src"
)

func Downloader(searchList *src.SearchList) {
	var wg sync.WaitGroup
	wg.Add(len(searchList.Data))
	completed := make(chan bool)

	for i := 0; i < len(searchList.Data); i++ {
		go downloadImage(&searchList.Data[i], &wg, completed)
	}

	// Code for spinner - it takes
	spinner := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	spinnerLen := 10
	spinnerIndex := 0
	completedCount := 0
	timeWait := 100

	for completedCount < len(searchList.Data) {
		select {
		case <-completed:
			completedCount++
			spinnerIndex = (spinnerIndex + 1) % spinnerLen
			fmt.Printf(
				"\rDownloading... %s %d/%d",
				spinner[spinnerIndex],
				completedCount,
				len(searchList.Data))
		default:
			time.Sleep(time.Duration(timeWait) * time.Millisecond)
			spinnerIndex = (spinnerIndex + 1) % spinnerLen
			fmt.Printf(
				"\rDownloading... %s %d/%d",
				spinner[spinnerIndex],
				completedCount,
				len(searchList.Data))
		}
	}

	wg.Wait()
	close(completed)
}

func downloadImage(imgInfo *src.ImageInfo, wg *sync.WaitGroup, completed chan bool) {
	defer wg.Done()
	dir := src.BASE_CONFIG.DIR

	resp, err := http.Get(imgInfo.Path)
	if err != nil {
		log.Printf("error while downloading (%s): %s\n", imgInfo.Url, err)
		return
	}
	defer resp.Body.Close()

	// check the dir is exists or not
	if _, err = os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			log.Printf("error while creating dir (%s): %s\n", dir, err)
			return
		}
	}

	// create the file to store download content
	imgName := imgInfo.Id + filepath.Ext(imgInfo.Path)
	filename := dir + "/" + imgName
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

	// fmt.Printf("File downloaded successfully: %s\n", imgInfo.ShortUrl) //nolint
	completed <- true
}
