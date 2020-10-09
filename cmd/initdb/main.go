package main

import (
	"context"

	"github.com/donnol/jdnote/app"
	"github.com/donnol/jdnote/models/action"
	"github.com/donnol/jdnote/models/role"
	"github.com/donnol/jdnote/models/roleaction"
	"github.com/donnol/jdnote/models/user"
	"github.com/donnol/jdnote/models/userrole"
)

func main() {
	if err := initdb(); err != nil {
		panic(err)
	}
}

// initdb 初始化数据库
func initdb() error {
	var err error

	sctx := context.Background()
	_, ctx := app.New(sctx)

	// 角色
	re := role.Entity{
		Role: "ALL",
	}
	r := role.New()
	if re.ID, err = r.Add(ctx, re); err != nil {
		return err
	}

	// 动作
	ae := action.Entity{
		Action: "ALL",
	}
	a := action.New()
	if ae.ID, err = a.Add(ctx, ae); err != nil {
		return err
	}

	// 角色动作关联
	rae := roleaction.Entity{
		RoleID:   re.ID,
		ActionID: ae.ID,
	}
	ra := roleaction.New()
	if rae.ID, err = ra.Add(ctx, rae); err != nil {
		return err
	}

	// 用户
	ue := user.Entity{
		Name:     "jd",
		Phone:    "jd",
		Email:    "jdlau@126.com",
		Password: "jd",
	}
	u := user.New()
	if ue.ID, err = u.Add(ctx, ue); err != nil {
		return err
	}

	// 用户角色关联
	ure := userrole.Entity{
		UserID: ue.ID,
		RoleID: re.ID,
	}
	ur := userrole.New()
	if _, err = ur.Add(ctx, ure); err != nil {
		return err
	}

	return nil
}
