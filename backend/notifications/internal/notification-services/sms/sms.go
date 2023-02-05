package sms

import (
	"encoding/json"
	"log"
	"time"

	n "github.com/gila-software/backend/notifications"
	ns "github.com/gila-software/backend/notifications/internal/notification-services"
)

type NotificationService struct {
	Storage n.StorageService
}

func (s *NotificationService) Notify(msg n.Message) error {
	sendMessage := false
	for _, c := range msg.Categories {
		if c == n.Category(n.SMS) {
			sendMessage = true
		}
	}
	if sendMessage {
		users, _ := s.Storage.UsersByChannel(n.SMS)
		for _, u := range users {
			logMsg := &ns.LogMessage{
				NotificationType: n.SMS.String(),
				User:             u,
				Timestamp:        time.Now().Format("2006-01-02 15:04:05"),
				Message:          msg,
			}
			out, _ := json.Marshal(logMsg)
			log.Println(string(out))
		}
	}
	return nil
}
