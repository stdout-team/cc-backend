package models

import (
	"strings"
	"time"
)

const (
	OrderByPopularity = iota
	OrderByNearby
)

type EventsLookupAPIRequest struct {
	Interests string `binding:"required" form:"interests"`
	DateMin   string `binding:"required" form:"minDate"`
	DateMax   string `binding:"required" form:"maxDate"`
	OrderBy   string `binding:"required" form:"orderBy"`
}

type EventsLookupRequest struct {
	Interests []string
	DateMin   time.Time
	DateMax   time.Time
	OrderBy   int
}

func NewEventsLookupRequest(apiReq *EventsLookupAPIRequest) (*EventsLookupRequest, error) {
	req := &EventsLookupRequest{}
	req.Interests = strings.Split(apiReq.Interests, ",")

	var err error

	req.DateMin, err = time.Parse("2006-01-02", apiReq.DateMin)
	if err != nil {
		return nil, err
	}

	req.DateMax, err = time.Parse("2006-01-02", apiReq.DateMax)
	if err != nil {
		return nil, err
	}

	switch apiReq.OrderBy {
	case "nearby":
		req.OrderBy = OrderByNearby

	case "popularity", "pop":
		fallthrough

	default:
		req.OrderBy = OrderByPopularity
	}

	return req, nil
}

type EventsLookupResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Events    []*Event  `json:"events"`
}
