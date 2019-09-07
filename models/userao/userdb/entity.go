package userdb

// Entity 实体-对应表结构
type Entity struct {
	ID       int    `json:"id" form:"id"`             // 记录ID
	Name     string `json:"name" form:"name"`         // 用户名
	Phone    string `json:"phone" form:"phone"`       // 手机号码
	Email    string `json:"email" form:"email"`       // 邮箱
	Password string `json:"password" form:"password"` // 密码
}

// Filter 过滤器
func (e Entity) Filter() interface{} {
	// 将密码置空
	e.Password = ""

	return e
}
