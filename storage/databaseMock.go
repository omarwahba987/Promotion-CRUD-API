package storage

type dbRepository interface {
	AddPromotion()
	GetPromotion()
	UpdatePromotion()
	DeletePromotion()
}

type DbMock struct {
}

func (d DbMock) AddPromotion() {
	panic("implement me")
}

func (d DbMock) GetPromotion() {
	panic("implement me")
}

func (d DbMock) UpdatePromotion() {
	panic("implement me")
}

func (d DbMock) DeletePromotion() {
	panic("implement me")
}
