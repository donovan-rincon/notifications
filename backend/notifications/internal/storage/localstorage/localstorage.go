package localstorage

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	n "github.com/gila-software/backend/notifications"
)

type StorageService struct {
	DB map[string]interface{}
}

var mockUsers []n.User = []n.User{
	{
		ID:          "33887837-f1d6-4e25-8a2a-2b1fea717493",
		Name:        "James Hetfield",
		Email:       "james@metallica.com",
		PhoneNumber: 123456789,
		Subscribed:  map[string]bool{n.Finance.String(): true, n.Movies.String(): false, n.Sports.String(): true},
		Channels:    map[string]bool{n.SMS.String(): false, n.Email.String(): true, n.PushNotifications.String(): false},
	},
	{
		ID:          "a2f8cc67-5ea2-426f-b115-2b59dcc4c2c2",
		Name:        "Kirk Hammet",
		Email:       "kirk@metallica.com",
		PhoneNumber: 123456789,
		Subscribed:  map[string]bool{n.Finance.String(): false, n.Movies.String(): true, n.Sports.String(): false},
		Channels:    map[string]bool{n.SMS.String(): true, n.Email.String(): true, n.PushNotifications.String(): false},
	},
	{
		ID:          "2c666e58-8286-44c0-9eef-410e6a9e1b31",
		Name:        "Lars Ulrich",
		Email:       "lars@metallica.com",
		PhoneNumber: 123456789,
		Subscribed:  map[string]bool{n.Finance.String(): true, n.Movies.String(): true, n.Sports.String(): false},
		Channels:    map[string]bool{n.SMS.String(): false, n.Email.String(): true, n.PushNotifications.String(): true},
	},
}

var categories map[string]n.Category = map[string]n.Category{n.Finance.String(): n.Finance, n.Movies.String(): n.Movies, n.Sports.String(): n.Sports}

func NewLocalStorage() StorageService {
	return StorageService{
		DB: map[string]interface{}{
			"users":      mockUsers,
			"messages":   []n.Message{},
			"categories": categories,
		},
	}
}

func (s *StorageService) Users() ([]n.User, error) {
	return s.DB["users"].([]n.User), nil
}

func (s *StorageService) UsersByCategory(category n.Category) ([]n.User, error) {
	res := []n.User{}
	for _, u := range s.DB["users"].([]n.User) {
		if u.Subscribed[category.String()] {
			res = append(res, u)
		}
	}
	return res, nil
}

func (s *StorageService) UsersByChannel(channel n.Channel) ([]n.User, error) {
	res := []n.User{}
	for _, u := range s.DB["users"].([]n.User) {
		if u.Channels[channel.String()] {
			res = append(res, u)
		}
	}
	return res, nil
}

func (s *StorageService) StoreMessage(msg n.Message) error {
	s.DB["messages"] = append(s.DB["messages"].([]n.Message), msg)
	return nil
}

func (s *StorageService) Messages() ([]n.Message, error) {
	return s.DB["messages"].([]n.Message), nil
}

func (s *StorageService) Categories() (map[string]n.Category, error) {
	return s.DB["categories"].(map[string]n.Category), nil
}

func (s *StorageService) MessagesHistory() []n.LogggerMessage {
	response := []n.LogggerMessage{}
	readFile, err := os.Open("messages_logs.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		var msg n.LogggerMessage
		err := json.Unmarshal(fileScanner.Bytes(), &msg)
		if err != nil {
			log.Println(err.Error())
		}
		response = append(response, msg)
	}

	for i := 0; i < len(response)/2; i++ {
		j := len(response) - i - 1
		response[i], response[j] = response[j], response[i]
	}

	return response
}
