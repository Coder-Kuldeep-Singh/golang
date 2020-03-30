package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func main() {

	response, err := http.Get("https://gist.github.com/Xeoncross/7dd11225e6f1484683301f0095b1dfbc")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	HTMLString := string(body)

	var c chan Node
	var title string
	var a []string
	wantedTokens := []string{
		"a", "title",
	}

	c = GetTokensFromHTMLString(HTMLString, wantedTokens)

	for node := range c {
		// fmt.Println(node.Type, node)

		if node.Type == "title" {
			tt := node.Doc.Next()

			if tt == html.TextToken {
				next := node.Doc.Token()
				title = strings.TrimSpace(next.Data)
			}
		}

		if node.Type == "a" {
			tt := node.Doc.Next()

			if tt == html.TextToken {
				next := node.Doc.Token()
				// a =
				a = append(a, strings.TrimSpace(next.Data))

			}
		}

	}

	fmt.Println("title", title)
	fmt.Println("a", a)

}

// Node foobar
type Node struct {
	Type  string
	Token html.Token
	Doc   *html.Tokenizer
}

// GetTokensFromHTMLString foobar
func GetTokensFromHTMLString(HTMLString string, wantedTokens []string) (c chan Node) {

	c = make(chan Node)

	go func() {

		defer close(c)

		// https://play.golang.org/p/0MRSefJ_-E
		r := strings.NewReader(HTMLString)
		z := html.NewTokenizer(r)

		// defer func() {
		// 	close(c)
		// }

		for {
			tt := z.Next()

			switch {
			case tt == html.ErrorToken:
				// End of the document, we're done
				return
			case tt == html.StartTagToken:
				token := z.Token()

				for _, name := range wantedTokens {
					if token.Data == name {
						c <- Node{token.Data, token, z}
					}
					continue
				}
			}
		}
	}()

	return c
}
