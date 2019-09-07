package main

import (
	"github.com/donnol/jdnote/utils/context"
	"github.com/donnol/jdnote/models/action"
	"github.com/donnol/jdnote/models/role"
	roleaction "github.com/donnol/jdnote/models/role_action"
	"github.com/donnol/jdnote/models/user"
	userrole "github.com/donnol/jdnote/models/user_role"
)

func main() {
	if err := initdb(); err != nil {
		panic(err)
	}
}

// initdb 初始化数据库
func initdb() error {
	var err error

	ctx := context.Default()

	// 角色
	re := role.Entity{
		Role: "ALL",
	}
	r := &role.Role{}
	if re.ID, err = r.Add(ctx, re); err != nil {
		return err
	}

	// 动作
	ae := action.Entity{
		Action: "ALL",
	}
	a := &action.Action{}
	if ae.ID, err = a.Add(ctx, ae); err != nil {
		return err
	}

	// 角色动作关联
	rae := roleaction.Entity{
		RoleID:   re.ID,
		ActionID: ae.ID,
	}
	ra := &roleaction.RoleAction{}
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
	u := &user.User{}
	if ue.ID, err = u.Add(ctx, ue); err != nil {
		return err
	}

	// 用户角色关联
	ure := userrole.Entity{
		UserID: ue.ID,
		RoleID: re.ID,
	}
	ur := &userrole.UserRole{}
	if _, err = ur.Add(ctx, ure); err != nil {
		return err
	}

	return nil
}
