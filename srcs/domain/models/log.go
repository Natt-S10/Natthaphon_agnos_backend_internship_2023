package models

import (
	"time"
)

type LogRecord struct {
	Timestamp    time.Time
	Route        string
	Status       int
	InitPassword string
	NumOfSteps   int
	Error        int
}
