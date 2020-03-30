package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func main() {

	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		tags := [15]string{"div", "header", "footer", "p", "span", "a", "ul", "li", "ol", "h1", "h2", "h3",
			"h4", "h5", "h6"}
		// fmt.Println(tags)
		for _, tag := range tags {
			// fmt.Println(tag)
			if n.Type == html.ElementNode && n.Data == string(tag) {
				fmt.Println(string(tag))
				// for _, a := range n.Attr {
				// 	if a.Key ==  {
				// 		fmt.Println(a.Val)
				// 		return
				// 		// break
				// 	}
				// }
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				f(c)
				// return
			}
		}

	}
	f(doc)

}
