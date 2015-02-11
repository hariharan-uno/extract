// Copyright 2014 Hari haran. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package extract

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExtractLinks(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,
			`<html><body><ul><li><a href="http://example.com/1"></a><li><a href="/2"></a></ul></body></html>`,
		)
	}))
	defer ts.Close()
	r, err := Links(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	if len(r) != 2 {
		t.Errorf("length of results slice should be 2 but got %d", len(r))
	}
	if r[1] == "/2" {
		t.Errorf("URLs not getting resolved into absolute URLs [Links]")
	}
}

func TestExtractImages(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,
			`<html><body><ul><img src="/smiley.png"><li><a href="http://example.com/1"></a><li><a href="/2"></a></ul></body></html>`,
		)
	}))
	defer ts.Close()
	s, err := Images(ts.URL)
	if err != nil {
		t.Fatal(err)
	}
	if len(s) != 1 {
		t.Errorf("length of results slice should be 1 but got %d", len(s))
	}
	if s[0] == "/smiley.png" {
		t.Errorf("URLs not getting resolved into absolute URLs [Images]")
	}

}

func BenchmarkExtractLinks(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,
			`<html><body><ul><li><a href="http://example.com/1"></a><li><a href="/2"></a></ul></body></html>`,
		)
	}))
	defer ts.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Links(ts.URL)
	}
}

func BenchmarkExtractImages(b *testing.B) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,
			`<html><body><ul><img src="/smiley.png"><li><a href="http://example.com/1"></a><li><a href="/2"></a></ul></body></html>`,
		)
	}))
	defer ts.Close()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Images(ts.URL)
	}
}
