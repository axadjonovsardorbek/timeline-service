package service

import (
	"context"
	mp "timeline/genproto"
	st "timeline/storage/mongo"
)

type HistoricalEventsService struct {
	storage st.Storage
	mp.UnimplementedHistoricalEventsServiceServer
}

func NewHistoricalEventsService(storage *st.Storage) *HistoricalEventsService {
	return &HistoricalEventsService{storage: *storage}
}

func (s *HistoricalEventsService) Create(ctx context.Context, req *mp.HistoricalEventsRes) (*mp.Void, error) {
	return s.storage.HistoricalEventsS.Create(req)
}
func (s *HistoricalEventsService) GetById(ctx context.Context, id *mp.ById) (*mp.HistoricalEventsGetByIdRes, error) {
	return s.storage.HistoricalEventsS.GetById(id)
}
func (s *HistoricalEventsService) GetAll(ctx context.Context, req *mp.HistoricalEventsGetAllReq) (*mp.HistoricalEventsGetAllRes, error) {
	return s.storage.HistoricalEventsS.GetAll(req)
}
func (s *HistoricalEventsService) Update(ctx context.Context, req *mp.HistoricalEventsUpdateReq) (*mp.Void, error) {
	return s.storage.HistoricalEventsS.Update(req)
}
func (s *HistoricalEventsService) Delete(ctx context.Context, id *mp.ById) (*mp.Void, error) {
	return s.storage.HistoricalEventsS.Delete(id)
}
func (s *HistoricalEventsService) Context(ctx context.Context, req *mp.ContextReq) (*mp.ContextRes, error){
	return s.storage.HistoricalEventsS.Context(req)
}