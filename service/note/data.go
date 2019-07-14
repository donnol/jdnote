package note

// Param 参数
type Param struct {
	Title  string `json:"title"`  // 标题
	Detail string `json:"detail"` // 详情
}

// Result 结果
type Result struct {
	UserName  string `json:"userName"`  // 用户名
	Title     string `json:"title"`     // 标题
	Detail    string `json:"detail"`    // 详情
	CreatedAt int64  `json:"createdAt"` // 创建时间
}
