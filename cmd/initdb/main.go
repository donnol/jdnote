package main

import (
	"context"

	"github.com/donnol/jdnote/internal/initializers"
	"github.com/donnol/jdnote/models/actionmodel"
	"github.com/donnol/jdnote/models/roleactionmodel"
	"github.com/donnol/jdnote/models/rolemodel"
	"github.com/donnol/jdnote/models/usermodel"
	"github.com/donnol/jdnote/models/userrolemodel"
	"github.com/donnol/jdnote/stores/actionstore"
	"github.com/donnol/jdnote/stores/roleactionstore"
	"github.com/donnol/jdnote/stores/rolestore"
	"github.com/donnol/jdnote/stores/userrolestore"
	"github.com/donnol/jdnote/stores/userstore"
)

func main() {
	if err := initdb(); err != nil {
		panic(err)
	}
}

// initdb 初始化数据库
func initdb() error {
	var err error

	ctx := context.Background()
	app := initializers.New()
	db := app.GetDB()

	// 角色
	re := rolemodel.Entity{
		Role: "ALL",
	}
	r := rolestore.New(db)
	if re.ID, err = r.Add(ctx, re); err != nil {
		return err
	}

	// 动作
	ae := actionmodel.Entity{
		Action: "ALL",
	}
	a := actionstore.New(db)
	if ae.ID, err = a.Add(ctx, ae); err != nil {
		return err
	}

	// 角色动作关联
	rae := roleactionmodel.Entity{
		RoleID:   re.ID,
		ActionID: ae.ID,
	}
	ra := roleactionstore.New(db)
	if rae.ID, err = ra.Add(ctx, rae); err != nil {
		return err
	}

	// 用户
	ue := usermodel.Entity{
		Name:     "jd",
		Phone:    "jd",
		Email:    "jdlau@126.com",
		Password: "jd",
	}
	u := userstore.New(db)
	if ue.ID, err = u.Add(ctx, ue); err != nil {
		return err
	}

	// 用户角色关联
	ure := userrolemodel.Entity{
		UserID: ue.ID,
		RoleID: re.ID,
	}
	ur := userrolestore.New(db)
	if _, err = ur.Add(ctx, ure); err != nil {
		return err
	}

	return nil
}
