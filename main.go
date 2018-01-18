package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/reujab/wallpaper"
)

const (
	bingDomain = "https://www.bing.com"
)

func main() {
	urlPic, err := downloadPhoto()
	if err != nil {
		log.Fatal(err)
	}
	if err = wallpaper.SetFromURL(bingDomain + urlPic); err != nil {
		log.Fatal(err)
	}
}

type bingResponse struct {
	Images []struct {
		URL           string        `json:"url"`
	} `json:"images"`
}

func downloadPhoto() (urlPic string, err error) {
	resp, err := http.Get(bingDomain + "/HPImageArchive.aspx?format=js&idx=0&n=1")
	if err != nil {
		return
	}
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Error status request: %s", resp.Status)
		return
	}

	respJSON := new(bingResponse)

	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(respJSON); err != nil {
		return
	}

	if len(respJSON.Images) == 0 {
		err = fmt.Errorf("error json response. Num of images is 0")
		return
	}

	urlPic = respJSON.Images[0].URL
	return
}
