package services

import (
	"cc/models"
	"context"
	"time"
)

func (es *EventsService) CountMeIn(eventID string) (*models.CountMeInResponse, error) {
	row := es.p.QueryRow(context.Background(), "update events set cmi = cmi + 1 where id = $1 returning cmi", eventID)

	var newValue int
	err := row.Scan(&newValue)

	return &models.CountMeInResponse{
		Timestamp: time.Now(),
		Event:     eventID,
		Count:     newValue,
	}, err
}
