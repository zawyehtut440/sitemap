package sitemap_builder

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"

	"sitemap/html_link_parser"
	"sitemap/sitemap_builder/data_structures"
)

type Sitemap struct {
	XMLName  xml.Name `xml:"urlset"`
	Protocol string   `xml:"xmlns,attr"`
	Urls     []Url    `xml:"url>url"`
}

type Url struct {
	Loc string `xml:"loc"`
}

func CreateSitemap(rootUrl string) string {
	u := findAllDomainLinks(rootUrl)
	fmt.Println(u)
	return `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9">
	<url>
		<loc>http://www.example.com/</loc>
	</url>
	<url>
		<loc>http://www.example.com/dogs</loc>
	</url>
</urlset>`
}

func findAllDomainLinks(rootUrl string) []Url {
	validUrls := make([]Url, 0)

	// using queue to visit all the websites, for each website url is unique and in the same domain
	urlQueue := data_structures.NewQueue[string]()
	urlQueue.Enqueue(rootUrl)                   // find root url first
	visited := data_structures.NewSet[string]() // visited: url that have already visited

	for !urlQueue.IsEmpty() { // while urlQueue is not empty
		url, _ := urlQueue.Dequeue()                               // get url from urlQueue
		htmlPage := getHtmlFromUrl(url)                            // get html page by visiting url page
		links := html_link_parser.GetLinksFromHtmlString(htmlPage) // get all links from  html page
		for _, link := range links {                               // for each link in links
			href := link.Href // get link's href
			if href == "" {   // if no href
				continue // check next link
			}
			targetUrl := href        // set targetUrl to href
			if targetUrl[0] == '/' { // if tartgetUrl is like "/about"
				targetUrl = rootUrl + targetUrl // prefix the domain, so make this targetUrl like "rootUrl/about"
			}

			if !visited.Contains(targetUrl) && strings.HasPrefix(targetUrl, rootUrl) { // if link haven't visited yet and is same domain as rootUrl
				fmt.Println(targetUrl)
				urlQueue.Enqueue(targetUrl)                        // add to urlQueue
				visited.Add(targetUrl)                             // targetUrl is visited
				validUrls = append(validUrls, Url{Loc: targetUrl}) // add targetUrl
			}
		}
	}

	return validUrls
}

// url: 要找的網頁
// return: html格式的內容
func getHtmlFromUrl(url string) string {
	resp, err := http.Get(url)
	check(err)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	check(err)

	return string(body)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
