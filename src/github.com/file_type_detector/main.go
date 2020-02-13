package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	filename := flag.String("f", "", "Provide the name of the file if it's exists")
	flag.Parse()
	//open the file
	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal("Error to open file", err)
	}
	defer file.Close()

	//read the content of the files
	typeofcontent, err := GetFileContentType(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Content Type: ", typeofcontent)
}

func GetFileContentType(out *os.File) (string, error) {
	buffer := make([]byte, 512)
	_, err := out.Read(buffer)
	if err != nil {
		return " ", err
	}
	// Use the net/http package's handy DectectContentType function. Always returns a valid
	// content-type by returning "application/octet-stream" if no others seemed to match.
	contentType := http.DetectContentType(buffer)
	return contentType, nil
}
