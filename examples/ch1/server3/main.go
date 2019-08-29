package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Counter struct {
	mu sync.Mutex
	count int
	http.Handler
}

func main() {
	c := NewCount()
	log.Fatal(http.ListenAndServe("localhost:8000", c))
}

func NewCount() *Counter {
	c := new(Counter)

	router := http.NewServeMux()
	router.HandleFunc("/", c.Echo)
	router.HandleFunc("/count", c.Count)
	c.Handler = router
	return c
}

// Echo echoes the Path component of the requested url
func (c  *Counter) Echo(w http.ResponseWriter, r *http.Request) {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// count echoes the number of calls so far
func (c *Counter) Count(w http.ResponseWriter, r *http.Request) {
	c.mu.Lock()
	fmt.Fprintf(w, "Count = %d\n", c.count)
	c.mu.Unlock()
}