package services

import "github.com/jackc/pgx/v5/pgxpool"

type EventsService struct {
	p *pgxpool.Pool
}

func NewEventsService(p *pgxpool.Pool) *EventsService {
	return &EventsService{p: p}
}
