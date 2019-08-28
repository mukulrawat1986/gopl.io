//server2 is a minimal echo and counter server
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type counter struct {
	mu sync.Mutex
	sum int
	http.Handler
}

func main() {
	//c := &counter{}
	c := NewCounter()
	log.Fatal(http.ListenAndServe("localhost:8000", c))
}

func NewCounter() *counter {
	c := new(counter)

	router := http.NewServeMux()
	router.HandleFunc("/", c.handler)
	router.HandleFunc("/count", c.count)

	c.Handler = router
	return c
}

//func (c *counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//	router := http.NewServeMux()
//	router.HandleFunc("/", c.handler)
//	router.HandleFunc("/count", c.count)
//
//	router.ServeHTTP(w, r)
//}

func (c *counter) handler(w http.ResponseWriter, r *http.Request) {
	c.mu.Lock()
	c.sum++
	c.mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func (c *counter) count (w http.ResponseWriter, r *http.Request) {
	c.mu.Lock()
	fmt.Fprintf(w, "Count %d\n", c.sum)
	c.mu.Unlock()
}
