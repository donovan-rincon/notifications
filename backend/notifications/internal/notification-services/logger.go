package notificationservices

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gila-software/backend/notifications"
)

func Init() {
	log.Default()
}

func LogMessage(msg notifications.LogggerMessage) {
	file, err := os.OpenFile("messages_logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	line, _ := json.Marshal(msg)
	_, err = file.Write(line)
	if err != nil {
		log.Println("Could not write text to messages_logs.txt")
	}
	file.WriteString("\n")
}
