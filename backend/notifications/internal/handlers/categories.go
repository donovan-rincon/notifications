package handlers

import (
	"encoding/json"
	"net/http"

	n "github.com/gila-software/backend/notifications"
)

type CategoriesResponse struct {
	Categories map[string]n.Category `json:"categories"`
}

func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	cat, _ := n.Storage.Categories()
	resp, _ := json.Marshal(&CategoriesResponse{cat})
	w.Write(resp)
}
