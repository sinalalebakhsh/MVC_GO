

package core

import (
    "net/http"
)

type Server struct {
    // define any server configuration or settings here
}

func NewServer() *Server {
    return &Server{}
}

func (s *Server) Start() error {
    // start the HTTP server
    return http.ListenAndServe(":8080", nil)
}