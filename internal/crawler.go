package internal

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

type CrawlerTemplate struct {
	BaseURL        string
	Depth          int
	DepthTemplates []DepthTemplate
	LinkHTML       string
}

type DepthTemplate struct {
	NumLinks int
	Link     HTMLElement
	URLTemp  string
}

type Song struct {
	Base        string
	Title       string
	ReleaseDate string
	Genre       string
	Cover       string
	Link        string
}

func CrawlWebsite(crawler CrawlerTemplate) error {
	urls, err := crawl(crawler.BaseURL, crawler.Depth, crawler)

	if err != nil {
		return err
	}

	for _, u := range urls {
		fmt.Printf("- %s\n", u)
	}

	return nil
}

func getHTML(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	responseBody := resp.Body
	defer responseBody.Close()

	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(responseBody)
	if err != nil {
		panic(err)
	}

	return html, nil
}

func crawl(url string, depth int, template CrawlerTemplate) ([]string, error) {

	if depth <= 0 {
		return nil, nil
	}

	var urls []string
	var temp = template.DepthTemplates[template.Depth-depth]
	fetchedUrls, err := fetchUrls(url, temp)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// fmt.Printf("FOUND, Depth %d:", depth)
	// for _, u := range fetchedUrls {
	// 	fmt.Printf("- %s\n", u)
	// }

	for _, u := range fetchedUrls {
		var formattedURL = u
		if temp.URLTemp != "" {
			formattedURL = fmt.Sprintf(temp.URLTemp, u)
		}

		urls = append(urls, u)
		dope, _ := crawl(formattedURL, depth-1, template)
		urls = append(urls, dope...)
	}

	return urls, nil
}

func fetchUrls(baseURL string, depth DepthTemplate) ([]string, error) {

	data, err := getHTML(baseURL)
	// show the HTML code as a string %s
	//fmt.Printf("%s\n", data)

	if err != nil {
		return nil, err
	}

	var urls []string
	linksFound := depth.NumLinks

	if document, err := html.ParseWithOptions(bytes.NewReader([]byte(data)), html.ParseOptionEnableScripting(false)); err == nil {
		//fmt.Printf("%s\n", document)
		var parser func(*html.Node)

		parser = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "a" {
				a := n.Attr[0]
				if a.Key == "itemprop" && a.Val == depth.Link.Class && linksFound > 0 {
					url := getHref(n)
					urls = append(urls, url)
					linksFound--
				}
			}

			for c := n.FirstChild; c != nil; c = c.NextSibling {
				parser(c)
			}
		}

		parser(document)

	} else {
		log.Panicln("Parse html error", err)
	}

	return urls, err
}

func getHref(n *html.Node) string {

	// Iterate over all of the Token's attributes until we find an "href"
	for _, a := range n.Attr {
		if a.Key == "href" {
			href := a.Val
			//fmt.Printf("Got href: %s\n", href)
			return href
		}
	}

	return ""
}

// NEED TO FLUSH OUT RULES TO VALIDATE THE CRAWLER TEMPLATE
func validateCrawlerTemplate(template CrawlerTemplate) (bool, string) {
	return false, ""
}
