package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gookit/color"
	"golang.org/x/net/html"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	fmt.Println(" ---------------------> Broken")
	fmt.Println(" <--------------------- Not Broken")
	lines, err := readLines("urls.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for _, line := range lines {
		// filter := regexp.MustCompile(`domain_name=(.*.*) site_name`)
		// stored := filter.FindAllStringSubmatch(line, -1)
		stored := strings.Split(line, "\n")
		for _, element := range stored {
			// fmt.Fprintln(w, element[1])
			url := element
			// url := "https://chaufferus.us/open-positions/Illinois/Chicago/60629/3127090/3"
			if resp, err := http.Get(url); err == nil {
				defer resp.Body.Close()

				// log.Println("Load page complete")

				if resp != nil {
					// log.Println("Page response is NOT nil")
					// --------------
					data, _ := ioutil.ReadAll(resp.Body)
					resp.Body.Close()

					hdata := strings.Replace(string(data), "<noscript>", "", -1)
					hdata = strings.Replace(hdata, "</noscript>", "", -1)
					// --------------

					if document, err := html.Parse(strings.NewReader(hdata)); err == nil {
						yellow := color.FgYellow.Render
						var parser func(*html.Node)
						parser = func(n *html.Node) {
							if n.Type == html.ElementNode && n.Data == "img" {

								var imgSrcUrl string

								for _, element := range n.Attr {
									if element.Key == "src" {
										imgSrcUrl = element.Val
									}
									// if element.Key == "data-original" {
									// 	imgDataOriginal = element.Val
									// }
								}

								// fmt.Println(url + " ----------------> " + imgSrcUrl)

								re := strings.Replace(imgSrcUrl, "/open-positions/", "", -1)
								// log.Println("imageUrl " + re)

								resp, err := http.Get(url + re)
								red := color.FgRed.Render
								if err != nil {
									fmt.Println(err)
								}
								if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
									fmt.Println(url + " <-------------- " + re)
									fmt.Print()
								} else {
									fmt.Println(url + " --------------> " + red(re))
								}
								defer resp.Body.Close()
								// resp.Body.Close()
								// bodyElement, err := ioutil.ReadAll(resp)
							}

							for c := n.FirstChild; c != nil; c = c.NextSibling {
								parser(c)
							}

						}
						fmt.Println(yellow("***************************************************************************************************"))
						parser(document)
					} else {
						log.Panicln("Parse html error", err)
					}

				} else {
					log.Println("Page response IS nil")
				}
			}

		}
	}
}
