package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	n "github.com/gila-software/backend/notifications"
	"github.com/google/uuid"
)

type payload struct {
	Categories []string `json:"categories"`
	Content    string   `json:"content"`
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Received request:", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(r.Body)
	var p payload
	err := decoder.Decode(&p)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	categories := []n.Category{}
	for _, c := range p.Categories {
		switch c {
		case n.Finance.String():
			categories = append(categories, n.Finance)
		case n.Movies.String():
			categories = append(categories, n.Movies)
		case n.Sports.String():
			categories = append(categories, n.Sports)

		}
	}

	msg := n.Message{
		ID:         uuid.NewString(),
		Categories: categories,
		Content:    p.Content,
	}
	n.Storage.StoreMessage(msg)
	for _, notifier := range n.Notifiers {
		notifier.Notify(msg)
	}
}
