package userroledata

// Entity 实体-对应数据库表
type Entity struct {
	ID     int `json:"id"`                                  // 记录ID
	UserID int `json:"userID" db:"user_id" rel:"t_user.id"` // 用户ID
	RoleID int `json:"roleID" db:"role_id" rel:"t_role.id"` // 角色ID
}
