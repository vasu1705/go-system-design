package typicode

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type TypicodeHandler struct {
	service *TypicodeService
}

func NewTypicodeHandler(service *TypicodeService) *TypicodeHandler {
	return &TypicodeHandler{
		service: service,
	}
}

func (th *TypicodeHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Received request for /users")
	users, err := th.service.GetUsers()
	if err != nil {
		http.Error(w, "Failed to fetch users", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Failed to encode users", http.StatusInternalServerError)
		return
	}
}
