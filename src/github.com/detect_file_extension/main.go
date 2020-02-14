package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
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
	csvfile, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer csvfile.Close()
	// str2 := strings.Join(csvfile, " ")
	// str2 = string(str2)
	// t := strings.Split(str2, "\n")
	// fmt.Println(string(t[1]))
	// for _, line := range csvfile {
	// 	str := string(line)
	//
	// fmt.Println(string(t))
	// for _, single := range t {
	// 	fmt.Println(string(single))
	// }
	// }
	r := csv.NewReader(csvfile)
	// Iterate through the records
	for {
		// Read each record from csv
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(record[0])
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
