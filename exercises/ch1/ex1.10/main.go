package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)

	// create a file
	f, err := os.Create("output.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetchall: error when creating file %v\n", err)
		os.Exit(1)
	}

	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}

	for range os.Args[1:] {
		//fmt.Println(<- ch)
		f.WriteString(<-ch)
		f.WriteString("\n")
	}

	//fmt.Printf("%.2fs elapsed \n", time.Since(start).Seconds())

	f.Write([]byte(fmt.Sprintf("%.2fs elapsed \n", time.Since(start).Seconds())))
	f.Sync()
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	defer resp.Body.Close()
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}