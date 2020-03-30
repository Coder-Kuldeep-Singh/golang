package main

import (
	"fmt"
	"log"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	// response, err := http.Get("https://www.progville.com/go/goquery-jquery-html-golang/")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer response.Body.Close()
	// body, err := ioutil.ReadAll(response.Body)
	// doc, err := html.Parse(strings.NewReader(string(body)))
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// removeTags(doc)
	// buf := bytes.NewBuffer([]bytes{})
	// if err := html.Render(buf, doc); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(buf.String())
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a
href="/bar/baz">BarBaz</a></ul><span>TEXT I WANT</span>`
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node, bool)
	f = func(n *html.Node, printText bool) {
		if printText && n.Type == html.TextNode {
			fmt.Printf("%q\n", n.Data)
		}
		printText = printText || (n.Type == html.ElementNode && n.Data == "span")
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c, printText)
		}
	}
	f(doc, false)
}

// func removeTags(n *html.Node) {
// 	//tag
// 	if n.Type == html.ElementNode && n.Data == "script" {
// 		n.Parent.RemoveChild(n)
// 		return //tag is gone...
// 	}

// 	//treverse DOM
// 	for c := n.FirstChild; c != nil; c = c.NextSibling {
// 		removeTags(c)
// 	}

// }
