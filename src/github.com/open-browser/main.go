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
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}