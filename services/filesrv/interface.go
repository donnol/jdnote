package filesrv

import "github.com/donnol/jdnote/utils/context"

type IFile interface {
	Add(ctx context.Context, param AddParam) (result AddResult, err error)
	Get(ctx context.Context, param GetParam) (result GetResult, err error)
}
