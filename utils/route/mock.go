package route

import "context"

type NewerMock struct {
	NewFunc func() interface{}
}

var _ Newer = &NewerMock{}

func (mockRecv *NewerMock) New() interface{} {
	return mockRecv.NewFunc()
}

type CheckerMock struct {
	CheckFunc func(context.Context) error
}

var _ Checker = &CheckerMock{}

func (mockRecv *CheckerMock) Check(p0 context.Context) error {
	return mockRecv.CheckFunc(p0)
}

type FilterMock struct {
	FilterFunc func() interface{}
}

var _ Filter = &FilterMock{}

func (mockRecv *FilterMock) Filter() interface{} {
	return mockRecv.FilterFunc()
}
