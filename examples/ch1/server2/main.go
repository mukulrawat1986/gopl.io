// server2 is a minimal echo and counter server
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Server struct {
	mu sync.Mutex
	count int
}

func main() {
	s := &Server{}
	log.Fatal(http.ListenAndServe("localhost:8000", s))
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()
	router.HandleFunc("/", s.Echo)
	router.HandleFunc("/count", s.Counter)
	router.ServeHTTP(w, r)
}

// Echo echoes the path  component of the requested url
func (s *Server) Echo(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	s.count++
	s.mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// Counter echoes the number of calls so far
func (s *Server) Counter(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	fmt.Fprintf(w, "Counter = %d\n", s.count)
	s.mu.Unlock()
}
