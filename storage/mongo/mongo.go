package mongo

import (
	"context"
	"fmt"
	"log"
	"timeline/config"
	"timeline/storage"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	Db                *mongo.Database
	PersonalEventsS   storage.PersonalEventsI
	HistoricalEventsS storage.HistoricalEventsI
}

func NewMongoStorage(config config.Config) (*Storage, error) {
	// uri := fmt.Sprintf("mongodb://%s:%d@%s:%d/%s",config.MONGO_DB_USER, config.MONGO_DB_PASS, config.MONGO_DB_HOST, config.MONGO_DB_PORT, config.MONGO_DB_NAME)
	uri := fmt.Sprintf("mongodb://%s:%d", config.MONGO_DB_HOST, config.MONGO_DB_PORT)

	clientOptions := options.Client().ApplyURI(uri).SetAuth(options.Credential{Username: config.MONGO_DB_USER, Password: config.MONGO_DB_PASS})

	// clientOptions := options.Client().ApplyURI(uri).SetAuth(options.Credential{Username: "sardorbek", Password: "1111"})
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(config.MONGO_DB_NAME)

	personal := NewPersonalEventsRepo(db)
	historical := NewHistoricalEventsRepo(db)

	fmt.Println("Connected to MongoDB!")
	return &Storage{
		Db:                db,
		PersonalEventsS:   personal,
		HistoricalEventsS: historical,
	}, nil
}

func (s *Storage) PersonalEvents() storage.PersonalEventsI {
	if s.PersonalEventsS == nil {
		s.PersonalEventsS = NewPersonalEventsRepo(s.Db)
	}
	return s.PersonalEventsS
}

func (s *Storage) HistoricalEvents() storage.HistoricalEventsI {
	if s.HistoricalEventsS == nil {
		s.HistoricalEventsS = NewHistoricalEventsRepo(s.Db)
	}
	return s.HistoricalEventsS
}
