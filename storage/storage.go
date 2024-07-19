package storage

import (
	mp "timeline/genproto"
)

type StorageI interface {
	HistoricalEvents() HistoricalEventsI
	PersonalEvents() PersonalEventsI
}

type CustomEventsI interface {
	Create(*mp.CustomEventsCreateReq) (*mp.Void, error)
	GetById(*mp.ById) (*mp.CustomEventsGetByIdRes, error)
	GetAll(*mp.CustomEventsGetAllReq) (*mp.CustomEventsGetAllRes, error)
	Update(*mp.CustomEventsUpdateReq) (*mp.Void, error)
	Delete(*mp.ById) (*mp.Void, error)
}

type MilestonesI interface {
	Create(*mp.MilestonesCreateReq) (*mp.Void, error)
	GetById(*mp.ById) (*mp.MilestonesGetByIdRes, error)
	GetAll(*mp.MilestonesGetAllReq) (*mp.MilestonesGetAllRes, error)
	Update(*mp.MilestonesUpdateReq) (*mp.Void, error)
	Delete(*mp.ById) (*mp.Void, error)
}

type PersonalEventsI interface {
	Create(*mp.PersonalEventsRes) (*mp.Void, error)
	GetById(*mp.ById) (*mp.PersonalEventsGetByIdRes, error)
	GetAll(*mp.PersonalEventsGetAllReq) (*mp.PersonalEventsGetAllRes, error)
	Update(*mp.PersonalEventsUpdateReq) (*mp.Void, error)
	Delete(*mp.ById) (*mp.Void, error)
}

type HistoricalEventsI interface {
	Create(*mp.HistoricalEventsRes) (*mp.Void, error)
	GetById(*mp.ById) (*mp.HistoricalEventsGetByIdRes, error)
	GetAll(*mp.HistoricalEventsGetAllReq) (*mp.HistoricalEventsGetAllRes, error)
	Update(*mp.HistoricalEventsUpdateReq) (*mp.Void, error)
	Delete(*mp.ById) (*mp.Void, error)
	Context(*mp.ContextReq) (*mp.ContextRes, error)
}
