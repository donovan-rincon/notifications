package notificationservices

import (
	"log"
	"os"

	"github.com/gila-software/backend/notifications"
)

type LogMessage struct {
	NotificationType string                `json:"notification-type"`
	User             notifications.User    `json:"user"`
	Timestamp        string                `json:"timestamp"`
	Message          notifications.Message `json:"message"`
}

func Init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.SetFlags(1)
}
