package notification

type Observer interface {
	ReceiveNotification(message string)
	GetId() string
}
