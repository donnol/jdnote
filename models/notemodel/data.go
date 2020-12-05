package notemodel

import (
	"strings"
	"time"
)

// Entity 笔记
// 与表结构对应，其它结构在此基础上进行增改
type Entity struct {
	ID        int       `json:"id"`                                  // 记录ID
	UserID    int       `json:"userID" db:"user_id" rel:"t_user.id"` // 用户ID
	Title     string    `json:"title"`                               // 标题
	Status    Status    `json:"status" db:"status"`                  // 状态: 1 草稿;2 发布;
	Detail    string    `json:"detail"`                              // 详情
	CreatedAt time.Time `json:"createdAt" db:"created_at"`           // 创建时间
}

// EntityList 列表
type EntityList []Entity

// Process 处理
func (e Entity) Process() Entity {
	ne := e

	// 不存在标题，则截取详情的一部分作为标题
	if ne.Title == "" {
		details := strings.Split(ne.Detail, "\n")
		if len(details) != 0 {
			ne.Title = details[0]
		}
	}

	return ne
}

// Filter 过滤器
func (e Entity) Filter() interface{} {
	// 如有需要，可将指定字段置空

	return e
}

// Join 连接
// 这里又怎么会知道要跟谁join呢？所以只能在srv里做，这里做是没意义的
func (e Entity) Join(f func(Entity) Entity) Entity {
	ne := f(e)
	return ne
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
