package models

import "time"

type Event struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Schedule    []DaySchedule `json:"schedule"`
	Loc         *Location     `json:"location"`
	Announced   time.Time     `json:"announced"`
	Updated     time.Time     `json:"updated"`
	Interests   []string      `json:"interests"`
	Description string        `json:"description"`
	Deleted     time.Time     `json:"-"`
	CountMeIn   int           `json:"countMeIn"`
}

type (
	DaySchedule = [3]string
	Coordinates = [2]float32
	Album       = []string
)

type Location struct {
	Place  string      `json:"place"`
	Coords Coordinates `json:"coords"`
}

type EventMedia struct {
	Preview string `json:"preview"`
	Alb     Album  `json:"album"`
}
