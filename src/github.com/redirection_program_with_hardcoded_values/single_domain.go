package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	_ "strings"
)

func main() {
	FetchData()
}

func FetchData() {
	var Domain_List [1]string
	// Domain_List[0] = "https://driverus.us"
	// Domain_List[1] = "https://chaufferus.us"
	// Domain_List[2] = "https://chaufferjob.us"
	// Domain_List[3] = "https://driveus.us"
	// Domain_List[4] = "https://ownboss.us"
	// Domain_List[5] = "https://driverjob.us"
	// Domain_List[6] = "https://doordash.givearide.us"
	// Domain_List[7] = "https://tutree.com"
	// Domain_List[8] = "https://givearide.us"
	Domain_List[0] = "https://51talkus.com"
	// Domain_List[10] = "https://sprout4future.us"
	// Domain_List[11] = "https://rouchi.us"
	// Domain_List[12] = "https://qkids.co"
	// Domain_List[13] = "https://delivery.givearide.us"
	//Driver_websites
	for domains := 0; domains < len(Domain_List); domains++ {
		data_store1 := Domain_List[domains]         //store all domains inside of this variable
		fmt.Println("<h1>" + data_store1 + "</h1>") //All domains name will print
		Collect_url := [6]string{"/fbapply/", "/fbapply-i/", "/fbapply-dd/", "/fbapply-dd-a/", "/fbapply-t/", "/fbapply-v/"}
		for url_segments := 0; url_segments < len(Collect_url); url_segments++ {
			data_store2 := Collect_url[url_segments]
			// fmt.Println("***********************************************************************************************")
			fmt.Println("<h2>" + data_store2 + "</h2>")
			// fmt.Println("<xmp>")

			// Execute Command to collect Data
			// out, err := exec.Command("curl", data_store1+data_store2).Output()
			out, err := http.Get(data_store1 + data_store2)
			if err != nil {
				fmt.Println(err)
			}
			// Print the HTTP Status Code and Status Name
			// fmt.Println("HTTP Response Status:", out.StatusCode, http.StatusText(out.StatusCode))

			if out.StatusCode >= 200 && out.StatusCode <= 299 {
				// fmt.Println("HTTP Status is in the 2xx range")
				body, err := ioutil.ReadAll(out.Body)
				if err != nil {
					fmt.Println(err)
				}
				output := string(body)
				// re := regexp.MustCompile(`location.href(.*);`)
				re := regexp.MustCompile(`content="5;(.*)"`)
				// re := regexp.MustCompile(`http://(.*)`)
				submatchall := re.FindAllStringSubmatch(output, -1)
				for _, element := range submatchall {
					fmt.Println(element[1])
				}
			} else {
				// fmt.Println("Argh! Broken")
				fmt.Println("HTTP Response Status:", out.StatusCode, http.StatusText(out.StatusCode))
			}

		}
		// fmt.Println("***********************************************************************************************")
		fmt.Println("<h2>" + "Ping" + "</h2>")
		// Collect data to see versions
		//out, err := exec.Command("curl", data_store1+"/open-positions/ping").Output()
		out, err := http.Get(data_store1 + "/open-positions/ping")
		if err != nil {
			fmt.Println(err)
		}

		// fmt.Println("HTTP Response Status:", out.StatusCode, http.StatusText(out.StatusCode))

		if out.StatusCode >= 200 && out.StatusCode <= 299 {
			// fmt.Println("HTTP Status is in the 2xx range")
			body, err := ioutil.ReadAll(out.Body)
			if err != nil {
				fmt.Println(err)
			}
			output := string(body)
			ping_regex := regexp.MustCompile(`Version:(.*),`)
			matchall := ping_regex.FindAllStringSubmatch(output, -1)
			for _, element := range matchall {
				fmt.Println("<ul>" + "<li>" + element[1] + "</li>" + "</ul>")
			}

		} else {
			// fmt.Println("Argh! Broken")
			fmt.Println("HTTP Response Status:", out.StatusCode, http.StatusText(out.StatusCode))
		}

		//output := string(out)
		//
		//
		//
		//output := string(out[20:36]) // Select range
		// fmt.Println(data_store1)
		fmt.Println("<hr/>")
		// fmt.Println("***********************************************************************************************")
		// fmt.Println("")

	}

}
