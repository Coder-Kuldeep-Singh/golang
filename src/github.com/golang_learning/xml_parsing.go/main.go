package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type SitemapIndex struct {
	Locations []Location `xml:"sitemap"`
}
type Location struct {
	Loc string `xml:"loc"`
}

func main() {
	xmlUrls()
}
func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
}
func xmlUrls() {
	domain := flag.String("d", "domain.com", "provide domain")
	protocol := flag.String("p", "https", "Protocol")
	flag.Parse()
	// url := lines(fetchUrl(*protocol + "://" + *domain + "/robots.txt"))
	// fmt.Printf("%T\n", url[1])
	// typecasting := string(url[1])
	xmlurl := fetchUrl(*protocol + "://" + *domain + "/https-sitemap_index.xml")
	var s SitemapIndex
	xml.Unmarshal(xmlurl, &s)
	// lastindex := len(s.Locations) - 1  //fetch last index of the array
	// fmt.Printf("%s\n", s.Locations[lastindex]) //
	fmt.Println(len(s.Locations))
	// for _, Location := range s.Locations {
	// 	// fmt.Printf("%L\n", Location)
	// }
}
func fetchUrl(url string) []uint8 {
	result, err := http.Get(url)
	Error(err)
	defer result.Body.Close()
	body, err := ioutil.ReadAll(result.Body)
	Error(err)
	return body
}

//parsing sitemap url using RegularExpression
// func lines(s []uint8) []string {
// 	sitemap := regexp.MustCompile(`sitemap: (.*)`)
// 	xmlpath := sitemap.FindAllStringSubmatch(string(s), -1)
// 	return xmlpath[0]
// }
func Error(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}
