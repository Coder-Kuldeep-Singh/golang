package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
)

// // readLines reads a whole file into memory
// // and returns a slice of its lines.
func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
// func writeLines(lines []string, path string) error {
// 	file, err := os.Create(path)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()
// 	w := bufio.NewWriter(file)
// 	for _, line := range lines {
// 		fmt.Fprintln(w, line)
// 	}
// 	return w.Flush()
// }

func Domains(w http.ResponseWriter, r *http.Request) {
	lines, err := readLines("domains.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for _, line := range lines {
		filter := regexp.MustCompile(`domain_name=(.*.*) site_name`)
		stored := filter.FindAllStringSubmatch(line, -1)
		for _, element := range stored {
			fmt.Fprintln(w, element[1])
		}
	}
	// fmt.Fprintf(w, lines)
}

//func Develop() {
//	fmt.Println("Server Started on localhost:8000/v1/domains")
//}

func main() {
	http.HandleFunc("/v1/domains", Domains)
	fmt.Println("Development server started on localhost:8000/v1/domains")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
