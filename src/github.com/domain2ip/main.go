package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

type Domains struct {
	Domain string `json:"domain"`
	IP     string `json:"ip"`
}

func main() {
	// var data []Domains
	url, err := http.Get("http://s.tutree.com:7635/v1/groups")
	Error(err)
	defer url.Body.Close()
	body, err := ioutil.ReadAll(url.Body)
	Error(err)
	converter := string(body)
	trimdata := strings.Split(converter, "\n")
	// fmt.Println(trimdata)
	for _, data := range trimdata {
		urls, err := http.Get("http://s.tutree.com:7635/v1/" + data)
		Error(err)
		defer urls.Body.Close()
		data, err := ioutil.ReadAll(urls.Body)
		Error(err)
		converter := string(data)
		trim := strings.Split(converter, "\n")
		// fmt.Println(converter)
		for _, domains := range trim {
			// fmt.Println(string(domains))
			convertdomain2ip(string(domains))
			// writeFile(domains)
		}
	}
	// fmt.Println("file written successfully")
	// write, _ := json.MarshalIndent(data, "", "")
	// // fmt.Println(string(write))
	// _ = ioutil.WriteFile("domain2ip.json", write, 0644)
	os.Exit(0)
}

func Error(err error) {
	if err != nil {
		log.Println(err)
	}
}

//Write data into json file
func convertdomain2ip(name string) {
	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		// os.Exit()
	}
	typeconvert := addr.String()
	var data []Domains
	// for _, names := range trimagain {
	data = append(data, Domains{
		Domain: name,
		IP:     typeconvert,
	})
	// }
	write, _ := json.MarshalIndent(data, "", "")
	// fmt.Println(string(write))
	_ = ioutil.WriteFile("domain2ip.json", write, 0644)
	fmt.Println("domain=" + name + "\t" + "ip=" + addr.String())
}

//WriteFile function writing data into a txt file
// func writeFile(data string) {
// 	// addr, err := net.ResolveIPAddr("ip", data)
// 	// if err != nil {
// 	// 	fmt.Println("Resolution error", err.Error())
// 	// 	// os.Exit()
// 	// }
// 	f, err := os.Create("address.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 		f.Close()
// 		return
// 	}
// 	trim := string(data)
// 	trimagain := strings.Split(trim, "\n")
// 	for i, v := range trimagain {
// 		fmt.Fprintln(f, v[i])
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 	}
// 	err = f.Close()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// }
