package model

type LogLevel string

const (
	Info     LogLevel = "Info"
	Warning  LogLevel = "Warning"
	Critical LogLevel = "Critical"
	Blocker  LogLevel = "Blocker"
)
