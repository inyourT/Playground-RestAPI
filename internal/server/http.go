package server

import (
	"net/http"
	"playground/internal/handler"

	"github.com/go-chi/chi"
)

type Server struct {
	router chi.Router
}

func NewServer() *Server {
	return &Server{
		router: chi.NewRouter(),
	}
}

func (s *Server) RegisterRouters(userHandler *handler.UserHandler) {
	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is running"))
	})

	s.router.Route("/users", func(r chi.Router) {
		r.Get("/", userHandler.GetUsers)
		r.Post("/", userHandler.CreateUser)
		r.Get("/{id}", userHandler.GetUserById)
		r.Put("/{id}", userHandler.UpdateUser)
	})
}

func (s *Server) Run(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
