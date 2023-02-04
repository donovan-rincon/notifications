package localstorage

import (
	n "github.com/gila-software/backend/notifications"
)

type StorageService struct {
	DB map[string]interface{}
}

func (s *StorageService) Users() ([]n.User, error) {

}

func (s *StorageService) UsersByCategory(category n.Categories) ([]n.User, error) {

}

func (s *StorageService) StoreMessage(msg n.Message) error {

}

func (s *StorageService) Messages() ([]n.Message, error) {

}
