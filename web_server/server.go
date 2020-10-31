package main

import "net/http"

type Server struct {
	port   string
	router *Router
}

func (s *Server) Handle(path string, handler http.HandlerFunc) {
	s.router.rules[path] = handler
}

func (s *Server) ListenServer() error {
	http.Handle("/", s.router)

	return http.ListenAndServe(s.port, nil)
}

func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}
