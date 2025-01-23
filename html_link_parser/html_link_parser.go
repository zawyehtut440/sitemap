package html_link_parser

import (
	"os"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func GetLinksFromHtmlFile(fileName string) []Link {
	doc := parseHtmlFile(fileName)
	links := getLinks(doc)
	return links
}

func GetLinksFromHtmlString(htmlContent string) []Link {
	doc := parseHtmlString(htmlContent)
	links := getLinks(doc)
	return links
}

func parseHtmlFile(fileName string) *html.Node {
	// read the whole content of fileName
	htmlContent, err := os.ReadFile(fileName)
	check(err)
	// parse htmlContent
	return parseHtmlString(string(htmlContent))
}

func parseHtmlString(htmlContent string) *html.Node {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	check(err)
	return doc
}

type Link struct {
	Href string
	Text string
}

// search all first meet anchor tags
// for every first meet anchor, find the text and href, then return the Link struct
func getLinks(root *html.Node) []Link {
	links := make([]Link, 0)

	anchorTags := searchAllFirstMeetATag(root)

	for _, aTags := range anchorTags {
		links = append(links, createLink(aTags))
	}

	return links
}

func searchAllFirstMeetATag(root *html.Node) []*html.Node {
	result := make([]*html.Node, 0)

	for n := range root.ChildNodes() { // for each child of root node
		if n.Type == html.ElementNode && n.DataAtom == atom.A { // find a anchor tag
			result = append(result, n)
		} else { // not an anchor tag
			result = append(result, searchAllFirstMeetATag(n)...) // find anchor tag from n
		}
	}

	return result
}

func createLink(anchorTag *html.Node) Link {
	href, text := getHref(anchorTag), ""
	for n := range anchorTag.Descendants() {
		// if n is the element node, then n.Data is the tag name
		// if n is the text node, then n.Data is literally text
		if n.Type == html.TextNode && !isBlankText(n.Data) {
			text += strings.TrimSpace(n.Data) + " "
		}
	}
	return Link{
		Href: href,
		Text: strings.TrimRight(text, " "),
	}
}

func isBlankText(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func getHref(anchorTag *html.Node) string {
	for _, attr := range anchorTag.Attr { // find anchorTag's attr
		if attr.Key == "href" { // if found a key is href
			return attr.Val // return the href's value
		}
	}
	return "" // no href attribute
}
