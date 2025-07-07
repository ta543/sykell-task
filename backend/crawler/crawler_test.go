package crawler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/require"
)

func TestCrawlCollectsData(t *testing.T) {
	html := `<html><head><title>Test</title></head><body><h1>Hi</h1><a href="/ok">ok</a><a href="/bad">bad</a><a href="https://ext">ext</a><form><input type='password'></form></body></html>`
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "bad") {
			http.Error(w, "not found", http.StatusNotFound)
			return
		}
		w.Write([]byte(html))
	}))
	defer srv.Close()

	res, err := Crawl(context.Background(), srv.URL)
	require.NoError(t, err)
	require.Equal(t, "Test", res.Title)
	require.Equal(t, 1, res.H1Count)
	require.Equal(t, 2, res.InternalLinks)
	require.Equal(t, 1, res.ExternalLinks)
	require.Equal(t, 2, res.BrokenLinksCount)
	require.True(t, res.HasLoginForm)
}

func TestCheckBroken(t *testing.T) {
	srv := httptest.NewServer(http.NotFoundHandler())
	defer srv.Close()

	code, broken := checkBroken(context.Background(), srv.URL)
	require.True(t, broken)
	require.Equal(t, http.StatusNotFound, code)
}

func TestDetectHTMLVersionNoDoctype(t *testing.T) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader("<html></html>"))
	require.NoError(t, err)
	v := detectHTMLVersion(doc)
	require.Equal(t, "HTML4", v)
}

func TestCrawlGoogle(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := Crawl(ctx, "https://www.google.com")
	require.NoError(t, err)
	require.Contains(t, res.Title, "Google")
}

func TestCrawlWikipedia(t *testing.T) {
       ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
       defer cancel()
       res, err := Crawl(ctx, "https://www.wikipedia.org")
       require.NoError(t, err)
       require.Contains(t, res.Title, "Wikipedia")
}

func TestCrawlGitHub(t *testing.T) {
       ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
       defer cancel()
       res, err := Crawl(ctx, "https://github.com")
       require.NoError(t, err)
       require.Contains(t, res.Title, "GitHub")
}
