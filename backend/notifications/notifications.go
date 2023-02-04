package notifications

type User struct {
	ID          string
	Name        string
	Email       string
	PhoneNumber int
	Subscribed  map[string]bool
	Channels    map[string]bool
}

type Categories int64

const (
	Finance Categories = iota
	Movies
	Sports
)

func (c Categories) String() string {
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

type Channels int64

const (
	Email Channels = iota
	PushNotifications
	SMS
)

func (c Channels) String() string {
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
	ID         string
	Categories []Categories
	Content    string
}

type NotificationService interface {
	Notify(msg Message) error
}

type StorageService interface {
	Users() ([]User, error)
	UsersByCategory(categories Categories) ([]User, error)
	UsersByChannel(channel Channels) ([]User, error)
	StoreMessage(msg Message) error
	Messages() ([]Message, error)
}
