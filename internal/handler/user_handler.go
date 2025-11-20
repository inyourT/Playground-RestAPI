package handler

import (
	"encoding/json"
	"net/http"
	"playground/internal/service"
	"strconv"
	"strings"
)

type UserHandler struct {
	service *service.UserService
}

type CreateUserRequest struct {
	Name string `json:"name"`
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users := h.service.GetAllUser()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req CreateUserRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.Name == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	newUser := h.service.CreateUser(req.Name)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	// example url: /users/3
	path := r.URL.Path
	parts := strings.Split(path, "/")

	if len(parts) != 3 {
		http.Error(w, "ivalid URL", http.StatusBadRequest)
	}

	idStr := parts[2]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id must be a nomber", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUserById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
