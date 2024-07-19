package service

import (
	"context"
	mp "timeline/genproto"
	st "timeline/storage/mongo"
)

type PersonalEventsService struct {
	storage st.Storage
	mp.UnimplementedPersonalEventsServiceServer
}

func NewPersonalEventsService(storage *st.Storage) *PersonalEventsService {
	return &PersonalEventsService{storage: *storage}
}

func (s *PersonalEventsService) Create(ctx context.Context, req *mp.PersonalEventsRes) (*mp.Void, error) {
	return s.storage.PersonalEventsS.Create(req)
}
func (s *PersonalEventsService) GetById(ctx context.Context, id *mp.ById) (*mp.PersonalEventsGetByIdRes, error) {
	return s.storage.PersonalEventsS.GetById(id)
}
func (s *PersonalEventsService) GetAll(ctx context.Context, req *mp.PersonalEventsGetAllReq) (*mp.PersonalEventsGetAllRes, error) {
	return s.storage.PersonalEventsS.GetAll(req)
}
func (s *PersonalEventsService) Update(ctx context.Context, req *mp.PersonalEventsUpdateReq) (*mp.Void, error) {
	return s.storage.PersonalEventsS.Update(req)
}
func (s *PersonalEventsService) Delete(ctx context.Context, id *mp.ById) (*mp.Void, error) {
	return s.storage.PersonalEventsS.Delete(id)
}
