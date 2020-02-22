package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
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
		log.Fatalln("Couldn't open the csv file", err)
	}
	defer csvfile.Close()
	// Parse the file
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
		fmt.Println(record)
	}
}

//TxtFileReader  reader the All data of the txt file
func TxtFileReader(filename string) {
	//Open the file
	txtfile, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Couldn't open the txt file", err)
	}
	// fmt.Println(txtfile)
	var lines []string
	//Parsing the txt file
	scanner := bufio.NewScanner(txtfile)
	//Iterate
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fmt.Println(lines)
}

//XmlFileReader reader the All data of the xml file
func XmlFileReader(filename string) {
	//Open the file
	xmlfile, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Couldn't open the xml file", err)
	}
	defer xmlfile.Close()
	// fmt.Println(xmlfile)
	var lines []string
	// Parsing the txt file
	scanner := bufio.NewScanner(xmlfile)
	// Iterate
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	fmt.Println(lines)
}

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

//JsonFileReader  reader the All data of the json file in array
func JsonFilereader(filename string) {
	//Open the file
	jsonFile, err := os.Open(filename)
	if err != nil {
		log.Fatalln("Couldn't open the json file", err)
	}
	defer jsonFile.Close()

	// we initialize our Users array
	var users Users

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
	}
	// fmt.Println(jsonFile)
	// var lines []string
	// //Parsing the txt file
	// scanner := bufio.NewScanner(jsonFile)
	// //Iterate
	// for scanner.Scan() {
	// 	lines = append(lines, scanner.Text())
	// }
	// fmt.Println(lines)
}
