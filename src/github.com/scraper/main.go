package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {

	url := "https://chaufferus.us/open-positions/Illinois/Chicago/60629/3127090/3"

	if resp, err := http.Get(url); err == nil {
		defer resp.Body.Close()

		log.Println("Load page complete")

		if resp != nil {
			log.Println("Page response is NOT nil")
			// --------------
			data, _ := ioutil.ReadAll(resp.Body)
			resp.Body.Close()

			hdata := strings.Replace(string(data), "<noscript>", "", -1)
			hdata = strings.Replace(hdata, "</noscript>", "", -1)
			// --------------

			if document, err := html.Parse(strings.NewReader(hdata)); err == nil {
				var parser func(*html.Node)
				parser = func(n *html.Node) {
					if n.Type == html.ElementNode && n.Data == "img" {

						var imgSrcUrl, imgDataOriginal string

						for _, element := range n.Attr {
							if element.Key == "src" {
								imgSrcUrl = element.Val
							}
							if element.Key == "data-original" {
								imgDataOriginal = element.Val
							}
						}

						log.Println(imgSrcUrl, imgDataOriginal)
					}

					for c := n.FirstChild; c != nil; c = c.NextSibling {
						parser(c)
					}

				}
				parser(document)
			} else {
				log.Panicln("Parse html error", err)
			}

		} else {
			log.Println("Page response IS nil")
		}
	}

}

