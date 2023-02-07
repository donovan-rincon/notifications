package handlers

import (
	"encoding/json"
	"net/http"

	n "github.com/gila-software/backend/notifications"
)

type UsersResponse struct {
	Users []n.User `json:"users"`
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	users, _ := n.Storage.Users()
	resp, _ := json.Marshal(&UsersResponse{users})
	w.Write(resp)
}
