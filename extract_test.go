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
	r := Links(ts.URL)
	if len(r) != 2 {
		t.Errorf("length of results slice should be 2 but got %d", len(r))
	}
}

func TestExtractImages(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w,
			`<html><body><ul><img src="/smiley.png"><li><a href="http://example.com/1"></a><li><a href="/2"></a></ul></body></html>`,
		)
	}))
	defer ts.Close()
	s := Images(ts.URL)
	if len(s) != 1 {
		t.Errorf("length of results slice should be 1 but got %d", len(s))
	}
	if s[0] == "smiley.png" {
		t.Errorf("URLs not getting resolved into absolute URLs")
	}

}
