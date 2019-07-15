package note

import (
	"fmt"
	"strings"
)

// Param 参数
type Param struct {
	Title  string `json:"title"`  // 标题
	Detail string `json:"detail"` // 详情
}

// Check 检查
func (p Param) Check() error {
	if strings.TrimSpace(p.Title) == "" {
		return fmt.Errorf("Empty Title")
	}
	if strings.TrimSpace(p.Detail) == "" {
		return fmt.Errorf("Empty Detail")
	}

	return nil
}

// ModParam 修改参数
type ModParam struct {
	ID int `json:"id"` // 记录ID
	Param
}

// Check 检查
func (m ModParam) Check() error {
	if m.ID == 0 {
		return fmt.Errorf("Empty ID")
	}
	if err := m.Param.Check(); err != nil {
		return err
	}

	return nil
}

// Result 结果
type Result struct {
	UserName  string `json:"userName"`  // 用户名
	Title     string `json:"title"`     // 标题
	Detail    string `json:"detail"`    // 详情
	CreatedAt int64  `json:"createdAt"` // 创建时间
}
