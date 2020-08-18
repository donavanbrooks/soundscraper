package parser

import (
	"io"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

type HTMLElement struct {
	Tag   string
	Class string
}

type WebsiteTemplate struct {
	Name        string
	URL         string
	Base        HTMLElement
	Artist      HTMLElement
	Title       HTMLElement
	ReleaseDate HTMLElement
	Genre       HTMLElement
	Cover       HTMLElement
}

type Date struct {
	Year  int
	Month int
	Day   int
}

type Album struct {
	Artist      string
	Title       string
	ReleaseDate string
	Genre       string
	Cover       string
}

func ScrapeWebsite(website WebsiteTemplate) ([]Album, error) {
	resp, err := http.Get(website.URL)

	if err != nil {
		return nil, err
	}

	responseBody := resp.Body
	defer responseBody.Close()

	newAlbums := parseHTML(responseBody, website)
	return newAlbums, nil
}

func parseHTML(htmlResp io.ReadCloser, website WebsiteTemplate) []Album {

	var newAlbums []Album
	tokenizer := html.NewTokenizer(htmlResp)

	for {

		token := tokenizer.Next()

		switch token {
		case html.ErrorToken:
			// End of the document, we're done. Return Albums
			return newAlbums

		case html.StartTagToken:
			t := tokenizer.Token()

			if t.Data == website.Base.Tag && len(t.Attr) > 0 {
				attr := t.Attr[0]
				if attr.Key == "class" && attr.Val == website.Base.Class {
					newAlbum := populateAlbum(tokenizer, website)
					newAlbums = append(newAlbums, newAlbum)
				}
			}
		}
	}
}

func populateAlbum(tokenizer *html.Tokenizer, website WebsiteTemplate) Album {
	var a Album

	// Keeps track of opening and closing tag matching
	baseTagCount := 1

	for {

		if baseTagCount == 0 {
			return a
		}

		token := tokenizer.Next()

		switch token {
		case html.EndTagToken:
			t := tokenizer.Token()

			// If matching base tag found, decrement count
			if t.Data == website.Base.Tag {
				baseTagCount--
			}

		case html.StartTagToken:
			t := tokenizer.Token()

			if t.Data == website.Base.Tag {
				baseTagCount++
			}

			if t.Data == website.Title.Tag && len(t.Attr) > 0 {
				attr := t.Attr[0]
				if attr.Key == "class" && attr.Val == website.Title.Class {
					a.Title = getText(tokenizer)
				}
			} else if t.Data == website.Artist.Tag && len(t.Attr) > 0 {
				attr := t.Attr[0]
				if attr.Key == "class" && attr.Val == website.Artist.Class {
					a.Artist = getText(tokenizer)
				}
			} else if t.Data == website.ReleaseDate.Tag && len(t.Attr) > 0 {
				attr := t.Attr[0]
				if attr.Key == "class" && attr.Val == website.ReleaseDate.Class {
					a.ReleaseDate = getText(tokenizer)
				}
			}
		}
	}
}

func getText(tokenizer *html.Tokenizer) string {

	tokenType := tokenizer.Next()

	if tokenType == html.TextToken {
		data := tokenizer.Token().Data
		return strings.TrimSpace(data)
	}

	return ""
}
