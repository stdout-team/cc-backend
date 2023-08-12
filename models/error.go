package models

import "time"

type ErrorResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Errors    []string  `json:"errors"`
}
