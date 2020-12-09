package filesrv

import "github.com/donnol/jdnote/utils/context"

type FileMock struct {
	AddFunc func(ctx context.Context, param AddParam) (result AddResult, err error)

	GetFunc func(ctx context.Context, param GetParam) (result GetResult, err error)
}

var _ IFile = &FileMock{}

func (mockRecv *FileMock) Add(ctx context.Context, param AddParam) (result AddResult, err error) {
	return mockRecv.AddFunc(ctx, param)
}

func (mockRecv *FileMock) Get(ctx context.Context, param GetParam) (result GetResult, err error) {
	return mockRecv.GetFunc(ctx, param)
}
