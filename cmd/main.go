package main

import (
	"fmt"

	parser "github.com/donavanbrooks/soundscraper/internal"
)

func main() {
	//testURL := "https://www.hotnewhiphop.com/mixtapes/"

	var website parser.WebsiteTemplate
	website.Name = "hotnewhiphop"
	website.URL = "https://www.hotnewhiphop.com/mixtapes/"
	website.Base = parser.HTMLElement{Tag: "div", Class: "grid-item song"}
	website.Artist = parser.HTMLElement{Tag: "em", Class: "default-artist"}
	website.Title = parser.HTMLElement{Tag: "a", Class: "cover-title grid-item-title"}
	website.ReleaseDate = parser.HTMLElement{Tag: "span", Class: "js-live-date-stopped"}

	newAlbums, err := parser.ScrapeWebsite(website)

	if err != nil {
		fmt.Println("Error Occurred: ", err)
		return
	}

	for i := 0; i < len(newAlbums); i++ {
		fmt.Printf("%d: %+v\n", i, newAlbums[i])
	}
}
