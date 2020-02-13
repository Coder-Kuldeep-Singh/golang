package main

import (
	"bufio"
	"flag"
	"fmt"
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
	csvfile, err := scanLines(filename)
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	for _, line := range csvfile {
		fmt.Println(line)
	}
	// defer csvfile.Close()
	// // Parse the file
	// r := csv.NewReader(csvfile)
	// // Iterate through the records
	// for {
	// 	// Read each record from csv
	// 	record, err := r.Read()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	fmt.Println(record)
	// }
}

//TxtFileReader  reader the All data of the txt file
func TxtFileReader(filename string) {
	//Open the file
	txtfile, err := scanLines(filename)
	if err != nil {
		log.Fatalln("Couldn't open the txt file", err)
	}
	// fmt.Println(txtfile)
	// var lines []string
	// //Parsing the txt file
	// scanner := bufio.NewScanner(txtfile)
	// //Iterate
	// for scanner.Scan() {
	// 	lines = append(lines, scanner.Text())
	// }
	// fmt.Println(lines)
	for _, line := range txtfile {
		fmt.Println(line)
	}

}

//XmlFileReader reader the All data of the xml file
func XmlFileReader(filename string) {
	//Open the file
	xmlfile, err := scanLines(filename)
	if err != nil {
		log.Fatalln("Couldn't open the xml file", err)
	}
	// defer xmlfile.Close()
	// fmt.Println(xmlfile)
	// var lines []string
	// // Parsing the txt file
	// scanner := bufio.NewScanner(xmlfile)
	// // Iterate
	// for scanner.Scan() {
	// 	lines = append(lines, scanner.Text())
	// }
	// fmt.Println(lines)
	for _, line := range xmlfile {
		fmt.Println(line)
	}
}

//JsonFileReader  reader the All data of the json file in array
func JsonFilereader(filename string) {
	//Open the file
	jsonFile, err := scanLines(filename)
	if err != nil {
		log.Fatalln("Couldn't open the json file", err)
	}
	for _, line := range jsonFile {
		fmt.Println(line)
	}
	// defer jsonFile.Close()
	// // fmt.Println(jsonFile)
	// var lines []string
	// //Parsing the txt file
	// scanner := bufio.NewScanner(jsonFile)
	// //Iterate
	// for scanner.Scan() {
	// 	lines = append(lines, scanner.Text())
	// }
	// fmt.Println(lines)
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
