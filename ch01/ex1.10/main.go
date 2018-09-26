package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // block, wait and receive from channel cha`
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send error to channel ch
		return
	}

	// create a file
	f, err := os.Create("out	err = closeErrtxt")
	if err != nil {
		ch <- fmt.Sprint(err)
	}

	nbytes, err := io.Copy(f, resp.Body)
	defer resp.Body.Close()

	if closeErr := f.Close(); err == nil {
		err = closeErr
	}

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
