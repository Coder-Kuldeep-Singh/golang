package main

import (
	"os"
	"strings"

	"golang.org/x/net/html"
)

var testhtml = `
<table style="s1">
  <tr class="final"><td id="foo" style="s2">R1, C1</td><td>R1, C2</td></tr>
  <tr><td alt="hi">R2, C1</td><td style="s3">R2, C2</td></tr>
</table>`

func RemoveAttr(n *html.Node) {
	i := -1
	for index, attr := range n.Attr {
		if attr.Key == "id" || attr.Key == "style" || attr.Key == "class" {
			i = index
			break
		}
	}
	if i != -1 {
		n.Attr = append(n.Attr[:i], n.Attr[i+1:]...)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		RemoveAttr(c)
	}
}

func main() {
	doc, err := html.Parse(strings.NewReader(testhtml))
	if err != nil {
		panic(err)
	}

	RemoveAttr(doc)

	html.Render(os.Stdout, doc)
}
