package exp

type DBMock struct {
	GetFunc func(id uint) (Data, error)
}

var _ DB = &DBMock{}

func (mockRecv *DBMock) Get(id uint) (Data, error) {
	return mockRecv.GetFunc(id)
}
