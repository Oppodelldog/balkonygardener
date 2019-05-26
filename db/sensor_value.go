package db

import "time"

type SensorValues []SensorValue
type SensorValue struct {
	T time.Time
	V interface{}
}
