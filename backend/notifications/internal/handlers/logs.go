package handlers

import (
	"encoding/json"
	"net/http"

	n "github.com/gila-software/backend/notifications"
)

type LogsMessagesResponse struct {
	Messages []n.LogggerMessage `json:"messages"`
}

func LogsMessagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	msgs := n.Storage.MessagesHistory()
	resp, _ := json.Marshal(&LogsMessagesResponse{msgs})
	w.Write(resp)
}
