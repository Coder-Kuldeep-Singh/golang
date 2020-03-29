package main

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
)

var wg sync.WaitGroup

func main() {
	res, _ := http.Head("http://localhost/rand.txt") // 187 MB file of random numbers per line
	maps := res.Header
	length, _ := strconv.Atoi(maps["Content-Length"][0]) // Get the content length from the header request
	limit := 10                                          // 10 Go-routines for the process so each downloads 18.7MB
	len_sub := length / limit                            // Bytes for each Go-routine
	diff := length % limit                               // Get the remaining for the last request
	body := make([]string, 11)                           // Make up a temporary array to hold the data to be written to the file
	for i := 0; i < limit; i++ {
		wg.Add(1)

		min := len_sub * i       // Min range
		max := len_sub * (i + 1) // Max range

		if i == limit-1 {
			max += diff // Add the remaining bytes in the last request
		}

		go func(min int, max int, i int) {
			client := &http.Client{}
			req, _ := http.NewRequest("GET", "http://localhost/rand.txt", nil)
			range_header := "bytes=" + strconv.Itoa(min) + "-" + strconv.Itoa(max-1) // Add the data for the Range header of the form "bytes=0-100"
			req.Header.Add("Range", range_header)
			resp, _ := client.Do(req)
			// defer resp.Close()

			reader, _ := ioutil.ReadAll(resp.Body)
			body[i] = string(reader)
			ioutil.WriteFile(strconv.Itoa(i), []byte(string(body[i])), 0x777) // Write to the file i as a byte array
			wg.Done()
			//          ioutil.WriteFile("new_oct.png", []byte(string(body)), 0x777)
		}(min, max, i)
	}
	wg.Wait()
}
