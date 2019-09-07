package note

import "time"

// Entity 笔记
type Entity struct {
	ID        int       `json:"id"`                                  // 记录ID
	UserID    int       `json:"userID" db:"user_id" rel:"t_user.id"` // 用户ID
	Title     string    `json:"title"`                               // 标题
	Detail    string    `json:"detail"`                              // 详情
	CreatedAt time.Time `json:"createdAt" db:"created_at"`           // 创建时间
}

// Pages 分页列表
type Pages []struct {
	Entity
	Total int
}

// Transfer 转换
func (p Pages) Transfer() (res []Entity, total int, err error) {
	for i, single := range p {
		if i == 0 {
			total = single.Total
		}
		res = append(res, single.Entity)
	}

	return
}
