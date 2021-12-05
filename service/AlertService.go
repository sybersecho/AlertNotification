package service

import (
	"AlertNotification/config"
	"AlertNotification/model"
	"AlertNotification/notification"
	"fmt"
	"strings"
	"time"
)

type AlertService struct {
	alertType     model.LogLevel
	queue         model.Queue
	freezeUntil   time.Time
	notifications []*notification.AlertNotification
}

func NewAlertService(alertType model.LogLevel, notifications []*notification.AlertNotification) *AlertService {
	return &AlertService{
		alertType:     alertType,
		queue:         model.Queue{Data: make([]time.Time, 0)},
		freezeUntil:   time.Time{},
		notifications: notifications,
	}
}

func (as *AlertService) ProcessData(data string) {
	alertConfig, err := config.GetAlertConfig(as.alertType)
	if err != nil {
		return
	}

	split := strings.Split(data, " ")
	dateTimeStr := fmt.Sprintf("%v %v", split[0], split[1])
	dateTime, _ := time.ParseInLocation("2006-01-02 15:04:05", dateTimeStr, time.Local)

	if !as.freezeUntil.IsZero() && dateTime.Before(as.freezeUntil) {
		fmt.Println(fmt.Printf("%v: freeze until %v", dateTimeStr, as.freezeUntil.Format("2006-01-02 15:04:05")))
		return
	} else {
		as.freezeUntil = time.Time{}
	}

	as.queue.Enqueue(dateTime)

	first, _ := as.queue.Front()
	last, _ := as.queue.Last()
	if (alertConfig.Frequency == as.queue.Size()) && (int(last.Sub(*first).Seconds()) < alertConfig.Duration) {
		as.freezeUntil = last.Add(time.Duration(alertConfig.WaitTime) * time.Second)
		as.queue.Clear()
		fmt.Println(fmt.Sprintf("send notification of %v alert", as.alertType))
		for _, val := range as.notifications {
			val.NotifyAll()
		}
	}
}
