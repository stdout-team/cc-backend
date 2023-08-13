package models

import "time"

type CountMeInResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Event     string    `json:"event"`
	Count     int       `json:"countMeIn"`
}
