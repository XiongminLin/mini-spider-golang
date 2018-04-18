/* webpage_parse.go - parse urls from web page	*/
/*
modification history
--------------------
2017/07/18, by Xiongmin LIN, create
*/
/*
DESCRIPTION
*/
package web_package

import (
	"bytes"
	"fmt"
	"net/url"
)

import (
	"golang.org/x/net/html"
)

type HtmlLinks struct {
	links []string
}

// create new HtmlLinks
func NewHtmlLinks() *HtmlLinks {
	hl := new(HtmlLinks)
	hl.links = make([]string, 0)
	return hl
}

/*
get all href in given html node

Params:
	- n: html node
	- refUrl: reference url
*/
func (hl *HtmlLinks) getLinks(n *html.Node, refUrl *url.URL) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				linkUrl, err := refUrl.Parse(a.Val)
				if err == nil {
					hl.links = append(hl.links, linkUrl.String())
				}
				break
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		hl.getLinks(c, refUrl)
	}
}

/*
get url links in given html page

Params:
	- data: data for html page
	- urlStr: url string of this html page

Returns:
	- links: parsed links
	- error: any failure
*/
func ParseWebPage(data []byte, urlStr string) ([]string, error) {
	// parse html
	doc, err := html.Parse(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("html.Parse():%s", err.Error())
	}

	// parse url
	refUrl, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, fmt.Errorf("url.ParseRequestURI(%s):%s", urlStr, err.Error())
	}
	
	// get all links
	hl := NewHtmlLinks()
	hl.getLinks(doc, refUrl)

	return hl.links, nil
}
