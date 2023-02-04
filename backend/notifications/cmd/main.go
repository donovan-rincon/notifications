package main

import (
	n "github.com/gila-software/backend/notifications"
	"github.com/gila-software/backend/notifications/internal/notification-services/email"
	pushnotifications "github.com/gila-software/backend/notifications/internal/notification-services/push-notifications"
	"github.com/gila-software/backend/notifications/internal/notification-services/sms"
	"github.com/gila-software/backend/notifications/internal/storage/localstorage"
)

func main() {

	ls := localstorage.NewLocalStorage()

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
			Categories: []n.Categories{n.Finance, n.Sports, n.Movies},
			Content:    "Test content 3",
		},
		{
			ID:         "4",
			Categories: []n.Categories{n.Finance, n.Movies},
			Content:    "Test content 4",
		},
		{
			ID:         "5",
			Categories: []n.Categories{n.Finance, n.Sports},
			Content:    "Test content 5",
		},
		{
			ID:         "6",
			Categories: []n.Categories{n.Finance, n.Sports, n.Movies},
			Content:    "Test content 6",
		},
	}

	notifiers := []n.NotificationService{
		&sms.NotificationService{
			Storage: &ls,
		},
		&pushnotifications.NotificationService{
			Storage: &ls,
		},
		&email.NotificationService{
			Storage: &ls,
		},
	}

	for _, m := range msgs {
		ls.StoreMessage(m)
		for _, n := range notifiers {
			n.Notify(m)
		}
	}
}
