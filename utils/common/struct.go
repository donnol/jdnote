package common

// Param 通用参数
type Param struct {
	BeginTime int64 `json:"beginTime"` // 开始时间
	EndTime   int64 `json:"endTime"`   // 结束时间

	PageIndex int `json:"pageIndex"` // 分页开始
	PageSize  int `json:"pageSize"`  // 分页大小

	OnlyPublish bool `json:"-"` // 只展示发布记录
}

// DefaultParam 默认值
var DefaultParam = Param{
	PageSize: 10,
}
