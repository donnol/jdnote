package commonmodel

import "time"

type IDBase struct {
	ID int `json:"id"` // 记录ID
}

// TableBase 表结构所需公用部分
type TableBase struct {
	IDBase
	CreatedAt time.Time `json:"createdAt" db:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"` // 创建时间
	CreatedBy int       `json:"createdBy" db:"created_by"` // 创建操作用户
	UpdatedBy int       `json:"updatedBy" db:"updated_by"` // 更新操作用户
}

// ResultBase 结果所需公用部分
type ResultBase struct {
	IDBase
	CreatedStamp    int64  `json:"createdStamp"`    // 创建时间时间戳
	UpdatedStamp    int64  `json:"updatedStamp"`    // 创建时间时间戳
	CreatedUserName string `json:"createdUserName"` // 创建操作用户名字
	UpdatedUserName string `json:"updatedUserName"` // 更新操作用户名字
}

func (vb *ResultBase) FromTableBase(tb TableBase, uf func(userIDs []int) map[int]string) {
	userMap := uf([]int{tb.CreatedBy, tb.UpdatedBy})

	vb.ID = tb.ID
	vb.CreatedStamp = tb.CreatedAt.Unix()
	vb.UpdatedStamp = tb.UpdatedAt.Unix()
	vb.CreatedUserName = userMap[tb.CreatedBy]
	vb.UpdatedUserName = userMap[tb.UpdatedBy]
}
