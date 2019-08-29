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
	http.Handler
}

func main() {
	s := NewServer()
	log.Fatal(http.ListenAndServe("localhost:8000", s))
}

func NewServer() *Server {
	s := new(Server)

	router := http.NewServeMux()
	router.HandleFunc("/", s.Echo)
	router.HandleFunc("/count", s.Counter)
	s.Handler = router
	return s
}

// Echo echoes the Path component of the requested url
func (s  *Server) Echo(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	s.count++
	s.mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// count echoes the number of calls so far
func (s *Server) Counter(w http.ResponseWriter, r *http.Request) {
	s.mu.Lock()
	fmt.Fprintf(w, "Counter = %d\n", s.count)
	s.mu.Unlock()
}