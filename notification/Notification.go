package notification

type Notification interface {
	Register(user Observer)
	Deregister(user Observer)
	NotifyAll()
}

type AlertNotification struct {
	userList []Observer
	name     string
}

func NewAlertNotification(name string) *AlertNotification {
	newAlert := new(AlertNotification)
	newAlert.userList = make([]Observer, 0)
	newAlert.name = name

	return newAlert
}

func (an *AlertNotification) Register(user Observer) {
	an.userList = append(an.userList, user)
}

func (an *AlertNotification) Deregister(user Observer) {
	an.userList = removeFromList(an.userList, user)
}

func removeFromList(list []Observer, userToRemove Observer) []Observer {
	listLength := len(list)
	for i, val := range list {
		if val.GetId() == userToRemove.GetId() {
			list[listLength-1], list[i] = list[i], list[listLength-1]
		}
	}

	return list
}

func (an *AlertNotification) NotifyAll() {
	for _, user := range an.userList {
		user.ReceiveNotification(an.name)
	}
}
