package services

import (
	"cc/models"
	"context"
	"time"
)

func (es *EventsService) GetInterests() (*models.GetInterestsResponse, error) {
	rows, err := es.p.Query(context.Background(), "select distinct tag from event_tags")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []string

	for rows.Next() {
		var tag string

		err := rows.Scan(&tag)
		if err != nil {
			return nil, err
		}

		result = append(result, tag)
	}

	return &models.GetInterestsResponse{
		Timestamp: time.Now(),
		Interests: result,
	}, nil
}
