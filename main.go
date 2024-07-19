package main

import (
	"log"
	"net"
	cf "timeline/config"
	mp "timeline/genproto"
	"timeline/service"
	"timeline/storage/mongo"
	"timeline/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	config := cf.Load()

	db, err := postgres.NewPostgresStorage(config)

	if err != nil {
		panic(err)
	}

	mongo, err := mongo.NewMongoStorage(config)

	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", config.TIMELINE_SERVICE_PORT)

	if err != nil {
		log.Fatalf("Failed to listen tcp: %v", err)
	}

	s := grpc.NewServer()

	mp.RegisterCustomEventsServiceServer(s, service.NewCustomEventsService(db))
	mp.RegisterMilestonesServiceServer(s, service.NewMilestonesService(db))
	mp.RegisterHistoricalEventsServiceServer(s, service.NewHistoricalEventsService(mongo))
	mp.RegisterPersonalEventsServiceServer(s, service.NewPersonalEventsService(mongo))

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
