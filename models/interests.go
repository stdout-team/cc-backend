package models

import "time"

type GetInterestsResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Interests []string  `json:"interests"`
}
