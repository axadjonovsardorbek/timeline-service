package postgres

import (
	"database/sql"
	"fmt"

	"timeline/config"
	"timeline/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db         *sql.DB
	EventS     storage.CustomEventsI
	MilestoneS storage.MilestonesI
}

func NewPostgresStorage(config config.Config) (*Storage, error) {
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD, config.DB_PORT)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	event := NewCustomEventsRepo(db)
	milestone := NewMilestonesRepo(db)

	return &Storage{
		Db:     db,
		EventS: event,
		MilestoneS: milestone,
	}, nil
}
