package config

import "AlertNotification/notification"

type LogLevelConfig struct {
	listNotification []*notification.AlertNotification
}

//func NewLogLevelConfig() *LogLevelConfig {
//	newLog := new(LogLevelConfig)
//	newLog.listNotification = make([]*notification.AlertNotification, 0)
//	return newLog
//}

func NewLogLevelConfig(alerts ...*notification.AlertNotification) *LogLevelConfig {
	newLog := new(LogLevelConfig)
	newLog.listNotification = make([]*notification.AlertNotification, 0)
	newLog.listNotification = append(newLog.listNotification, alerts...)
	return newLog
}

func (ll *LogLevelConfig) GetListNotification() []*notification.AlertNotification {
	return ll.listNotification
}

func (ll *LogLevelConfig) AddNotification(alert *notification.AlertNotification) {
	ll.listNotification = append(ll.listNotification, alert)
}

func (ll *LogLevelConfig) DeleteNotification(alert *notification.AlertNotification) {
	for i, val := range ll.listNotification {
		if val == alert {
			ll.listNotification = append(ll.listNotification[:i], ll.listNotification[i+1])
		}
	}
}
