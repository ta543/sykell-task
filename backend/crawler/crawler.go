package crawler

import (
        "context"
        "net/http"
        "net/url"
        "strings"

        "github.com/PuerkitoBio/goquery"
)

// Result holds data about a crawled page
// Fields correspond to models.URL but without DB struct tags.
// LinkError contains information about a broken link on the page.
type LinkError struct {
        URL        string
        StatusCode int
}

type Result struct {
        Title            string
        HTMLVersion      string
        H1Count          int
        H2Count          int
        H3Count          int
        InternalLinks    int
        ExternalLinks    int
        BrokenLinksCount int
        BrokenLinks      []LinkError
        HasLoginForm     bool
}

func Crawl(ctx context.Context, target string) (*Result, error) {
        req, err := http.NewRequestWithContext(ctx, http.MethodGet, target, nil)
        if err != nil {
                return nil, err
        }
        resp, err := http.DefaultClient.Do(req)
        if err != nil {
                return nil, err
        }
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	result := &Result{}
	result.Title = strings.TrimSpace(doc.Find("title").Text())

	result.HTMLVersion = detectHTMLVersion(doc)

	result.H1Count = doc.Find("h1").Length()
	result.H2Count = doc.Find("h2").Length()
	result.H3Count = doc.Find("h3").Length()

	baseURL, err := url.Parse(target)
	if err != nil {
		baseURL = &url.URL{}
	}

       doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
               href, exists := s.Attr("href")
               if !exists {
                       return
               }
               u, err := url.Parse(href)
               if err != nil {
                       return
               }
               if !u.IsAbs() {
                       result.InternalLinks++
                       u = baseURL.ResolveReference(u)
               } else {
                       if u.Host == baseURL.Host {
                               result.InternalLinks++
                       } else {
                               result.ExternalLinks++
                       }
               }
               if code, broken := checkBroken(ctx, u.String()); broken {
                        result.BrokenLinksCount++
                        result.BrokenLinks = append(result.BrokenLinks, LinkError{URL: u.String(), StatusCode: code})
                }
       })

	result.HasLoginForm = doc.Find("input[type='password']").Length() > 0

	return result, nil
}

func checkBroken(ctx context.Context, link string) (int, bool) {
        req, err := http.NewRequestWithContext(ctx, http.MethodHead, link, nil)
        if err != nil {
                return 0, true
        }
        resp, err := http.DefaultClient.Do(req)
        if err != nil {
                return 0, true
        }
        defer resp.Body.Close()
        code := resp.StatusCode
        return code, code >= 400 && code <= 599
}

func detectHTMLVersion(doc *goquery.Document) string {
	nodes := doc.Nodes
	if len(nodes) == 0 {
		return "unknown"
	}
	node := nodes[0]
	if node.Type == 10 && node.Data == "html" { // doctype
		if strings.Contains(strings.ToLower(node.Data), "html") {
			return "HTML5"
		}
	}
	return "HTML4"
}
