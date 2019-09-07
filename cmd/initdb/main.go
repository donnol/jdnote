package main

import (
	"github.com/donnol/jdnote/models"
	"github.com/donnol/jdnote/models/actionao/actiondb"
	"github.com/donnol/jdnote/models/roleactionao/roleactiondb"
	"github.com/donnol/jdnote/models/roleao/roledb"
	"github.com/donnol/jdnote/models/userao/userdb"
	"github.com/donnol/jdnote/models/userroleao/userroledb"
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
	re := roledb.Entity{
		Role: "ALL",
	}
	r := &roledb.Role{}
	if re.ID, err = r.Add(ctx, re); err != nil {
		return err
	}

	// 动作
	ae := actiondb.Entity{
		Action: "ALL",
	}
	a := &actiondb.Action{}
	if ae.ID, err = a.Add(ctx, ae); err != nil {
		return err
	}

	// 角色动作关联
	rae := roleactiondb.Entity{
		RoleID:   re.ID,
		ActionID: ae.ID,
	}
	ra := &roleactiondb.RoleAction{}
	if rae.ID, err = ra.Add(ctx, rae); err != nil {
		return err
	}

	// 用户
	ue := userdb.Entity{
		Name:     "jd",
		Phone:    "jd",
		Email:    "jdlau@126.com",
		Password: "jd",
	}
	u := &userdb.User{}
	if ue.ID, err = u.Add(ctx, ue); err != nil {
		return err
	}

	// 用户角色关联
	ure := userroledb.Entity{
		UserID: ue.ID,
		RoleID: re.ID,
	}
	ur := &userroledb.UserRole{}
	if _, err = ur.Add(ctx, ure); err != nil {
		return err
	}

	return nil
}
