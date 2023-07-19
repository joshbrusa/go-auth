package api

import (
	"net/http"
	"strconv"
)

type Server struct {
	port int
}

func NewServer(port int) *Server {
	return &Server{port}
}

func (server *Server) Start() error {
	addr := ":" + strconv.Itoa(server.port)
	return http.ListenAndServe(addr, nil)
}

