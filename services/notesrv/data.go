package notesrv

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/donnol/jdnote/models/notemodel"
	"github.com/donnol/jdnote/utils/common"
	"github.com/donnol/jdnote/utils/context"
)

// Param 参数
type Param struct {
	Title  string `json:"title"`  // 标题
	Detail string `json:"detail"` // 详情
}

// Check 检查
func (p Param) Check(ctx context.Context) error {
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
func (m ModParam) Check(ctx context.Context) error {
	if m.NoteID == 0 {
		return fmt.Errorf("Empty ID")
	}
	if err := m.Param.Check(ctx); err != nil {
		return err
	}

	return nil
}

// PageParam 分页参数
type PageParam struct {
	Title  string `json:"title"`  // 标题
	Detail string `json:"detail"` // 详情

	common.Param
}

// PageResult 分页结果
type PageResult struct {
	Total int      `json:"total"` // 总数
	List  []Result `json:"list"`  // 列表
}

// Result 结果
type Result struct {
	NoteID    int    `json:"noteID"`    // 笔记ID
	UserName  string `json:"userName"`  // 用户名
	Title     string `json:"title"`     // 标题
	Detail    string `json:"detail"`    // 详情
	CreatedAt int64  `json:"createdAt"` // 创建时间
}

// Init 初始化
func (r Result) Init(single notemodel.Entity) (Result, error) {
	tmp := r

	tmp.NoteID = single.ID
	tmp.UserName = strconv.Itoa(single.UserID)
	tmp.Title = single.Title
	tmp.Detail = single.Detail
	tmp.CreatedAt = single.CreatedAt.Unix()

	return tmp, nil
}
