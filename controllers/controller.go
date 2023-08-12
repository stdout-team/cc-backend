package controllers

import "github.com/jackc/pgx/v5/pgxpool"

type Controller struct {
	connPool *pgxpool.Pool
}

func NewController(p *pgxpool.Pool) *Controller {
	return &Controller{connPool: p}
}
