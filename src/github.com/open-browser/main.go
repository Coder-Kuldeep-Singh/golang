package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"runtime"
)

func main() {
	url := flag.String("u", "", "Provide the  Url")
	flag.Parse()
	OpenBrowser(*url)
}

func OpenBrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err := exec.Command("xdg-open", url)
		err.Start()
		// if error := err.Process.Kill(); error != nil {
		// 	log.Fatal("close the opened browser failed", error)
		// }
	case "windows":
		err := exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
		err.Start()
		// if error := err.Process.Kill(); error != nil {
		// 	log.Fatal("close the opened browser failed", error)
		// }
	case "darwin":
		err := exec.Command("open", url)
		err.Start()
		// if error := err.Process.Kill(); error != nil {
		// 	log.Fatal("close the opened browser failed", error)
		// }
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
