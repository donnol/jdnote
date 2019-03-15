package main

import (
	"github.com/donnol/jdnote/model/action"
	"github.com/donnol/jdnote/model/role"
	roleaction "github.com/donnol/jdnote/model/role_action"
	"github.com/donnol/jdnote/model/user"
	userrole "github.com/donnol/jdnote/model/user_role"
	pg "github.com/donnol/jdnote/store/db/postgresql"
)

func main() {
	if err := initdb(); err != nil {
		panic(err)
	}
}

// initdb 初始化数据库
func initdb() error {
	// 角色
	r := &role.Role{
		Role: "ALL",
	}
	r.DB = pg.New()
	if err := r.Add(); err != nil {
		return err
	}

	// 动作
	a := &action.Action{
		Action: "ALL",
	}
	a.DB = pg.New()
	if err := a.Add(); err != nil {
		return err
	}

	// 角色动作关联
	ra := &roleaction.RoleAction{
		RoleID:   r.ID,
		ActionID: a.ID,
	}
	ra.DB = pg.New()
	if err := ra.Add(); err != nil {
		return err
	}

	// 用户
	u := &user.User{
		Name:     "jd",
		Phone:    "13420693396",
		Email:    "jdlau@126.com",
		Password: "13420693396",
	}
	u.DB = pg.New()
	if err := u.Add(); err != nil {
		return err
	}

	// 用户角色关联
	ur := &userrole.UserRole{
		UserID: u.ID,
		RoleID: r.ID,
	}
	ur.DB = pg.New()
	if err := ur.Add(); err != nil {
		return err
	}

	return nil
}
