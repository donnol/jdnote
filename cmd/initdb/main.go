package main

import (
	"github.com/donnol/jdnote/models"
	"github.com/donnol/jdnote/models/action/actiondata"
	"github.com/donnol/jdnote/models/role/roledata"
	"github.com/donnol/jdnote/models/roleaction/roleactiondata"
	"github.com/donnol/jdnote/models/user/userdata"
	"github.com/donnol/jdnote/models/userrole/userroledata"
)

func main() {
	if err := initdb(); err != nil {
		panic(err)
	}
}

// initdb 初始化数据库
func initdb() error {
	var err error

	ctx := models.DefaultCtx()

	// 角色
	re := roledata.Entity{
		Role: "ALL",
	}
	r := &roledata.Role{}
	if re.ID, err = r.Add(ctx, re); err != nil {
		return err
	}

	// 动作
	ae := actiondata.Entity{
		Action: "ALL",
	}
	a := &actiondata.Action{}
	if ae.ID, err = a.Add(ctx, ae); err != nil {
		return err
	}

	// 角色动作关联
	rae := roleactiondata.Entity{
		RoleID:   re.ID,
		ActionID: ae.ID,
	}
	ra := &roleactiondata.RoleAction{}
	if rae.ID, err = ra.Add(ctx, rae); err != nil {
		return err
	}

	// 用户
	ue := userdata.Entity{
		Name:     "jd",
		Phone:    "jd",
		Email:    "jdlau@126.com",
		Password: "jd",
	}
	u := &userdata.User{}
	if ue.ID, err = u.Add(ctx, ue); err != nil {
		return err
	}

	// 用户角色关联
	ure := userroledata.Entity{
		UserID: ue.ID,
		RoleID: re.ID,
	}
	ur := &userroledata.UserRole{}
	if _, err = ur.Add(ctx, ure); err != nil {
		return err
	}

	return nil
}
