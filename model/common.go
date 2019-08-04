package model

// CommonParam 通用参数
type CommonParam struct {
	BeginTime int64 `json:"beginTime"` // 开始时间
	EndTime   int64 `json:"endTime"`   // 结束时间

	PageIndex int `json:"pageIndex"` // 分页开始
	PageSize  int `json:"pageSize"`  // 分页大小
}

// DefaultCommonParam 默认值
var DefaultCommonParam = CommonParam{
	PageSize: 10,
}
