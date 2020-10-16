package main

import (
	crawler "github.com/donavanbrooks/soundscraper/internal"
)

func main() {
	//testURL := "https://www.hotnewhiphop.com/mixtapes/"

	// var website parser.WebsiteTemplate
	// website.Name = "hotnewhiphop"
	// website.BaseURL = "https://www.hotnewhiphop.com/mixtapes/"
	// website.Base = parser.HTMLElement{Tag: "div", Class: "grid-item song"}
	// website.Artist = parser.HTMLElement{Tag: "em", Class: "default-artist"}
	// website.Title = parser.HTMLElement{Tag: "a", Class: "cover-title grid-item-title"}
	// website.ReleaseDate = parser.HTMLElement{Tag: "span", Class: "js-live-date-stopped"}

	// NumLinks int
	// Link     HTMLElement
	// URLTemp  string

	var website crawler.CrawlerTemplate
	website.BaseURL = "https://www.soundcloud.com/partyomo/partynextdoor-west-district/recommended/"
	website.Depth = 2
	website.DepthTemplates = []crawler.DepthTemplate{
		crawler.DepthTemplate{
			NumLinks: 5,
			Link:     crawler.HTMLElement{Tag: "a", Class: "url"},
			URLTemp:  "https://www.soundcloud.com%s/recommended/"},
		crawler.DepthTemplate{
			NumLinks: 5,
			Link:     crawler.HTMLElement{Tag: "a", Class: "url"},
			URLTemp:  "https://www.soundcloud.com%s/recommended/"},
		crawler.DepthTemplate{
			NumLinks: 5,
			Link:     crawler.HTMLElement{Tag: "a", Class: "url"},
			URLTemp:  "https://www.soundcloud.com%s/recommended/"},
	}

	crawler.CrawlWebsite(website)
	// newAlbums, err := parser.ScrapeWebsite(website)

	// if err != nil {
	// 	fmt.Println("Error Occurred: ", err)
	// 	return
	// }

	// for i := 0; i < len(newAlbums); i++ {
	// 	fmt.Printf("%d: %+v\n", i, newAlbums[i])
	// }
}
