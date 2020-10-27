package timesrv

import "time"

type ITime interface {
	ToTime(int64) time.Time
	ToUnix(time.Time) int64
}

func New() ITime {
	return &timeImpl{}
}

type timeImpl struct {
}

func (impl *timeImpl) ToTime(in int64) time.Time {
	return time.Unix(in, 0)
}

func (impl *timeImpl) ToUnix(in time.Time) int64 {
	return in.Unix()
}
