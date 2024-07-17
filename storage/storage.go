package storage

import (
	mp "timeline/genproto"
)

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

// type MediasI interface {
// 	Create(*mp.MediasCreateReq) (*mp.Void, error)
// 	GetById(*mp.ById) (*mp.MediasGetByIdRes, error)
// 	GetAll(*mp.MediasGetAllReq) (*mp.MediasGetAllRes, error)
// 	Update(*mp.MediasUpdateReq) (*mp.Void, error)
// 	Delete(*mp.ById) (*mp.Void, error)
// }

// type SharedMemoriesI interface {
// 	Create(*mp.SharedMemoriesCreateReq) (*mp.Void, error)
// 	GetById(*mp.ById) (*mp.SharedMemoriesGetByIdRes, error)
// 	GetAll(*mp.SharedMemoriesGetAllReq) (*mp.SharedMemoriesGetAllRes, error)
// 	Update(*mp.SharedMemoriesUpdateReq) (*mp.Void, error)
// 	Delete(*mp.ById) (*mp.Void, error)
// }
