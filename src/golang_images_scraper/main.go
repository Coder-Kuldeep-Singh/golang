var (
	Images = make(map[string]string)
)

//Checking images into embed css
func ChecksEmbedCss(pageContent string, protocol, domain string, errorCount int) {
	Openingtag := "<style>"
	ClosingTag := "</style>"
	Length := len(Openingtag)
	StartIndex := strings.Index(pageContent, Openingtag)
	if StartIndex == -1 {
		log.Println("No element found")
		return
	}
	//Find the index of the closing tag
	EndIndex := strings.Index(pageContent, ClosingTag)
	if EndIndex == -1 {
		log.Println("No closing tag found.")
		return
	}

	pageImages := []byte(pageContent[StartIndex+Length : EndIndex])
	re := regexp.MustCompile(`background-image: url(.*);`)
	bgimages := re.FindAllStringSubmatch(string(pageImages), -1)
	if bgimages == nil {
		// log.Println("No matches.")
		return
	} else {
		for _, bgimage := range bgimages {
			Images["img"] = bgimage[1]
			CheckStatusOfImages(protocol, domain, Images["img"], errorCount)
			return
		}
	}
}

var (
	CssLinks = make(map[string]string)
)

func ChecksCssLinks(protocol, domain, pageContent string) {
	// doc, err := goquery.NewDocument(url)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// // response, err := url.Parse(Url)
	// // if err != nil {
	// // 	log.Println("Failed to Parse Url", err)
	// // 	return
	// // }
	// // var cssFile string
	// doc.Find("link").Each(func(index int, item *goquery.Selection) {
	// 	linkTag := item
	// 	link1, _ := linkTag.Attr("href")
	// 	// link, _ := linkTag.Attr("rel")
	// 	// linkText := linkTag.Text()
	// 	fmt.Printf("%s\n", link1)
	// 	// cssFile += link1 + "\n"
	// 	// fmt.Println(cssFile)
	// 	// cssFile += link1
	// 	// trim := strings.Trim(cssFile, ",")
	// 	// fmt.Println(trim)

	// })
	// // fmt.Println(cssFile)
	// // CssLinks["css"] = cssFile
	// // fmt.Println(CssLinks["css"])
	re := regexp.MustCompile(`<link rel='stylesheet' href='(.*)'`)
	Links := re.FindAllStringSubmatch(pageContent, -1)
	if Links == nil {
		// fmt.Println("No matches.")
	} else {
		for _, link := range Links {
			// fmt.Println(link[1])
			CheckCssJavascriptLinks(protocol, domain, link[1])
		}
	}

}

//CheckBodyOfCssFiles functions checks the background images in css files
func CheckBodyofCssFiles(response string) {
	re := regexp.MustCompile(`background-image: url(.*);`)
	bgimages := re.FindAllStringSubmatch(response, -1)
	if bgimages == nil {
		// log.Println("No matches.")
		return
	} else {
		for _, bgimage := range bgimages {
			Images["img"] = bgimage[1]
			fmt.Println(Images["img"])
			// CheckStatusOfImages(protocol, domain, Images["img"], errorCount)
			return
		}
	}
}

// Checking images into embed javascript
// func ChecksEmbedJavascript(pageContent string) {
// 	re := regexp.MustCompile(`<script>(.|\n)*?</script>`)
// 	scripts := re.FindAllString(string(pageContent), -1)
// 	if scripts == nil {
// 		fmt.Println("No matches.")
// 	} else {
// 		for _, script := range scripts {
// 			fmt.Println(script)
// 		}
// 	}

// }

//Checking Image tags in the body of the page
func ChecksImageInBody(pageContent string, protocol, domain string, errorCount int) {
	//replacing the tags to find only images tag
	TrimNoScript := strings.Replace(pageContent, "<noscript>", "", -1)
	TrimNoScript = strings.Replace(TrimNoScript, "</noscript>", "", -1)
	// --------------
	//searching for the all image tags inside of the hdata variable
	if document, err := html.Parse(strings.NewReader(TrimNoScript)); err == nil {
		var parser func(*html.Node)
		parser = func(n *html.Node) {
			if n.Type == html.ElementNode && n.Data == "img" {
				var imgSrcUrl string
				//searching for attributes of the image tags
				for _, element := range n.Attr {
					if element.Key == "src" {
						imgSrcUrl = element.Val
					}
					// if element.Key == "data-original" {
					// 	imgDataOriginal = element.Val
					// }
				}
				// fmt.Println(" ---------> " + imgSrcUrl)
				//running to finding only schema and host of the domain
				response, err := url.Parse(imgSrcUrl)
				if err != nil {
					log.Println("Failed to parsing Image path", err)
					return
				}
				path := response.Path
				Images["img"] = path
				// fmt.Println(imagepath)
				for _, imgurl := range Images {
					CheckStatusOfImages(protocol, domain, imgurl, errorCount)
				}
			}
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				parser(c)
			}
		}
		parser(document)
	} else {
		log.Println("Parse html error", err)
	}
}

//runExternalImageChecker checks the broken images
func runExternalImageChecker(Url string) {
	response, err := url.Parse(Url)
	if err != nil {
		log.Println("Failed to Parse Url", err)
		return
	}

	protocol := response.Scheme
	domain := response.Host
	errorCount := 0
	//accessing the url to fetch data from the page
	if resp, err := http.Get(Url); err == nil {
		//closing the body of the page after getting reponse
		defer resp.Body.Close()
		// log.Println("Load page complete")
		if resp != nil {
			// log.Println("Page response is NOT nil")
			//reading the data of the body
			pagebody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
				return
			}
			pageContent := string(pagebody)

			ChecksCssLinks(protocol, domain, pageContent)
			ChecksEmbedCss(pageContent, protocol, domain, errorCount)
			// ChecksEmbedJavascript(pageContent)
			ChecksImageInBody(pageContent, protocol, domain, errorCount)
		} else {
			log.Println(err)
		}
	}
	log.Println("Broken Image:", errorCount)
}

//Visit to css and javscript  url's
func CheckCssJavascriptLinks(protocol, domain, path string) {
	resp, err := http.Get(protocol + "://" + domain + path)
	if err != nil {
		log.Println(Red("Broken Url --> "), err)
		return
	}
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		// fmt.Println("status Ok")
		responsebody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("Not css Found")
			return
		}
		CheckBodyofCssFiles(string(responsebody))
	} else {
		// fmt.Println(resp.StatusCode)
		return
	}
}

//CheckStatusOfImages visiting to each image url throw Image map
func CheckStatusOfImages(protocol, domain, imgurl string, errorCount int) {
	resp, err := http.Get(protocol + "://" + domain + imgurl)
	if err != nil {
		log.Println(Red("Broken Url --> "), err)
		return
	}
	//Checks the status if status is not ok then it means image is broken
	//and print the path of the image
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		// fmt.Println(re)
		// fmt.Print("\n")
	} else {
		// log.Println(Red(imagepath["img"]))
		fmt.Println(CssLinks["css"])
		log.Println(Red(Images["img"]))
		errorCount++
		return
	}
}
