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

const (
	filename string = "domain2ip.json"
)

type DomainIP struct {
	Domain string `json:"domain"`
	IP     string `json:"ip"`
}

func main() {
	trimdata := lines(fetchUrl("http://s.tutree.com:7635/v1/groups"))
	// fmt.Println(trimdata)
	for _, data := range trimdata {
		trim := lines(fetchUrl("http://s.tutree.com:7635/v1/" + data))
		// fmt.Println(trim)
		for _, domains := range trim {
			// fmt.Println(string(domains))
			domainIP, err := convertdomain2ip(string(domains))
			if err == nil {
				appendToFile(filename, domainIP)
			}
		}
	}
	// fmt.Println("file written successfully")
	os.Exit(0)
}

func fetchUrl(url string) string {
	result, err := http.Get(url)
	Error(err)
	defer result.Body.Close()
	body, err := ioutil.ReadAll(result.Body)
	Error(err)
	return string(body)
}

func lines(s string) []string {
	return strings.Split(s, "\n")
}

func Error(err error) {
	if err != nil {
		log.Println(err)
	}
}

func convertdomain2ip(name string) (*DomainIP, error) {
	addr, err := net.ResolveIPAddr("ip", name)
	if err != nil {
		fmt.Println("Resolution error", err.Error())
		return nil, err
	}
	typeconvert := addr.String()
	return &DomainIP{
		Domain: name,
		IP:     typeconvert}, nil
}

func appendToFile(filename string, data interface{}) {
	jsonData, _ := json.MarshalIndent(data, "", "")
	// fmt.Println(string(write))
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.Write(jsonData); err != nil {
		log.Println(err)
	}
}
