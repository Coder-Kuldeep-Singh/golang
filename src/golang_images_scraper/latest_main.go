//Checking images into embed css
func ChecksEmbedCss(pageContent string, protocol, domain string, errorCount int) {
	Openingtag := "<style>"
	ClosingTag := "</style>"
	Length := len(Openingtag)
	StartIndex := strings.Index(pageContent, Openingtag)
	if StartIndex == -1 {
		log.Println("No element found", StartIndex)
		return
	}
	//Find the index of the closing tag
	EndIndex := strings.Index(pageContent, ClosingTag)
	if EndIndex == -1 {
		log.Println("No closing tag found:", EndIndex)
		return
	}
	var Images = make(map[string]string)
	pageImages := []byte(pageContent[StartIndex+Length : EndIndex])
	re := regexp.MustCompile(`background-image: url(.*);`)
	bgimages := re.FindAllStringSubmatch(string(pageImages), -1)
	if bgimages == nil {
		// log.Println("No background Image Found in Embed Css")
		return
	} else {
		for _, bgimage := range bgimages {
			Images["img"] = bgimage[1]
			CheckStatusOfImages(protocol, domain, Images["img"], errorCount)
			// log.Println("background Image Found in Embed Css", Images["img"])
			return
		}
	}
}

//Checking images in css files
func ChecksCssLinks(protocol, domain, pageContent string, errorCount int) {
	re := regexp.MustCompile(`<link rel='stylesheet' href='(.*)'`)
	Links := re.FindAllStringSubmatch(pageContent, -1)
	if Links == nil {
		log.Println("No Css files Found.")
	} else {
		for _, link := range Links {
			CheckCssJavascriptLinks(protocol, domain, link[1], errorCount)
		}
	}

}

//CheckBodyOfCssFiles functions checks the background images in css files
func CheckBodyofCssFiles(protocol, domain, response string, errorCount int) {
	// fmt.Println(response)
	var Images = make(map[string]string)
	re := regexp.MustCompile(`background-image: url(.*);`)
	bgimages := re.FindAllStringSubmatch(response, -1)
	if bgimages == nil {
		// log.Println("No Background Image found in Css files")
		return
	} else {
		for _, bgimage := range bgimages {
			Images["img"] = bgimage[1]
			CheckStatusOfImages(protocol, domain, Images["img"], errorCount)
			return
		}
	}
}

//Checking images into embed javascript
// func ChecksEmbedJavascript(pageContent string) {
// 	re := regexp.MustCompile(`<script>(.|\n)*?</script>`)
// 	scripts := re.FindAllString(string(pageContent), -1)
// 	if scripts == nil {
// 		log.Println("No image found in embed javascript.")
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
	document, err := html.Parse(strings.NewReader(TrimNoScript))
	if err != nil {
		log.Println("Error to parse in html", err)
	}
	var Images = make(map[string]string)
	var ImageFinderinHtml func(*html.Node)
	//Searching for images tag inside of the html page
	ImageFinderinHtml = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			var imgSrcUrl string
			//searching for attributes of the image tags
			for _, element := range n.Attr {
				if element.Key == "src" {
					imgSrcUrl = element.Val
				}
			}
			response, err := url.Parse(imgSrcUrl)
			if err != nil {
				log.Println("Failed to parsing Image path", err)
				return
			}
			path := response.Path
			Images["img"] = path
			for _, imgurl := range Images {
				CheckStatusOfImages(protocol, domain, imgurl, errorCount)
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			ImageFinderinHtml(c)
		}
	}
	//Calling parser function for each url
	ImageFinderinHtml(document)
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
	resp, err := http.Get(Url)
	if err != nil {
		log.Println("Unable to Locate Url:", err)
	}
	defer resp.Body.Close()
	pagebody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Unable to fetch body of the page:", err)
	}
	pageContent := string(pagebody)
	ChecksCssLinks(protocol, domain, pageContent, errorCount)
	ChecksEmbedCss(pageContent, protocol, domain, errorCount)
	// ChecksEmbedJavascript(pageContent)
	ChecksImageInBody(pageContent, protocol, domain, errorCount)
	// log.Println("Broken Image:", errorCount)
}

//Visit to css and javscript  url's
func CheckCssJavascriptLinks(protocol, domain, path string, errorCount int) {
	resp, err := http.Get(protocol + "://" + domain + path)
	if err != nil {
		log.Println(Red("Broken Url --> "), err)
		return
	}
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		// fmt.Println("status Ok")
		responsebody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Println("No content found in Pages", err)
			return
		}
		CheckBodyofCssFiles(protocol, domain, string(responsebody), errorCount)
	} else {
		log.Println("response code of the body", resp.Status, resp)
		return
	}
}

//CheckStatusOfImages visiting to each image url throw Image map
func CheckStatusOfImages(protocol, domain, imgurl string, errorCount int) {
	resp, err := http.Get(protocol + "://" + domain + imgurl)
	if err != nil {
		log.Println(Red("Invalid Url: " + imgurl))
		return
	}
	//Checks the status if status is not ok then it means image is broken
	//and print the path of the image
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		// fmt.Println(re)
		// fmt.Print("\n")
	} else {
		log.Println(Red("Broken Image: " + imgurl))
		errorCount++
		return
	}
}
