package main

import (
	"flag"
	"io"
	"net/http"
	"os"
)

func main() {
	url := flag.String("u", " ", "Provide the URL")
	filename := flag.String("f", "sample.html", "Provide the name of the which you want to store the content of the web page")
	flag.Parse()
	err := DownloadFile(*url, *filename)
	if err != nil {
		panic(err)
	}
}

func DownloadFile(url, filename string) error {
	//Create file
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	//Get the Data from web page
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	//Write the body into file
	_, err = io.Copy(out, response.Body)
	if err != nil {
		return err
	}
	return nil
}
