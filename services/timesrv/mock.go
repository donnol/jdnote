package timesrv

import "time"

type TimeMock struct {
	ToTimeFunc func(int64) time.Time

	ToUnixFunc func(time.Time) int64
}

var _ ITime = &TimeMock{}

func (mockRecv *TimeMock) ToTime(p0 int64) time.Time {
	return mockRecv.ToTimeFunc(p0)
}

func (mockRecv *TimeMock) ToUnix(p0 time.Time) int64 {
	return mockRecv.ToUnixFunc(p0)
}
