package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/google/uuid"
	mp "timeline/genproto"
)

type CustomEventsRepo struct {
	db *sql.DB
}

func NewCustomEventsRepo(db *sql.DB) *CustomEventsRepo {
	return &CustomEventsRepo{db: db}
}

func (c *CustomEventsRepo) Create(req *mp.CustomEventsCreateReq) (*mp.Void, error) {
	id := uuid.New().String()

	query := `
	INSERT INTO custom_events (
		id,
		user_id,
		title,
		description,
		date,
		category
	) VALUES ($1, $2, $3, $4, $5, $6)
	`

	_, err := c.db.Exec(query, id, req.UserId, req.Title, req.Description, req.Date, req.Category)
	if err != nil {
		log.Println("Error while creating custom event: ", err)
		return nil, err
	}

	log.Println("Successfully created custom event")
	return nil, nil
}

func (c *CustomEventsRepo) GetById(id *mp.ById) (*mp.CustomEventsGetByIdRes, error) {
	event := mp.CustomEventsGetByIdRes{
		Event: &mp.CustomEventsRes{},
	}

	query := `
	SELECT 
		id,
		user_id,
		title,
		description,
		date,
		category
	FROM 
		custom_events
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	row := c.db.QueryRow(query, id.Id)
	err := row.Scan(
		&event.Event.Id,
		&event.Event.UserId,
		&event.Event.Title,
		&event.Event.Description,
		&event.Event.Date,
		&event.Event.Category,
	)

	if err != nil {
		log.Println("Error while getting custom event by id: ", err)
		return nil, err
	}

	log.Println("Successfully got custom event")
	return &event, nil
}

func (c *CustomEventsRepo) GetAll(req *mp.CustomEventsGetAllReq) (*mp.CustomEventsGetAllRes, error) {
	events := mp.CustomEventsGetAllRes{}

	query := `
	SELECT 
		id,
		user_id,
		title,
		description,
		date,
		category
	FROM 
		custom_events
	WHERE 
		deleted_at = 0
	`

	var args []interface{}
	var conditions []string

	if req.UserId != "" && req.UserId != "string" {
		conditions = append(conditions, " user_id = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.UserId)
	}

	if len(conditions) > 0 {
		query += " AND " + strings.Join(conditions, " AND ")
	}

	var limit int32 = 10
	var offset int32 = req.Filter.Page * limit

	args = append(args, limit, offset)
	query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", len(args)-1, len(args))

	rows, err := c.db.Query(query, args...)
	if err != nil {
		log.Println("Error while retrieving custom events: ", err)
		return nil, err
	}

	for rows.Next() {
		event := mp.CustomEventsRes{}

		err := rows.Scan(
			&event.Id,
			&event.UserId,
			&event.Title,
			&event.Description,
			&event.Date,
			&event.Category,
		)

		if err != nil {
			log.Println("Error while scanning all custom events: ", err)
			return nil, err
		}

		events.Events = append(events.Events, &event)
	}

	events.Count = int32(len(events.Events))
	log.Println("Successfully fetched all custom events")
	return &events, nil
}

func (c *CustomEventsRepo) Update(req *mp.CustomEventsUpdateReq) (*mp.Void, error) {

	query := `
	UPDATE
		custom_events
	SET 
	`

	var conditions []string
	var args []interface{}

	if req.Title != "" && req.Title != "string" {
		conditions = append(conditions, " title = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Title)
	}
	if req.Description != "" && req.Description != "string" {
		conditions = append(conditions, " description = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Description)
	}

	if len(conditions) == 0 {
		return nil, errors.New("nothing to update")
	}

	conditions = append(conditions, " updated_at = now()")
	query += strings.Join(conditions, ", ")
	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0 "

	args = append(args, req.Id)

	_, err := c.db.Exec(query, args...)
	if err != nil {
		log.Println("Error while updating custom event: ", err)
		return nil, err
	}

	log.Println("Successfully updated custom event")
	return nil, nil
}

func (c *CustomEventsRepo) Delete(id *mp.ById) (*mp.Void, error) {

	query := `
	UPDATE 
		custom_events
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := c.db.Exec(query, id.Id)
	if err != nil {
		log.Println("Error while deleting custom event: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("custom event with id %s not found", id.Id)
	}

	log.Println("Successfully deleted custom event")
	return nil, nil
}
