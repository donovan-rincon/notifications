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

	/*
		msgs := []n.Message{
			{
				ID:         "1",
				Categories: []n.Categories{n.Finance, n.Sports},
				Content:    "Test content 1",
			},
			{
				ID:         "2",
				Categories: []n.Categories{n.Movies, n.Sports},
				Content:    "Test content 2",
			},
			{
				ID:         "3",
				Categories: []n.Categories{n.Movies},
				Content:    "Test content 3",
			},
			{
				ID:         "4",
				Categories: []n.Categories{n.Finance},
				Content:    "Test content 4",
			},
			{
				ID:         "5",
				Categories: []n.Categories{n.Sports},
				Content:    "Test content 5",
			},
			{
				ID:         "6",
				Categories: []n.Categories{n.Finance, n.Sports, n.Movies},
				Content:    "Test content 6",
			},
		}
	*/

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

	http.HandleFunc("/message", handlers.Handler)
	http.ListenAndServe(":8080", handlers.LoggingMiddleware(http.DefaultServeMux))
}
