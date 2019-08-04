package note

import (
	"fmt"
	"strings"

	"github.com/donnol/jdnote/model"
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
	NoteID int `json:"noteID"` // 记录ID
	Param
}

// Check 检查
func (m ModParam) Check() error {
	if m.NoteID == 0 {
		return fmt.Errorf("Empty ID")
	}
	if err := m.Param.Check(); err != nil {
		return err
	}

	return nil
}

// Result 结果
type Result struct {
	NoteID    int    `json:"noteID"`    // 笔记ID
	UserName  string `json:"userName"`  // 用户名
	Title     string `json:"title"`     // 标题
	Detail    string `json:"detail"`    // 详情
	CreatedAt int64  `json:"createdAt"` // 创建时间
}

// PageParam 分页参数
type PageParam struct {
	Title  string `json:"title"`  // 标题
	Detail string `json:"detail"` // 详情

	model.CommonParam
}

// PageResult 分页结果
type PageResult struct {
	Total int      `json:"total"` // 总数
	List  []Result `json:"list"`  // 列表
}
