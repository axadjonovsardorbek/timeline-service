package main

import (
	"log"
	cf "timeline/config"
	mp "timeline/genproto"
	"timeline/storage/postgres"
	"timeline/service"
	"net"

	"google.golang.org/grpc"
)

func main() {
	config := cf.Load()

	db, err := postgres.NewPostgresStorage(config)

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

	log.Printf("server listening at %v", listener.Addr())
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
