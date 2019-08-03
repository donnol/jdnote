package model

// CommonParam 通用参数
type CommonParam struct {
	Start int `json:"start"` // 分页开始
	Size  int `json:"size"`  // 分页大小
}

// DefaultCommonParam 默认值
var DefaultCommonParam = CommonParam{
	Size: 10,
}
