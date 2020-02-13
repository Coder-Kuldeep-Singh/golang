package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	filename := flag.String("f", "", "Provide the name|path of the file")
	flag.Parse()
	fileextension := filepath.Ext(*filename)
	// fmt.Println(fileextension)
	if fileextension == ".csv" {
		CsvFileReader(*filename)
	} else if fileextension == ".txt" {
		TxtFileReader(*filename)
	} else if fileextension == ".xml" {
		XmlFileReader(*filename)
	} else if fileextension == ".json" {
		JsonFilereader(*filename)
	} else {
		fmt.Println("This Tool doesn't support " + fileextension + " file")
	}

	// fmt.Println("Extension of the file:", fileextension)
}

//CsvFileReader  reader the All data of the csv file
func CsvFileReader(filename string) {
	// Open the file
	csvfile, err := scanLines(filename)
	if err != nil {
		panic(err)
	}
	for _, line := range csvfile {
		fmt.Println(line)
	}
}

//TxtFileReader  reader the All data of the txt file
func TxtFileReader(filename string) {
	//Open the file
	txtfile, err := scanLines(filename)
	if err != nil {
		panic(err)
	}
	for _, line := range txtfile {
		fmt.Println(line)
	}

}

//XmlFileReader reader the All data of the xml file
func XmlFileReader(filename string) {
	//Open the file
	xmlfile, err := scanLines(filename)
	if err != nil {
		panic(err)
	}
	for _, line := range xmlfile {
		fmt.Println(line)
	}
}

//JsonFileReader  reader the All data of the json file in array
func JsonFilereader(filename string) {
	//Open the file
	jsonFile, err := scanLines(filename)
	if err != nil {
		panic(err)
	}
	for _, line := range jsonFile {
		fmt.Println(line)
	}
}

//scanLines line by line in all format files
func scanLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}
