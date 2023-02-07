package main

import (
	"net/http"

	n "github.com/gila-software/backend/notifications"
	"github.com/gila-software/backend/notifications/internal/handlers"
	notificationservices "github.com/gila-software/backend/notifications/internal/notification-services"
	"github.com/gila-software/backend/notifications/internal/notification-services/email"
	pushnotifications "github.com/gila-software/backend/notifications/internal/notification-services/push-notifications"
	"github.com/gila-software/backend/notifications/internal/notification-services/sms"
	"github.com/gila-software/backend/notifications/internal/storage/localstorage"
)

func main() {
	notificationservices.Init()
	ls := localstorage.NewLocalStorage()
	n.Storage = &ls

	n.Notifiers = []n.NotificationService{
		&sms.NotificationService{
			Storage: n.Storage,
		},
		&pushnotifications.NotificationService{
			Storage: n.Storage,
		},
		&email.NotificationService{
			Storage: n.Storage,
		},
	}

	http.HandleFunc("/message", handlers.MessagesHandler)
	http.HandleFunc("/messages", handlers.MessagesHandler)
	http.HandleFunc("/logs/messages", handlers.LogsMessagesHandler)
	http.HandleFunc("/categories", handlers.CategoriesHandler)
	http.HandleFunc("/users", handlers.UsersHandler)
	http.ListenAndServe(":8080", handlers.LoggingMiddleware(http.DefaultServeMux))
}
