package main

import (
	"AlertNotification/config"
	"AlertNotification/model"
	"AlertNotification/service"
	"fmt"
	"strings"
)

func main() {
	config.InitAlertConfig()
	config.InitNotification()

	fransisko := model.NewUser("1", "Fransisko")
	abs := model.NewUser("2", "Abs")

	smsAlert := config.GetAlert(config.SMS)
	emailAlert := config.GetAlert(config.Email)
	phoneAlert := config.GetAlert(config.Phone)

	smsAlert.Register(fransisko)
	smsAlert.Register(abs)
	emailAlert.Register(fransisko)
	phoneAlert.Register(fransisko)

	data := make([]string, 0)
	data = append(data, "2019-01-07 14:52:33 Warning data")
	data = append(data, "2019-01-07 14:52:34 Critical data")
	data = append(data, "2019-01-07 14:52:35 Info data")
	data = append(data, "2019-01-07 14:52:36 Critical data")
	data = append(data, "2019-01-07 14:52:37 Critical data")
	data = append(data, "2019-01-07 14:52:38 Critical data")
	data = append(data, "2019-01-07 14:52:39 Critical data")
	data = append(data, "2019-01-07 14:52:40 Critical data")
	data = append(data, "2019-01-07 14:52:41 Warning data")
	data = append(data, "2019-01-07 14:52:42 Critical data")
	data = append(data, "2019-01-07 14:52:43 Warning data")
	data = append(data, "2019-01-07 14:52:44 Critical data")
	data = append(data, "2019-01-07 14:52:45 Critical data")
	data = append(data, "2019-01-07 14:52:46 Critical data")
	data = append(data, "2019-01-07 14:52:47 Critical data")
	data = append(data, "2019-01-07 14:52:48 Critical data")
	data = append(data, "2019-01-07 14:52:49 Critical data")
	data = append(data, "2019-01-07 14:52:50 Critical data")
	data = append(data, "2019-01-07 14:52:51 Critical data")

	criticalNotificationConfig := config.NewLogLevelConfig(smsAlert, emailAlert, phoneAlert)

	blockerNotificationConfig := config.NewLogLevelConfig()
	blockerNotificationConfig.AddNotification(smsAlert)
	blockerNotificationConfig.AddNotification(emailAlert)
	blockerNotificationConfig.AddNotification(phoneAlert)

	warningNotificationConfig := config.NewLogLevelConfig()
	warningNotificationConfig.AddNotification(smsAlert)

	criticalService := service.NewAlertService(model.Critical, criticalNotificationConfig.GetListNotification())
	warningService := service.NewAlertService(model.Warning, warningNotificationConfig.GetListNotification())
	blockerService := service.NewAlertService(model.Blocker, blockerNotificationConfig.GetListNotification())

	for _, val := range data {
		split := strings.Split(val, " ")

		switch model.LogLevel(split[2]) {
		case model.Critical:
			criticalService.ProcessData(val)
		case model.Warning:
			warningService.ProcessData(val)
		case model.Blocker:
			blockerService.ProcessData(val)
		default:
			fmt.Println(fmt.Sprintf("'%v' service not found", split[2]))
		}
	}

}
