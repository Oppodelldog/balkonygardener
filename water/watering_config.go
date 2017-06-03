package water

import "time"

type WateringConfig struct {
	Duration time.Duration
	PinName  string
	Comment  string
}