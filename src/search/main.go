package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type GoogleResult struct {
	ResultRank  int
	ResultURL   string
	ResultTitle string
	ResultDesc  string
}

var googleDomains = map[string]string{
	"com": "https://www.google.com/search?q=",
	"uk":  "https://www.google.co.uk/search?q=",
	"ru":  "https://www.google.ru/search?q=",
	"fr":  "https://www.google.fr/search?q=",
}

func buildGoogleUrl(searchTerm string, countryCode string, languageCode string) string {
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	if googleBase, found := googleDomains[countryCode]; found {
		return fmt.Sprintf("%s%s&num=100&hl=%s", googleBase, searchTerm, languageCode)
	} else {
		return fmt.Sprintf("%s%s&num=100&hl=%s", googleDomains["com"], searchTerm, languageCode)
	}
}

func googleRequest(searchURL string) (*http.Response, error) {

	baseClient := &http.Client{}

	req, _ := http.NewRequest("GET", searchURL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")

	res, err := baseClient.Do(req)

	if err != nil {
		return nil, err
	} else {
		return res, nil
	}
}

func googleResultParser(response *http.Response) ([]GoogleResult, error) {
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return nil, err
	}

	results := []GoogleResult{}
	sel := doc.Find("div.g")
	rank := 1
	for i := range sel.Nodes {
		item := sel.Eq(i)
		linkTag := item.Find("a")
		link, _ := linkTag.Attr("href")
		titleTag := item.Find("h3.r")
		descTag := item.Find("span.st")
		desc := descTag.Text()
		title := titleTag.Text()
		link = strings.Trim(link, " ")
		if link != "" && link != "#" {
			result := GoogleResult{
				rank,
				link,
				title,
				desc,
			}
			results = append(results, result)
			rank += 1
		}
	}
	return results, err
}

func GoogleScrape(searchTerm string, countryCode string, languageCode string) ([]GoogleResult, error) {
	googleUrl := buildGoogleUrl(searchTerm, countryCode, languageCode)
	res, err := googleRequest(googleUrl)
	if err != nil {
		return nil, err
	}
	scrapes, err := googleResultParser(res)
	if err != nil {
		return nil, err
	} else {
		return scrapes, nil
	}
}

func PageResponse(response *http.Response) (string, error) {
	body, err := ReadAllResponse(response.Body)
	if body == nil {
		return "", err
	}
	return string(body), nil
}

func VisitToEachDomain(domain string) (string, error) {
	resp, err := googleRequest(domain)
	if err != nil {
		return "", err
	}
	body, err := PageResponse(resp)
	if err != nil {
		return "", err
	}
	// fmt.Println(resp)
	// return string(body), nil
	filename, err := BuildFileName(domain)
	if err != nil {
		return "", err
	}
	generatefile, err := GenerateFile(filename)
	if err != nil {
		return "", err
	}
	build, err := Put(generatefile, string(body))
	if err != nil {
		return "", err
	}
	link, err := TookLinks(domain)
	if err != nil {
		return "", err
	}
	dom, err := url.Parse(domain)
	if err != nil {
		return "", err
	}
	fmt.Println(dom.Host)
	for _, l := range link {
		uri, err := url.Parse(l)
		if err != nil {
			return "", err
		}
		if uri.Host != "" {
			fmt.Println(l)
		}

		if uri.Host == "" {
			NewLink := dom.Scheme + "://" + dom.Host + l
			fmt.Println(NewLink)
		}

	}
	return string(build), nil

}

func TookLinks(domain string) ([]string, error) {
	doc, err := goquery.NewDocument(domain)
	if err != nil {
		return nil, err
	}
	var links []string
	doc.Find("a").Each(func(i int, items *goquery.Selection) {
		href, exists := items.Attr("href")
		if !exists {
			// fmt.Println(href)
			fmt.Println("No Tag Exists")
		}
		href = strings.Trim(href, " ")
		if href != "" && href != "#" {
			links = append(links, href)
		}
		// fmt.Println(links)
	})
	return links, nil
}

func BuildFileName(url string) (string, error) {
	removeSpecialCharacter := strings.Replace(url, "/", "-", -1)
	if removeSpecialCharacter == "" {
		return "", nil
	}
	return removeSpecialCharacter, nil
}

func GenerateFile(url string) (*os.File, error) {
	filename, err := BuildFileName(url)
	if err != nil {
		return nil, err
	}
	create, err := os.Create("./output/" + filename + ".html")
	if err != nil {
		return nil, err
	}
	return create, nil
}

func Put(file *os.File, data string) (string, error) {
	// size, err := io.Copy(file, data)
	size, err := file.WriteString(data)
	if err != nil {
		return "", err
	}
	return string(size), nil
}

// readAll reads from r until an error or EOF and returns the data it read
// from the internal buffer allocated with a specified capacity.
func readAll(r io.Reader, capacity int64) (b []byte, err error) {
	var buf bytes.Buffer
	// If the buffer overflows, we will get bytes.ErrTooLarge.
	// Return that as an error. Any other panic remains.
	defer func() {
		e := recover()
		if e == nil {
			return
		}
		if panicErr, ok := e.(error); ok && panicErr == bytes.ErrTooLarge {
			err = panicErr
		} else {
			panic(e)
		}
	}()
	if int64(int(capacity)) == capacity {
		buf.Grow(int(capacity))
	}
	_, err = buf.ReadFrom(r)
	return buf.Bytes(), err
}

// ReadAll reads from r until an error or EOF and returns the data it read.
// A successful call returns err == nil, not err == EOF. Because ReadAll is
// defined to read from src until EOF, it does not treat an EOF from Read
// as an error to be reported.
func ReadAllResponse(r io.Reader) ([]byte, error) {
	return readAll(r, bytes.MinRead)
}

func main() {
	query := flag.String("q", "", "Please provide the query you want to search")
	countryCode := flag.String("s", "com", "Please provide the country code")
	languageCode := flag.String("l", "en", "Please provide the language code")
	Domain := flag.String("D", "", "Provide the direct domain")
	flag.Parse()
	fmt.Println("Did you want to Run the Software")
	fmt.Print("Yes or No: ") //Print function is used to display output in same line
	var first string
	fmt.Scanln(&first)
	if first == "yes" || first == "Yes" || first == "YES" {
		fmt.Println("Crawling Started")
		if *query == "" {
			resp, err := VisitToEachDomain(*Domain)
			if err != nil {
				log.Println("Having Error to visiting url: ", err)
			}
			fmt.Println(resp)
		}
		if *Domain == "" {
			res, _ := GoogleScrape(*query, *countryCode, *languageCode)
			fmt.Println("-----------------------")
			// fmt.Println(*query)
			for _, item := range res {
				fmt.Println()
				// fmt.Println(item.ResultRank)
				// fmt.Println(item.ResultTitle)
				// fmt.Println(item.ResultURL)
				resp, err := VisitToEachDomain(item.ResultURL)
				if err != nil {
					log.Println("Having Error to visiting url: ", err)
				}
				fmt.Println(resp)
				// fmt.Println(item.ResultDesc)
				// fmt.Println()
			}

		}
		fmt.Println("Crawling Finished...")

	} else {
		os.Exit(1)
	}
	time.Sleep(time.Second * 30)
}
