package main

import (
	"log"
	"os/exec"
	"time"

	"github.com/raff/godet"
)

func main() {
	chromeApp := "chromium-browser"
	chromeAppArg := []string{"--headless", "--hide-scrollbars", "--remote-debugging-port=9222", "--disable-gpu", "--allow-insecure-localhost"}
	cmd := exec.Command(chromeApp, chromeAppArg...)
	err := cmd.Start()
	if err != nil {
		log.Println("cannot start browser", err)
	}
	// Will wait for chromium to start
	time.Sleep(5 * time.Second)

	/// connect to Chromium instance
	remote, err := godet.Connect("localhost:9222", true)
	if err != nil {
		log.Println("cannot connect to Chrome instance:", err)
		return
	}

	// disconnect when done
	defer remote.Close()
	remote.PageEvents(true)
	remote.DOMEvents(true)

	_, err = remote.Navigate("https://www.google.com")
	if err != nil {
		log.Println("cannot connect to Chrome instance:", err)
		return
	}
	count := 0

	remote.CallbackEvent("Page.frameStoppedLoading", func(params godet.Params) {
		count = count + 1
		log.Println("TCL: remote.CallbackEvent -> count", count)
		switch count {
		case 1:
			_, err = remote.EvaluateWrap(`
                document.getElementsByClassName('gLFyf gsfi')[0].value = "Bacancy Technology"
                document.getElementsByClassName('gNO89b')[1].click()
            `)
			if err != nil {
				log.Println("Error executing js block 1:", err)
				return
			}
			time.Sleep(5 * time.Second)
		case 2:
			_, err = remote.EvaluateWrap(`
                document.getElementsByClassName('LC20lb')[0].click()
            `)
			if err != nil {
				log.Println("Error executing js block 1:", err)
				return
			}
			time.Sleep(5 * time.Second)
			// take a screenshot
			_ = remote.SaveScreenshot("SearchResultScreenshot.png", 0644, 0, true)
			// print a pdf
			err = remote.SavePDF(`./SearchResultPdf.pdf`, 0644)
			if err != nil {
				log.Println("Error saving pdf.")
			}
		}
	})

	time.Sleep(30 * time.Second)

	killApp := "kill"
	killAppArg := []string{"$(lsof -t -i:9222)"}
	cmd = exec.Command(killApp, killAppArg...)
	err = cmd.Start()
	if err != nil {
		log.Println("cannot kill processes", err)
	}
}
