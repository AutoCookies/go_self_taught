package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()

	// Get the HTTP response
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("error fetching %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	// Determine a filename based on the URL
	fileName := path.Base(resp.Request.URL.Path)
	if fileName == "" || fileName == "/" {
		fileName = "index.html" // fallback
	}

	// Create the file
	outFile, err := os.Create(fileName)
	if err != nil {
		ch <- fmt.Sprintf("error creating file for %s: %v", url, err)
		return
	}
	defer outFile.Close()

	// Copy HTTP response body to file
	nbytes, err := io.Copy(outFile, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("error writing %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d bytes written to %s", secs, nbytes, fileName)
}
