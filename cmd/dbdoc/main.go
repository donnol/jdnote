package main

import (
	"log"
	"os"
	"os/exec"

	"github.com/donnol/jdnote/models/actionmodel"
	"github.com/donnol/jdnote/models/roleactionmodel"
	"github.com/donnol/jdnote/models/rolemodel"
	"github.com/donnol/jdnote/models/usermodel"
	"github.com/donnol/jdnote/models/userrolemodel"
	"github.com/donnol/tools/dbdoc"
)

func main() {
	makeDoc()
	makeGraph()
}

var ins = []interface{}{
	&usermodel.Entity{},
	&rolemodel.Entity{},
	&userrolemodel.Entity{},
	&actionmodel.Entity{},
	&roleactionmodel.Entity{},
}

func makeDoc() {
	// 打开文件
	filename := "DB_README.md"
	w, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	// 标题
	_, err = w.WriteString("# 数据库表定义\n\n![ER图](./db_cicada.svg)\n\n")
	if err != nil {
		panic(err)
	}

	// 解析
	if err := dbdoc.Resolve(w,
		ins...,
	); err != nil {
		panic(err)
	}
}

func makeGraph() {
	filename := "DB_cicada.dot"
	makeDot(filename)
	execDot(filename)
}

func makeDot(filename string) {
	// 打开文件
	w, err := os.OpenFile(filename, os.O_CREATE|os.O_TRUNC|os.O_RDWR, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer w.Close()

	// 解析
	if err := dbdoc.ResolveGraph(w,
		ins...,
	); err != nil {
		panic(err)
	}
}

func execDot(filename string) {
	// 如果有安装dot命令，则执行命名生成svg图
	cmdName := "dot"
	if _, err := exec.LookPath(cmdName); err == nil {
		// dot -Tpng -o hello.png tmp.dot
		cmd := exec.Command(cmdName, "-Tsvg", "-odb_cicada.svg", filename)
		out, err := cmd.Output()
		if err != nil {
			panic(err)
		}
		if len(out) != 0 {
			log.Printf("Output: %s\n", out)
		}
	} else {
		log.Printf("Not find dot command\n")
	}
}
