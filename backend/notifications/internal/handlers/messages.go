package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	n "github.com/gila-software/backend/notifications"
	"github.com/google/uuid"
)

type payload struct {
	Categories []n.Category `json:"categories"`
	Content    string       `json:"content"`
}

func postMessage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var p payload
	err := decoder.Decode(&p)

	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	categories := []n.Category{}
	categories = append(categories, p.Categories...)

	msg := n.Message{
		ID:         uuid.NewString(),
		Categories: categories,
		Content:    p.Content,
	}
	n.Storage.StoreMessage(msg)
	for _, notifier := range n.Notifiers {
		log.Println(notifier)
		notifier.Notify(msg)
	}
}

type MessagesResponse struct {
	Messages []n.Message `json:"messages"`
}

func getMessages(w http.ResponseWriter, r *http.Request) {
	msgs, _ := n.Storage.Messages()
	resp, _ := json.Marshal(&MessagesResponse{msgs})
	w.Write(resp)
}

func MessagesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		postMessage(w, r)
		return
	}
	if r.Method == http.MethodGet {
		getMessages(w, r)
		return
	}
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}
