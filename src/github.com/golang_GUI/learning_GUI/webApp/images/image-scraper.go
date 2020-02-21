package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

)

func main() {
	url := flag.String("u", " ", "Provide the link of the image you want to download")
	file := flag.String("fc", "", "Provide the name of the file")
	flag.Parse()
	response, err := http.Get(*url)
	if err != nil {
		fmt.Println("Having trouble to find url", err)
		os.Exit(1)
	}
	if response.StatusCode != http.StatusOK {
		fmt.Println(response.Status)
		os.Exit(1)
	}
	defer response.Body.Close()

	//create the file
	createfile, err := os.Create(*file)
	if err != nil {
		fmt.Println("Error to creating file", err)
		os.Exit(1)
	}
	defer createfile.Close()

	_, err = io.Copy(createfile, response.Body)
	if err != nil {
		fmt.Println("Error to append image into file")
		os.Exit(1)
	}
	fmt.Println("Success......")
}
