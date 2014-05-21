// extract provides functions for parsing HTML and extract specific items
package extract

import (
	"log"
	"net/http"

	"code.google.com/p/cascadia"
	"code.google.com/p/go.net/html"
)

type Selection struct {
	Selector string
	URL      string
}

// NewSelection is a constructor for Selection type.
// It takes the selector string and the url as inputs, in that order.
func NewSelection(s, u string) *Selection {
	return &Selection{s, u}
}

// ExtractLinks extracts all the referincing urls from a webpage.
func ExtractLinks(url string) []string {
	s := NewSelection("a[href]", url)
	link, err := url.Parse(s.URL)
	if err != nil {
		log.Fatal("Incorrect url")
		return nil
	}
	r, err := http.Get(link.String())
	if err != nil {
		log.Fatal(err)

	}
	doc, err := html.Parse(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	sel, err := cascadia.Compile(s.Selector)
	if err != nil {
		log.Fatal(err)
	}
	matches := sel.MatchAll(doc)
	var result []string
	for _, m := range matches {
		result = append(result, nodeString(m))
	}
	return result
}

func nodeString(n *html.Node) string {
	switch n.Type {
	case html.TextNode:
		return n.Data
	case html.ElementNode:
		return attribute(
			html.Token{
				Type: html.StartTagToken,
				Data: n.Data,
				Attr: n.Attr,
			}, "href")
	}
	return ""
}

// attribute takes a html Token and the attribute key as inputs
// and returns the value of the attribute.
func attribute(t html.Token, a string) string {
	for _, x := range t.Attr {
		if x.Key == a {
			return x.Val
		}
	}
	return nil
}
