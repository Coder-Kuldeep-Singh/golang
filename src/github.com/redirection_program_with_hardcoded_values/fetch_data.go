package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// type Domains struct {
// 	domains_name string
// }
var Domains_List []string

func main() {
	fmt.Println("Starting the application...")
	response, err := http.Get("http://s.tutree.com:7635/v1/driver_websites")
	// response, err := http.Get("https://medium.com/@etiennerouzeaud/golang-string-to-array-string-dacc6b78a92e")

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	data, _ := ioutil.ReadAll(response.Body)
	Domain := append(Domains_List, string(data))
	fmt.Println(Domain)
	// for domains := 0; domains < len(Domain)-1; domains++ {
	// 	data_store1 := Domain[domains]              //store all domains inside of this variable
	// 	fmt.Println("<h1>" + data_store1 + "</h1>") //All domains name will print
	// 	Collect_url := [4]string{"/fbapply/", "/fbapply-i/", "/fbapply-dd/", "/fbapply-dd-a/"}
	// 	for url_segments := 0; url_segments < len(Collect_url); url_segments++ {
	// 		data_store2 := Collect_url[url_segments]
	// 		// fmt.Println("***********************************************************************************************")
	// 		fmt.Println("<h2>" + data_store2 + "</h2>")
	// 		// fmt.Println("<xmp>")
	// 		// Execute Command to collect Data
	// 		out, err := exec.Command("curl", data_store1+data_store2).Output()
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 		output := string(out)
	// 		re := regexp.MustCompile(`<script>;location.href='(.*)';</script>`)
	// 		submatchall := re.FindAllStringSubmatch(output, -1)
	// 		//fmt.Println(submatchall);
	// 		for _, element := range submatchall {
	// 			fmt.Println(element[1])
	// 		}

	// 	}
	// 	// fmt.Println("***********************************************************************************************")
	// 	fmt.Println("<h2>" + "Ping" + "</h2>")
	// 	out, err := exec.Command("curl", data_store1+"/open-positions/ping").Output()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	output := string(out)
	// 	ping_regex := regexp.MustCompile(`Version:(.*),`)
	// 	matchall := ping_regex.FindAllStringSubmatch(output, -1)
	// 	for _, element := range matchall {
	// 		fmt.Println("<ul>" + "<li>" + element[1] + "</li>" + "</ul>")
	// 	}
	// 	//output := string(out[20:36]) // Select range
	// 	// fmt.Println(data_store1)
	// 	fmt.Println("<hr/>")
	// 	// fmt.Println("***********************************************************************************************")
	// 	// fmt.Println("")

	// }
}
