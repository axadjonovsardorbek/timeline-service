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

type MilestonesRepo struct {
	db *sql.DB
}

func NewMilestonesRepo(db *sql.DB) *MilestonesRepo {
	return &MilestonesRepo{db: db}
}

func (m *MilestonesRepo) Create(req *mp.MilestonesCreateReq) (*mp.Void, error) {
	id := uuid.New().String()

	query := `
	INSERT INTO milestones (
		id,
		user_id,
		title,
		date,
		category
	) VALUES ($1, $2, $3, $4, $5)
	`

	_, err := m.db.Exec(query, id, req.UserId, req.Title, req.Date, req.Category)
	if err != nil {
		log.Println("Error while creating milestone: ", err)
		return nil, err
	}

	log.Println("Successfully created milestone")
	return nil, nil
}

func (m *MilestonesRepo) GetById(id *mp.ById) (*mp.MilestonesGetByIdRes, error) {
	milestone := mp.MilestonesGetByIdRes{
		Milestone: &mp.MilestonesRes{},
	}

	query := `
	SELECT 
		id,
		user_id,
		title,
		date,
		category
	FROM 
		milestones
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	row := m.db.QueryRow(query, id.Id)
	err := row.Scan(
		&milestone.Milestone.Id,
		&milestone.Milestone.UserId,
		&milestone.Milestone.Title,
		&milestone.Milestone.Date,
		&milestone.Milestone.Category,
	)

	if err != nil {
		log.Println("Error while getting milestone by id: ", err)
		return nil, err
	}

	log.Println("Successfully got milestone")
	return &milestone, nil
}

func (m *MilestonesRepo) GetAll(req *mp.MilestonesGetAllReq) (*mp.MilestonesGetAllRes, error) {
	milestones := mp.MilestonesGetAllRes{}

	query := `
	SELECT 
		id,
		user_id,
		title,
		date,
		category
	FROM 
		milestones
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

	rows, err := m.db.Query(query, args...)
	if err != nil {
		log.Println("Error while retrieving milestones: ", err)
		return nil, err
	}

	for rows.Next() {
		milestone := mp.MilestonesRes{}

		err := rows.Scan(
			&milestone.Id,
			&milestone.UserId,
			&milestone.Title,
			&milestone.Date,
			&milestone.Category,
		)

		if err != nil {
			log.Println("Error while scanning all milestones: ", err)
			return nil, err
		}

		milestones.Milestones = append(milestones.Milestones, &milestone)
	}

	milestones.Count = int32(len(milestones.Milestones))
	log.Println("Successfully fetched all milestones")
	return &milestones, nil
}

func (m *MilestonesRepo) Update(req *mp.MilestonesUpdateReq) (*mp.Void, error) {

	query := `
	UPDATE
		milestones
	SET 
	`

	var conditions []string
	var args []interface{}

	if req.Title != "" && req.Title != "string" {
		conditions = append(conditions, " title = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Title)
	}
	if req.Date != "" {
		conditions = append(conditions, " date = $"+strconv.Itoa(len(args)+1))
		args = append(args, req.Date)
	}

	if len(conditions) == 0 {
		return nil, errors.New("nothing to update")
	}

	conditions = append(conditions, " updated_at = now()")
	query += strings.Join(conditions, ", ")
	query += " WHERE id = $" + strconv.Itoa(len(args)+1) + " AND deleted_at = 0 "

	args = append(args, req.Id)

	_, err := m.db.Exec(query, args...)
	if err != nil {
		log.Println("Error while updating milestone: ", err)
		return nil, err
	}

	log.Println("Successfully updated milestone")
	return nil, nil
}

func (m *MilestonesRepo) Delete(id *mp.ById) (*mp.Void, error) {

	query := `
	UPDATE 
		milestones
	SET 
		deleted_at = EXTRACT(EPOCH FROM NOW())
	WHERE 
		id = $1
	AND 
		deleted_at = 0
	`

	res, err := m.db.Exec(query, id.Id)
	if err != nil {
		log.Println("Error while deleting milestone: ", err)
		return nil, err
	}

	if r, err := res.RowsAffected(); r == 0 {
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("milestone with id %s not found", id.Id)
	}

	log.Println("Successfully deleted milestone")
	return nil, nil
}
