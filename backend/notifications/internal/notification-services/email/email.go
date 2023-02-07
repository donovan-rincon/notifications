package email

import (
	"log"
	"time"

	n "github.com/gila-software/backend/notifications"
	notificationservices "github.com/gila-software/backend/notifications/internal/notification-services"
)

type NotificationService struct {
	Storage n.StorageService
}

func (s *NotificationService) Notify(msg n.Message) error {
	users, _ := s.Storage.UsersByChannel(n.Email)
	for _, u := range users {
		for _, c := range msg.Categories {
			if u.Subscribed[c.String()] {
				logMsg := &n.LogggerMessage{
					NotificationType: n.Email.String(),
					User:             u,
					Timestamp:        time.Now().Format("2006-01-02 15:04:05"),
					Message:          msg,
				}
				notificationservices.LogMessage(*logMsg)
				log.Printf("Message logged: %s %v\n", logMsg.NotificationType, logMsg.Message)
			}
		}
	}

	return nil
}
