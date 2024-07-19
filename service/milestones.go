package service

import (
	"context"
	mp "timeline/genproto"
	st "timeline/storage/postgres"
)

type MilestonesService struct {
	storage st.Storage
	mp.UnimplementedMilestonesServiceServer
}

func NewMilestonesService(storage *st.Storage) *MilestonesService {
	return &MilestonesService{storage: *storage}
}

func (s *MilestonesService) Create(ctx context.Context, req *mp.MilestonesCreateReq) (*mp.Void, error) {
	return s.storage.MilestoneS.Create(req)
}
func (s *MilestonesService) GetById(ctx context.Context, id *mp.ById) (*mp.MilestonesGetByIdRes, error) {
	return s.storage.MilestoneS.GetById(id)
}
func (s *MilestonesService) GetAll(ctx context.Context, req *mp.MilestonesGetAllReq) (*mp.MilestonesGetAllRes, error) {
	return s.storage.MilestoneS.GetAll(req)
}
func (s *MilestonesService) Update(ctx context.Context, req *mp.MilestonesUpdateReq) (*mp.Void, error) {
	return s.storage.MilestoneS.Update(req)
}
func (s *MilestonesService) Delete(ctx context.Context, id *mp.ById) (*mp.Void, error) {
	return s.storage.MilestoneS.Delete(id)
}
