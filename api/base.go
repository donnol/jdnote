package api

import (
	"github.com/donnol/jdnote/models/auth"
	"github.com/donnol/jdnote/utils/context"
	"github.com/pkg/errors"
)

// Base 基底
type Base struct {
	AuthAo auth.Auth
}

// CheckLogin 检查登录态
func (b Base) CheckLogin(ctx context.Context) error {
	if ctx.UserID() == 0 {
		return errors.Errorf("Please login")
	}
	return nil
}

// CheckPerm 检查权限
func (b Base) CheckPerm(ctx context.Context, perms []string) error {
	// 先要登录
	if err := b.CheckLogin(ctx); err != nil {
		return err
	}

	// 检查权限
	if err := b.AuthAo.CheckPerm(ctx, perms); err != nil {
		return err
	}

	return nil
}

// AddResult 添加记录后的结果
type AddResult struct {
	ID int `json:"id"` // 新纪录ID
}
