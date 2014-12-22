// Copyright 2014 Hari haran. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

// Package extract provides simple methods for extracting
// specific items from a HTML page.
package extract

import (
	"net/http"
	"net/url"

	"code.google.com/p/cascadia"
	"golang.org/x/net/html"
)

// A selection contains the required elements for extraction.
type selection struct {
	Selector string // CSS Selector
	URL      string
}

// newSelection is a constructor for selection type.
// It takes the selector string and the url as inputs, in that order.
func newSelection(s, u string) *selection {
	return &selection{s, u}
}

// Links extracts all the referencing absolute URLs from a webpage.
func Links(u string) ([]string, error) {
	s := newSelection("a[href]", u)
	link, err := url.Parse(s.URL)
	if err != nil {
		return nil, err
	}
	r, err := http.Get(link.String())
	if err != nil {
		return nil, err
	}
	doc, err := html.Parse(r.Body)
	if err != nil {
		return nil, err
	}
	sel, err := cascadia.Compile(s.Selector)
	if err != nil {
		return nil, err
	}
	matches := sel.MatchAll(doc)
	var result []string
	for _, m := range matches {
		r, err := resolveURL(hrefString(m), link)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}
	return result, nil
}

// hrefString takes an *html.Node as input and
// returns the value of attribute href.
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

// imageString takes an *html.Node as input and
// returns the value of attribute src.
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

// attribute takes an html Token and the attribute key as inputs
// and returns the value of the attribute.
func attribute(t html.Token, a string) string {
	for _, x := range t.Attr {
		if x.Key == a {
			return x.Val
		}
	}
	return ""
}

// Images returns the absolute URLs of all the images in a HTML page.
// It takes the URL of the page as the input.
func Images(u string) ([]string, error) {
	s := newSelection("img[src]", u)
	link, err := url.Parse(s.URL)
	if err != nil {
		return nil, err
	}
	r, err := http.Get(link.String())
	if err != nil {
		return nil, err
	}
	doc, err := html.Parse(r.Body)
	if err != nil {
		return nil, err
	}
	sel, err := cascadia.Compile(s.Selector)
	if err != nil {
		return nil, err
	}
	matches := sel.MatchAll(doc)
	var result []string
	for _, m := range matches {
		r, err := resolveURL(imageString(m), link)
		if err != nil {
			return nil, err
		}
		result = append(result, r)
	}
	return result, nil
}

// resolveURL converts all input URLs into absolute URLs
func resolveURL(s string, link *url.URL) (string, error) {
	// x may or may not be an absolute URL.
	x, err := url.Parse(s)
	if err != nil {
		return "", err
	}
	// y is guaranteed to be an absolute URL
	y := link.ResolveReference(x)
	return y.String(), nil
}
