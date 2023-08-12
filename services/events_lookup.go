package services

import (
	"cc/models"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

func (es *EventsService) Lookup(req *models.EventsLookupRequest) (*models.EventsLookupResponse, error) {
	switch req.OrderBy {
	case models.OrderByNearby:
		return nil, errors.New("nearby is not supported in this version")

	case models.OrderByPopularity:
		return es.lookupByPopularity(req)

	default:
		return nil, errors.New("wrong order clause")
	}
}

func (es *EventsService) lookupByPopularity(req *models.EventsLookupRequest) (*models.EventsLookupResponse, error) {
	query := `select e.id, e.title, e.cmi, e.coords, e.placedescription, e.eventdescription, e.announced, e.updated from events e
		inner join event_tags et on e.id = et.event_id
		                                and et.tag = any ($4)
		  and e.id in (select event_id from schedules where day >= $1 and day <= $2) 
		  and e.deleted is null group by e.id, e.title, e.cmi having count(distinct et.tag) = $3 order by e.cmi`

	rows, err := es.p.Query(context.Background(), query, req.DateMin, req.DateMax, len(req.Interests), req.Interests)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	resp := &models.EventsLookupResponse{Events: make([]*models.Event, 0)}
	resp.Timestamp = time.Now()

	resp.Events, err = es.fetchEventRows(rows)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// pgPoint represents a point that may be null.
type pgPoint struct {
	X, Y float32 // Coordinates of point
}

func (p *pgPoint) ScanPoint(v pgtype.Point) error {
	*p = pgPoint{
		X: float32(v.P.X),
		Y: float32(v.P.Y),
	}
	return nil
}

func (p *pgPoint) PointValue() (pgtype.Point, error) {
	return pgtype.Point{
		P:     pgtype.Vec2{X: float64(p.X), Y: float64(p.Y)},
		Valid: true,
	}, nil
}

func (es *EventsService) fetchEventRows(rows pgx.Rows) ([]*models.Event, error) {
	resp := make([]*models.Event, 0)

	for rows.Next() {
		event := &models.Event{}

		loc := &models.Location{}
		pnt := &pgPoint{}

		err := rows.Scan(&event.ID, &event.Title, &event.CountMeIn, &pnt, &loc.Place, &event.Description, &event.Announced, &event.Updated)
		if err != nil {
			return nil, err
		}

		event.Schedule, err = es.getSchedule(event.ID)
		if err != nil {
			return nil, err
		}

		event.Interests, err = es.getTags(event.ID)
		if err != nil {
			return nil, err
		}

		loc.Coords = [2]float32{pnt.X, pnt.Y}
		event.Loc = loc
		resp = append(resp, event)
	}

	return resp, nil
}

func (es *EventsService) getSchedule(eventID string) ([]models.DaySchedule, error) {
	schedule := "select day, opensAt, closesAt from schedules where event_id = $1"

	rows, err := es.p.Query(context.Background(), schedule, eventID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []models.DaySchedule

	for rows.Next() {
		var date, opensAt, closesAt time.Time

		var ds models.DaySchedule

		err = rows.Scan(&date, &opensAt, &closesAt)
		if err != nil {
			return nil, err
		}

		ds[0] = date.Format("2006-01-02")
		ds[1] = date.Format("15:04")
		ds[2] = date.Format("15:04")

		result = append(result, ds)
	}

	return result, nil
}

func (es *EventsService) getTags(eventID string) ([]string, error) {
	tags := "select tag from event_tags where event_id = $1"

	rows, err := es.p.Query(context.Background(), tags, eventID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var result []string

	for rows.Next() {
		var s string
		err = rows.Scan(&s)

		if err != nil {
			return nil, err
		}

		result = append(result, s)
	}

	return result, nil
}
