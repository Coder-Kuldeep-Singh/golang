package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
)

func readCurrentDir(folderpath string) {
	file, err := os.Open(folderpath)
	if err != nil {
		log.Println("failed opening directory: ", err)
		os.Exit(1)
	}
	defer file.Close()

	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	// fmt.Println(len(list))
	for _, name := range list {
		// fmt.Println(name)
		// OpentxtFiles(folderpath, name)
		getDomainFromFile(name)
	}
}

// func OpentxtFiles(folderpath, filename string) {
// 	file, err := os.Open(folderpath + "/" + filename)
// 	if err != nil {
// 		log.Println("failed opening directory: ", err)
// 	}
// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		// log.Println(scanner.Text())
// 		getDomainFromFile(string(scanner.Text()))
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Println(err)
// 		return
// 	}
// }

func CreateDomainsFile(domains string) {
	filename, err := os.OpenFile("All-Sitemap-Domains.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("Error to create txt file", err)
	}
	defer filename.Close()
	_, err = filename.WriteString(domains + "\n")
	if err != nil {
		log.Println("Error to append data into txt file", err)
	}
	filename.Sync()
}

func getDomainFromFile(PageContent string) {
	re := regexp.MustCompile(`(.*)-robots.txt`)
	// re := regexp.MustCompile(`Sitemap: (.*)`)
	FileToDomain := re.FindAllStringSubmatch(PageContent, -1)
	for _, Domain := range FileToDomain {
		// log.Println(Domain[1])
		// log.Println(Domain[1])
		CreateDomainsFile(string(Domain[1]))
		return
	}

}

func showData(w http.ResponseWriter, r *http.Request) {
	// (w, "Server")
	readfile, err := os.Open("All-Sitemap-Domains.txt")
	if err != nil {
		log.Println("Error to open the Domains file", err)
	}
	scanner := bufio.NewScanner(readfile)
	for scanner.Scan() {
		// log.Println(scanner.Text())
		fmt.Fprintf(w, scanner.Text()+"\n")
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
		return
	}
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Ping")
}
func main() {
	folder := flag.String("f", "", "Provide the path of the folder")
	flag.Parse()
	readCurrentDir(*folder)
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/sitemap_Domains/", showData)
	fmt.Println("Development Server Started localhost:8081")
	http.ListenAndServe(":8081", nil)

}
