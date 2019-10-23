package roledata

// Entity 实体-对应数据库表
type Entity struct {
	ID   int    `json:"id"`   // 记录ID
	Role string `json:"role"` // 角色
}
