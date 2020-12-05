package commonmodel

import "time"

type Base struct {
	ID        int       `json:"id"`                        // 记录ID
	CreatedAt time.Time `json:"createdAt" db:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"` // 创建时间
	CreatedBy int       `json:"createdBy" db:"created_by"` // 创建操作用户
	UpdatedBy int       `json:"updatedBy" db:"updated_by"` // 更新操作用户
}
