package timeapi

import (
	"context"
	"time"

	"github.com/donnol/jdnote/services/timesrv"
	"github.com/donnol/tools/route"
)

type Time struct {
	timeSrv timesrv.ITime
}

func (t *Time) GetToUnix(ctx context.Context, p route.Param) (r route.Result, err error) {
	param := struct {
		Time time.Time `json:"time"`
	}{}
	if err = p.Parse(ctx, &param); err != nil {
		return
	}

	r.Data = t.timeSrv.ToUnix(param.Time)

	return
}

func (t *Time) GetToTime(ctx context.Context, p route.Param) (r route.Result, err error) {
	param := struct {
		Timestamp int64 `json:"timestamp"`
	}{}
	if err = p.Parse(ctx, &param); err != nil {
		return
	}

	r.Data = t.timeSrv.ToTime(param.Timestamp)

	return
}
