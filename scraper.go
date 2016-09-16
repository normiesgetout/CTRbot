// Scraper
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"os"
)

type _4chanAPI struct {
	Threads []struct {
		Posts []struct {
			ID          string `json:"id"`
			Closed      int64  `json:"closed"`
			Com         string `json:"com"`
			Country     string `json:"country"`
			CountryName string `json:"country_name"`
			Ext         string `json:"ext"`
			Filename    string `json:"filename"`
			Fsize       int64  `json:"fsize"`
			H           int64  `json:"h"`
			Images      int64  `json:"images"`
			Md5         string `json:"md5"`
			Name        string `json:"name"`
			No          int64  `json:"no"`
			Now         string `json:"now"`
			Replies     int64  `json:"replies"`
			Resto       int64  `json:"resto"`
			SemanticURL string `json:"semantic_url"`
			Sticky      int64  `json:"sticky"`
			Sub         string `json:"sub"`
			Tim         int64  `json:"tim"`
			Time        int64  `json:"time"`
			TnH         int64  `json:"tn_h"`
			TnW         int64  `json:"tn_w"`
			W           int64  `json:"w"`
		} `json:"posts"`
	} `json:"threads"`
}

func Scrape(url *string, file *os.File) {
	w := bufio.NewWriter(file)

	response, err := http.Get(*url)

	if err != nil {
		fmt.Println("Could not load URL!")
		os.Exit(1)
	}

	b, _ := ioutil.ReadAll(response.Body)
	// b is a char array ... ahem ... I mean slice
	// fucking naming conventions

	defer response.Body.Close()

	var chanfour _4chanAPI
	if err := json.Unmarshal(b, &chanfour); err != nil {
		fmt.Println("Not getting the correct API. It's Hiro's fault!")
		fmt.Println(err)
		os.Exit(1)
	}

	for i := range chanfour.Threads {
		for j := range chanfour.Threads[i].Posts {
			if chanfour.Threads[i].Posts[j].Sticky == 0 {
				_, err := w.WriteString(html.UnescapeString(StripTags(chanfour.Threads[i].Posts[j].Com)))
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
	fmt.Println("Cancer successfully loaded!")
	file.Sync()
}
