package services

import (
	"cc/models"
	"context"
	"github.com/google/uuid"
	"time"
)

func (es *EventsService) AddEvent(input *models.EventInput) (*models.Event, error) {
	coords, err := ResolveCoordinatesByStreetAddress(input.Location)
	if err != nil {
		panic(err)
	}

	tx, err := es.p.Begin(context.Background())
	if err != nil {
		return nil, err
	}

	eventID := uuid.Must(uuid.NewRandom()).String()
	_, err = tx.Exec(context.Background(), "insert into events (id, title, placedescription, eventdescription, coords, announced, updated, cmi, created) values ($1, $2, $3, $4, $5, $6, $7, $8, $9)",
		eventID, input.Title, input.Location, input.Description, &pgPoint{
			X: float32(coords[0]), Y: float32(coords[1]),
		}, time.Now(), time.Now(), 0, time.Now())

	if err != nil {
		return nil, err
	}

	for _, tag := range input.Interests {
		_, err := tx.Exec(context.Background(), "insert into event_tags (event_id, tag) values ($1, $2)", eventID, tag)
		if err != nil {
			return nil, err
		}
	}

	for _, day := range input.Schedule {
		_, err := tx.Exec(context.Background(), "insert into schedules (event_id, day, opensat, closesat) values ($1, $2, $3, $4)",
			eventID, day[0], day[1], day[2])

		if err != nil {
			return nil, err
		}
	}

	tx.Commit(context.Background())

	return es.getByID(eventID)
}
