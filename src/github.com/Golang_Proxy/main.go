package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {

	//Creating the ProxyURL
	proxtStr := "http://localhost:7000"
	proxyURL, err := url.Parse(proxtStr)
	if err != nil {
		log.Println(err)
	}

	//Creating the URL to be loaded through the proxy
	urlStr := "https://www.admfactory.com/how-to-setup-a-proxy-for-http-client-in-golang/"
	url, err := url.Parse(urlStr)
	if err != nil {
		log.Println(err)
	}

	//Adding the proxy settings to the Transport object
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	//adding the Transport object to the http Client
	client := &http.Client{
		Transport: transport,
	}

	//Generating the HTTP GET request
	request, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		log.Println(err)
	}

	//Calling the URL
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}

	//Getting the response
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println(err)
	}

	//Printing the response
	log.Println(string(data))

}
