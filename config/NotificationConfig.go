package config

import "AlertNotification/notification"

type Notification string

const (
	SMS   Notification = "sms"
	Phone              = "phone"
	Email              = "Email"
)

var listNotification map[Notification]*notification.AlertNotification

func InitNotification() {
	listNotification = make(map[Notification]*notification.AlertNotification)
	listNotification[SMS] = notification.NewAlertNotification(string(SMS))
	listNotification[Phone] = notification.NewAlertNotification(Phone)
	listNotification[Email] = notification.NewAlertNotification(Email)
}

func GetAlert(notification Notification) *notification.AlertNotification {
	return listNotification[notification]
}
