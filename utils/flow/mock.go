package flow

import "context"

type FlowMock struct {
	CleanupFunc func(context.Context) error

	DoFunc func(context.Context) error

	FinishFunc func(context.Context) error

	PrepareFunc func(context.Context) error
}

var _ Flow = &FlowMock{}

func (mockRecv *FlowMock) Cleanup(p0 context.Context) error {
	return mockRecv.CleanupFunc(p0)
}

func (mockRecv *FlowMock) Do(p0 context.Context) error {
	return mockRecv.DoFunc(p0)
}

func (mockRecv *FlowMock) Finish(p0 context.Context) error {
	return mockRecv.FinishFunc(p0)
}

func (mockRecv *FlowMock) Prepare(p0 context.Context) error {
	return mockRecv.PrepareFunc(p0)
}

type HookMock struct {
	DoFunc func(context.Context) error

	WhenFunc func() Timing
}

var _ Hook = &HookMock{}

func (mockRecv *HookMock) Do(p0 context.Context) error {
	return mockRecv.DoFunc(p0)
}

func (mockRecv *HookMock) When() Timing {
	return mockRecv.WhenFunc()
}
