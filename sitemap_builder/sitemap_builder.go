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
	Urls     []Url    `xml:"url"`
}

type Url struct {
	Loc string `xml:"loc"`
}

func CreateSitemap(rootUrl string) string {
	sitemap := &Sitemap{Protocol: "http://www.sitemaps.org/schemas/sitemap/0.9"}
	if rootUrl[len(rootUrl)-1] == '/' {
		rootUrl = rootUrl[:len(rootUrl)-1]
	}
	sitemap.Urls = findAllDomainLinks(rootUrl)
	out, _ := xml.MarshalIndent(sitemap, "", "    ")
	return xml.Header + string(out)
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
			href := link.Href      // get link's href
			if canSkipHref(href) { // if href can be skipped
				continue // check next link
			}
			targetUrl := formatUrl(href, rootUrl) // get full url

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

func canSkipHref(href string) bool {
	// if href is empty or href is html id or href is prefix with mailto:
	return href == "" || strings.Contains(href, "#") || strings.HasPrefix(href, "mailto:")
}

// get full url of href
func formatUrl(href, rootUrl string) string {
	// if href have full url with porefix protocol like https or http
	if strings.HasPrefix(href, "https://") || strings.HasPrefix(href, "http://") {
		return href // href is full url
	}

	if href[0] == '.' {
		href = href[1:]
	}

	if href[0] == '/' { // if href is like "/about"
		href = rootUrl + href // prefix the domain, so make this href like "rootUrl/about"
	} else {
		href = rootUrl + "/" + href
	}

	return href
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
