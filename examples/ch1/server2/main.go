// server2 is a minimal echo and counter server
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Count struct {
	mu sync.Mutex
	count int
}

func main() {
	c := &Count{}
	log.Fatal(http.ListenAndServe("localhost:8000", c))
}

func (c *Count) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()
	router.HandleFunc("/", c.Handler)
	router.HandleFunc("/count", c.Count)
	router.ServeHTTP(w, r)
}

// Handler echoes the path  component of the requested url
func (c *Count) Handler(w http.ResponseWriter, r *http.Request) {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// Count echoes the number of calls so far
func (c *Count) Count(w http.ResponseWriter, r *http.Request) {
	c.mu.Lock()
	fmt.Fprintf(w, "Count = %d\n", c.count)
	c.mu.Unlock()
}
