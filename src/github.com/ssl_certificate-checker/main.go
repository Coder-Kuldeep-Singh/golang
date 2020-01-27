package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"strings"
	"time"
)

// type Expired struct {
// 	Domains []string
// }

// type Expired struct {
// 	Domains string
// }

func FetchDomains() {
	//fetch data from given url
	response, err := http.Get("http://s.tutree.com:7635/v1/groups")
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	// Read data from url
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	storeDomains := string(body)
	trimdata := strings.Split(storeDomains, "\n")
	for _, url := range trimdata {
		response, err := http.Get("http://s.tutree.com:7635/v1/" + url)
		if err != nil {
			log.Fatal(err)
		}

		defer response.Body.Close()
		// Read data from url
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		storeDomains := string(body)
		trimdata := strings.Split(storeDomains, "\n")
		for _, url := range trimdata {
			checkURL(url)

		}
	}
}

func main() {
	// var urls = []string{
	// 	"https://chaufferjob.us",
	// 	"https://qkids.com/",
	// }
	FetchDomains()
	// urls := []string FetchDomains(url)
	// for _, url := range urls {
	// 	checkURL(url)
	// }
}

func checkURL(url string) {
	resp, err := http.Head("https://" + url)
	// errorCounts := 0
	if err != nil {
		log.Printf("Unable to get %q: %s\n", url, err)
		return
	}
	resp.Body.Close()
	if resp.TLS == nil {
		log.Printf("%q is not HTTPS\n", url)
		return
	}

	for _, cert := range resp.TLS.PeerCertificates {
		for _, name := range cert.DNSNames {
			if !strings.Contains(url, name) {
				continue
			}
			issuer := strings.Join(cert.Issuer.Organization, ", ")
			dur := cert.NotAfter.Sub(time.Now())
			expiredate := dur.Hours() / 24
			dates := cert.NotAfter
			fmt.Printf("Certificate for %q from %q expires %s (%.0f days).\n", name, issuer, dates, expiredate)
			if expiredate <= 0 {
				// Sending  email to admin
				// changeType := strconv.Itoa(dates)
				// changeType2 := strconv.Itoa(expiredate)
				// Expired = append(Expired, name)
				// errorCounts = strings(name)
				body := "Certificate for " + name + " from " + issuer + " Expired "
				// body := Expired{Domains: name}
				sendEmail(body)
			}
		}
	}
}

//sendEmail function sends msg to admin
func sendEmail(body string) {
	// body := "Certificate for " + name + " from " + issuer + " Expired" + dates + expiredate
	from := os.Getenv("FROM")
	pass := os.Getenv("PASS")
	to := os.Getenv("TO")
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: Certificate Expire Alert\n" +
		body
	err := smtp.SendMail("smtp.gmail.com:587",
		smtp.PlainAuth("", from, pass, "smtp.gmail.com"),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("smtp error: %s", err)
		return
	}
	log.Print("sent, visit ", to)
}
