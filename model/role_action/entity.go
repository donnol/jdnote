package roleaction

// Entity 实体-对应数据库表
type Entity struct {
	ID       int `json:"id"`                                        // 记录ID
	RoleID   int `json:"roleID" db:"role_id" rel:"t_role.id"`       // 角色ID
	ActionID int `json:"actionID" db:"action_id" rel:"t_action.id"` // 动作ID
}
