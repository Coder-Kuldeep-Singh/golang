package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"strings"
	"time"
)

func main() {

	var urls = []string{
		"https://chaufferjob.us",
		"https://qkids.com/",
	}

	for _, url := range urls {
		checkURL(url)
	}
}

func checkURL(url string) {
	resp, err := http.Head(url)
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
				body := "Certificate for " + name + " from " + issuer + " Expired"
				from := "somebody@gmail.com"
				pass := ""
				to := "somebody@gmail.com"
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
		}
	}
}
