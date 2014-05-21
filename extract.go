// extract provides functions for parsing HTML and extract specific items
package extract

import (
	"log"
	"net/http"
	"net/url"

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

// Links extracts all the referincing urls from a webpage.
func Links(u string) []string {
	s := NewSelection("a[href]", u)
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
		// x may or may not be a absolute URL.
		x, err := url.Parse(hrefString(m))
		if err != nil {
			log.Fatal(err)
		}
		//y is guaranteed to be a absolute URL
		y := link.ResolveReference(x)
		result = append(result, y.String())
	}
	return result
}

// hrefString takes a *html.Node as input and
// returns the value of attribute href
func hrefString(n *html.Node) string {
	switch n.Type {
	case html.TextNode:
		return ""
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

func imageString(n *html.Node) string {
	switch n.Type {
	case html.TextNode:
		return ""
	case html.ElementNode:
		return attribute(
			html.Token{
				Type: html.StartTagToken,
				Data: n.Data,
				Attr: n.Attr,
			}, "src")
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
	return ""
}

func Images(u string) []string {
	s := NewSelection("img[src]", u)
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
		// x may or may not be a absolute URL.
		x, err := url.Parse(imageString(m))
		if err != nil {
			log.Fatal(err)
		}
		//y is guaranteed to be a absolute URL
		y := link.ResolveReference(x)
		result = append(result, y.String())
	}
	return result
}
