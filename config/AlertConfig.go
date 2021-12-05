package config

import (
	"AlertNotification/model"
	"fmt"
)

type AlertConfig struct {
	Type      model.LogLevel
	Frequency int
	Duration  int
	WaitTime  int
}

var alertConfigs map[model.LogLevel]*AlertConfig

func NewAlertConfig(alertType model.LogLevel, frequency, duration, waitTime int) *AlertConfig {
	return &AlertConfig{
		Type:      alertType,
		Frequency: frequency,
		Duration:  duration,
		WaitTime:  waitTime,
	}
}

func InitAlertConfig() {
	alertConfigs = make(map[model.LogLevel]*AlertConfig)
	alertConfigs[model.Critical] = NewAlertConfig(model.Critical, 5, 100, 3)
	alertConfigs[model.Warning] = NewAlertConfig(model.Warning, 20, 100, 100)
	alertConfigs[model.Info] = NewAlertConfig(model.Info, 30, 100, 100)
	alertConfigs[model.Blocker] = NewAlertConfig(model.Blocker, 30, 100, 100)
}

func GetAlertConfig(level model.LogLevel) (*AlertConfig, error) {
	if a, ok := alertConfigs[level]; ok {
		return a, nil
	}

	return nil, fmt.Errorf("config not found")
}
