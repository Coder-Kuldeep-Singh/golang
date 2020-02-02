package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/net/html"
)

var (
	config = &tls.Config{
		InsecureSkipVerify: true,
	}
	transport = &http.Transport{
		TLSClientConfig: config,
	}
	URLclient = &http.Client{
		Transport: transport,
	}
	queue   = make(chan string)
	Visited = make(map[string]bool)
)

func main() {
	URL := os.Args[1:]
	if len(URL) == 0 {
		fmt.Println("Missing URL")
		os.Exit(1)
	}
	baseURL := URL[0]
	go func() {
		queue <- baseURL
	}()
	for href := range queue {
		if !Visited[href] && SameDomain(href, baseURL) {
			CrawlingUrls(href)
		}
	}
	// body, err := ioutil.ReadAll(response.Body)
	// CheckErr(err)
	// fmt.Println(string(body))
}

func CrawlingUrls(href string) {
	Visited[href] = true
	fmt.Printf("Crawling urls ---> %v \n", href)
	response, err := URLclient.Get(href)
	CheckErr(err)
	// fmt.Println(response)
	defer response.Body.Close()
	links := getLinks(response.Body)
	for _, link := range links {
		// fmt.Printf("index %v -- link %+v \n", i, link)
		// CrawlingUrls(FixedUrl(link, href))
		absoluteURL := FixedUrl(link, href)
		go func() {
			queue <- absoluteURL
		}()
		// CrawlingUrls(link)
	}
}

//checkk if the links are related to same domain
func SameDomain(href, baseURL string) bool {
	uri, err := url.Parse(href)
	if err != nil {
		return false
	}

	parentUri, err := url.Parse(baseURL)
	if err != nil {
		return false
	}
	if uri.Host != parentUri.Host {
		return false
	}
	return true
}

func FixedUrl(href, baseURL string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return ""
	}

	base, err := url.Parse(baseURL)
	if err != nil {
		return ""
	}
	FixedURLS := base.ResolveReference(uri)
	return FixedURLS.String()
}

//Collect all links from response body and return it as an array of strings
func getLinks(body io.Reader) []string {
	var links []string
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			//todo: links list shoudn't contain duplicates
			return links
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}

				}
			}

		}
	}
}

func CheckErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
