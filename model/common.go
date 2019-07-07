package model

// CommonParam 通用参数
type CommonParam struct {
	Size   int `json:"size"`   // 分页大小
	Offset int `json:"offset"` // 分页偏移
}

// DefaultCommonParam 默认值
var DefaultCommonParam = CommonParam{
	Size: 10,
}
