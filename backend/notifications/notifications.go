package notifications

var Storage StorageService
var Notifiers []NotificationService

type LogggerMessage struct {
	NotificationType string  `json:"notification-type"`
	User             User    `json:"user"`
	Timestamp        string  `json:"timestamp"`
	Message          Message `json:"message"`
}

type User struct {
	ID          string          `json:"id"`
	Name        string          `json:"name"`
	Email       string          `json:"email"`
	PhoneNumber int             `json:"phone-number"`
	Subscribed  map[string]bool `json:"subscribed"`
	Channels    map[string]bool `json:"channels"`
}

type Category int64

const (
	Finance Category = iota
	Movies
	Sports
)

func (c Category) String() string {
	switch c {
	case Finance:
		return "finance"
	case Movies:
		return "movies"
	case Sports:
		return "sports"
	}
	return "unknown"
}

type Channel int64

const (
	Email Channel = iota
	PushNotifications
	SMS
)

func (c Channel) String() string {
	switch c {
	case Email:
		return "email"
	case PushNotifications:
		return "push-notifications"
	case SMS:
		return "sms"
	}
	return "unknown"
}

type Message struct {
	ID         string     `json:"id"`
	Categories []Category `json:"categories"`
	Content    string     `json:"content"`
}

type NotificationService interface {
	Notify(msg Message) error
}

type StorageService interface {
	Users() ([]User, error)
	UsersByCategory(categories Category) ([]User, error)
	UsersByChannel(channel Channel) ([]User, error)
	StoreMessage(msg Message) error
	Messages() ([]Message, error)
	Categories() (map[string]Category, error)
	MessagesHistory() []LogggerMessage
}
