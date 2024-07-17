package service

import (
	"context"
	mp "timeline/genproto"
	st "timeline/storage/postgres"
)

type CustomEventsService struct {
	storage st.Storage
	mp.UnimplementedCustomEventsServiceServer
}

func NewCustomEventsService(storage *st.Storage) *CustomEventsService {
	return &CustomEventsService{storage: *storage}
}

func (s *CustomEventsService) Create(ctx context.Context, req *mp.CustomEventsCreateReq) (*mp.Void, error) {
	return s.storage.EventS.Create(req)
}
func (s *CustomEventsService) GetById(ctx context.Context, id *mp.ById) (*mp.CustomEventsGetByIdRes, error) {
	return s.storage.EventS.GetById(id)
}
func (s *CustomEventsService) GetAll(ctx context.Context, req *mp.CustomEventsGetAllReq) (*mp.CustomEventsGetAllRes, error) {
	return s.storage.EventS.GetAll(req)
}
func (s *CustomEventsService) Update(ctx context.Context, req *mp.CustomEventsUpdateReq) (*mp.Void, error) {
	return s.storage.EventS.Update(req)
}
func (s *CustomEventsService) Delete(ctx context.Context, id *mp.ById) (*mp.Void, error) {
	return s.storage.EventS.Delete(id)
}
