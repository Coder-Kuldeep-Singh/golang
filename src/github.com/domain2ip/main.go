package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

// type Domains struct {
// 	Domain []string `json:"domain"`
// 	// IP     []int    `json:"ip"`
// }

func main() {
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
		}
	}
	os.Exit(0)
}

func Error(err error) {
	if err != nil {
		log.Println(err)
	}
}

func convertdomain2ip(name string) {
	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		// os.Exit()
	}
	// var data []Domains
	// // for _, names := range name {
	// data = append(data, Domains{
	// 	Domain: []string{name},
	// 	// IP:     []int{addr},
	// })
	// // }
	// write, _ := json.MarshalIndent(data, "", "")
	// fmt.Println(string(write))
	// _ = ioutil.WriteFile("domain2ip.json", write, 0644)
	fmt.Println("domain=" + name + "\t" + "ip=" + addr.String())
}
