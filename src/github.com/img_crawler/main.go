package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func Domains() {
	//fetch data from given url
	// domain := os.Args
	response, err := http.Get("https://chaufferjob.us/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	// Read data from url
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	storeDomains := string(body)
	re := regexp.MustCompile(`sitemap: (.*)`)
	// re := regexp.MustCompile(`http://(.*)`)
	submatchall := re.FindAllStringSubmatch(storeDomains, -1)
	for _, element := range submatchall {
		// fmt.Println(element[1])
		response, err := http.Get(element[1])
		if err != nil {
			log.Fatal(err)
		}
		defer response.Body.Close()
		// Read data from url
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		storeDomains := string(body)
		re := regexp.MustCompile(`<loc>(.*)</loc>`)
		// re := regexp.MustCompile(`http://(.*)`)
		submatchall := re.FindAllStringSubmatch(storeDomains, -1)
		// fmt.Println(submatchall)
		for _, elements := range submatchall {
			// fmt.Println(element[1])
			response, err := http.Get(elements[1])
			if err != nil {
				log.Fatal(err)
			}
			defer response.Body.Close()
			// Read data from url
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				log.Fatal(err)
			}
			storeDomains := string(body)
			// fmt.Println(storeDomains)
			re := regexp.MustCompile(`<loc>(.*)</loc>`)
			// re := regexp.MustCompile(`http://(.*)`)
			submatchall := re.FindAllStringSubmatch(storeDomains, -1)
			for _, element := range submatchall {
				// fmt.Println(element[1])
				response, err := http.Get(element[1])
				if err != nil {
					log.Fatal(err)
				}

				defer response.Body.Close()
				// Read data from url
				body, err := ioutil.ReadAll(response.Body)
				if err != nil {
					log.Fatal(err)
				}
				storeDomains := string(body)
				re := regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)
				// re := regexp.MustCompile(`http://(.*)`)
				submatchall := re.FindAllStringSubmatch(storeDomains, -1)
				// fmt.Println(storeDomains)
				// imgs := re.FindAllStringSubmatch(h, -1)
				out := make([]string, len(submatchall))
				for i := range out {
					out[i] = submatchall[i][1]
					fmt.Println(out[i])
				}
			}

		}

	}
}

func main() {
	Domains()
	// http.HandleFunc("/v1/domains", Domains)
	// fmt.Println("Development server started localhost:8000/v1/domains")
	// log.Fatal(http.ListenAndServe(":8000", nil))
}
