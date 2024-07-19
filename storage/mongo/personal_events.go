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

type PersonalEventsRepo struct {
	db *mongo.Database
}

func NewPersonalEventsRepo(db *mongo.Database) *PersonalEventsRepo {
	return &PersonalEventsRepo{db: db}
}

func (c *PersonalEventsRepo) Create(req *mp.PersonalEventsRes) (*mp.Void, error) {
	collection := c.db.Collection("personal_event")
	id := uuid.New().String()

	req.Id = id

	res, err := collection.InsertOne(context.TODO(), req)

	fmt.Println(res)
	if err != nil {
		log.Println("Error while creating personal event: ", err)
		return nil, err
	}

	log.Println("Successfully created personal event")

	return nil, nil
}
func (c *PersonalEventsRepo) GetById(req *mp.ById) (*mp.PersonalEventsGetByIdRes, error) {
	collection := c.db.Collection("personal_event")
	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	var event mp.PersonalEventsGetByIdRes
	err := collection.FindOne(context.TODO(), bson.M{"id": req.Id}).Decode(&event.Event)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("event not found")
	} else if err != nil {
		return nil, err
	}

	return &event, nil
}
func (c *PersonalEventsRepo) GetAll(req *mp.PersonalEventsGetAllReq) (*mp.PersonalEventsGetAllRes, error) {
	collection := c.db.Collection("personal_event")
	mongoFilter := bson.M{}
	if req.UserId != "" {
		mongoFilter["userid"] = req.UserId
	}
	if req.Date != "" {
		mongoFilter = bson.M{"date": primitive.Regex{Pattern: req.Date, Options: "i"}}
	}
	if req.Type != "" {
		mongoFilter = bson.M{"type": primitive.Regex{Pattern: req.Type, Options: "i"}}
	}

	cursor, err := collection.Find(context.TODO(), mongoFilter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var events mp.PersonalEventsGetAllRes

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
		var reading mp.PersonalEventsRes
		err := cursor.Decode(&reading)
		if err != nil {
			return nil, err
		}

		event := &mp.PersonalEventsRes{
			Id:      reading.Id,
			UserId:  reading.UserId,
			Type:    reading.Type,
			Title:   reading.Title,
			Date:    reading.Date,
			Preview: reading.Preview,
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
func (c *PersonalEventsRepo) Update(req *mp.PersonalEventsUpdateReq) (*mp.Void, error) {
	collection := c.db.Collection("personal_event")
	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	updateDoc := bson.M{}

	if req.Title != "" {
		updateDoc["title"] = req.Title
	}
	if req.Preview != "" {
		updateDoc["preview"] = req.Preview
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
func (c *PersonalEventsRepo) Delete(req *mp.ById) (*mp.Void, error) {
	collection := c.db.Collection("personal_event")
	if req.Id == "" {
		return nil, errors.New("id cannot be empty")
	}

	_, err := collection.DeleteOne(context.TODO(), bson.M{"id": req.Id})
	if err != nil {
		return nil, err
	}

	return nil, nil
}
