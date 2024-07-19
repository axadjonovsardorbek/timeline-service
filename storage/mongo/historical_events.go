package mongo

import (
	"context"
	"errors"
	"fmt"
	"log"
	mp "timeline/genproto"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type HistoricalEventsRepo struct {
	db *mongo.Database
}

func NewHistoricalEventsRepo(db *mongo.Database) *HistoricalEventsRepo {
	return &HistoricalEventsRepo{db: db}
}

func (c *HistoricalEventsRepo) Create(req *mp.HistoricalEventsRes) (*mp.Void, error) {
	collection := c.db.Collection("historical_event")
	id := uuid.New().String()

	req.Id = id

	res, err := collection.InsertOne(context.TODO(), req)

	fmt.Println(res)
	if err != nil {
		log.Println("Error while creating historical event: ", err)
		return nil, err
	}

	log.Println("Successfully created historical event")

	return nil, nil
}
func (c *HistoricalEventsRepo) GetById(req *mp.ById) (*mp.HistoricalEventsGetByIdRes, error) {
	collection := c.db.Collection("historical_event")
	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	var event mp.HistoricalEventsGetByIdRes
	err := collection.FindOne(context.TODO(), bson.M{"id": req.Id}).Decode(&event.Event)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("event not found")
	} else if err != nil {
		return nil, err
	}

	return &event, nil
}
func (c *HistoricalEventsRepo) GetAll(req *mp.HistoricalEventsGetAllReq) (*mp.HistoricalEventsGetAllRes, error) {
	collection := c.db.Collection("historical_event")
	mongoFilter := bson.M{}
	if req.UserId != "" {
		mongoFilter["userid"] = req.UserId
	}
	if req.Date != "" {
		mongoFilter = bson.M{"date": primitive.Regex{Pattern: req.Date, Options: "i"}}
	}
	if req.Category != "" {
		mongoFilter = bson.M{"category": primitive.Regex{Pattern: req.Category, Options: "i"}}
	}

	cursor, err := collection.Find(context.TODO(), mongoFilter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var events mp.HistoricalEventsGetAllRes

	var defaultLimit int32
	var check int32
	var offset int32
	var count int32

	defaultLimit = 10
	offset = req.Filter.Page * defaultLimit

	for cursor.Next(context.TODO()) {
		if check <= offset {
			check += 1
			continue
		}
		if count == defaultLimit {
			break
		}
		var reading mp.HistoricalEventsRes
		err := cursor.Decode(&reading)
		if err != nil {
			return nil, err
		}

		event := &mp.HistoricalEventsRes{
			Id:          reading.Id,
			UserId:      reading.UserId,
			Category:    reading.Category,
			Title:       reading.Title,
			Date:        reading.Date,
			Description: reading.Description,
		}

		count += 1
		events.Events = append(events.Events, event)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	events.Count = count
	return &events, nil
}
func (c *HistoricalEventsRepo) Update(req *mp.HistoricalEventsUpdateReq) (*mp.Void, error) {
	collection := c.db.Collection("historical_event")
	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	updateDoc := bson.M{}

	if req.Title != "" {
		updateDoc["title"] = req.Title
	}
	if req.Description != "" {
		updateDoc["description"] = req.Description
	}

	if len(updateDoc) == 0 {
		return nil, errors.New("no fields to update")
	}
	_, err := collection.UpdateOne(
		context.TODO(),
		bson.M{"id": req.Id},
		bson.M{"$set": updateDoc},
	)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
func (c *HistoricalEventsRepo) Delete(req *mp.ById) (*mp.Void, error) {
	collection := c.db.Collection("historical_event")
	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	_, err := collection.DeleteOne(context.TODO(), bson.M{"id": req.Id})
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (c *HistoricalEventsRepo) Context(req *mp.ContextReq) (*mp.ContextRes, error) {
	// |--------------------------------|
	// |-----------HISTORICAL-----------|
	// |--------------------------------|

	historical := c.db.Collection("historical_event")
	hFilter := bson.M{}
	if req.UserId != "" {
		hFilter["userid"] = req.UserId
	}
	if req.Date != "" {
		hFilter = bson.M{"date": primitive.Regex{Pattern: req.Date, Options: "i"}}
	}

	hCursor, err := historical.Find(context.TODO(), hFilter)
	if err != nil {
		return nil, err
	}
	defer hCursor.Close(context.TODO())

	var hEvents mp.HistoricalEventsGetAllRes

	for hCursor.Next(context.TODO()) {
		var reading mp.HistoricalEventsRes
		err := hCursor.Decode(&reading)
		if err != nil {
			return nil, err
		}

		pEvent := &mp.HistoricalEventsRes{
			Id:          reading.Id,
			UserId:      reading.UserId,
			Category:    reading.Category,
			Title:       reading.Title,
			Date:        reading.Date,
			Description: reading.Description,
		}

		hEvents.Events = append(hEvents.Events, pEvent)
	}

	if err := hCursor.Err(); err != nil {
		return nil, err
	}

	// |--------------------------------|
	// |------------PERSONAL------------|
	// |--------------------------------|

	personal := c.db.Collection("personal_event")
	pFilter := bson.M{}
	if req.UserId != "" {
		pFilter["userid"] = req.UserId
	}
	if req.Date != "" {
		pFilter = bson.M{"date": primitive.Regex{Pattern: req.Date, Options: "i"}}
	}

	pCursor, err := personal.Find(context.TODO(), pFilter)
	if err != nil {
		return nil, err
	}
	defer pCursor.Close(context.TODO())

	var pEvents mp.PersonalEventsGetAllRes

	for pCursor.Next(context.TODO()) {
		var preading mp.PersonalEventsRes
		err := pCursor.Decode(&preading)
		if err != nil {
			return nil, err
		}

		pEvent := &mp.PersonalEventsRes{
			Id:      preading.Id,
			UserId:  preading.UserId,
			Type:    preading.Type,
			Title:   preading.Title,
			Date:    preading.Date,
			Preview: preading.Preview,
		}

		pEvents.Events = append(pEvents.Events, pEvent)
	}

	if err := pCursor.Err(); err != nil {
		return nil, err
	}

	contextRes := mp.ContextRes{
		Historical: &hEvents,
		Personal:   &pEvents,
	}
	return &contextRes, nil

}
