package server

import (
	"fmt"
	"net/http"
	"playground/internal/handler"
)

type Server struct {
	mux *http.ServeMux
}

func NewServer() *Server {
	return &Server{
		mux: http.NewServeMux(),
	}
}


func (s *Server) RegisterRouters(userhandler *handler.UserHandler) {
	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "API is running")
	})

	s.mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodGet {
			userhandler.GetUsers(w, r)		
			return
		}
	
		if r.Method == http.MethodPost {
			userhandler.CreateUser(w, r)
			return
		}

		if r.Method == http.MethodGet{
			userhandler.GetUserById(w, r)
			return
		}
	
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})
}

func (s *Server) Run(addr string) error {
	return http.ListenAndServe(addr, s.mux)
}
