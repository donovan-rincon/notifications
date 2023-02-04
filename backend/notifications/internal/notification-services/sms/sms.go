package sms

import (
	"log"
	"time"

	n "github.com/gila-software/backend/notifications"
)

type NotificationService struct {
	Storage n.StorageService
}

func (s *NotificationService) Notify(msg n.Message) error {
	users, _ := s.Storage.UsersByChannel(n.SMS)
	for _, u := range users {
		log.Printf("Notification type: %s, Sent to user: %s. Timestamp: %s, Message content: %s", n.SMS.String(), string(u.Name), time.Now(), msg.Content)
	}
	return nil
}
