package main

import (
	"os"

	"github.com/donnol/jdnote/model/action"
	"github.com/donnol/jdnote/model/role"
	roleaction "github.com/donnol/jdnote/model/role_action"
	"github.com/donnol/jdnote/model/user"
	userrole "github.com/donnol/jdnote/model/user_role"
	"github.com/donnol/jdnote/utils/dbdoc"
)

func main() {
	// 打开文件
	filename := "DB_README.md"
	w, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	// 标题
	_, err = w.WriteString("# 数据库表定义\n\n")
	if err != nil {
		panic(err)
	}

	// 解析
	if err := dbdoc.Resolve(w,
		&user.User{},
		&role.Role{},
		&userrole.UserRole{},
		&action.Action{},
		&roleaction.RoleAction{},
	); err != nil {
		panic(err)
	}
}
