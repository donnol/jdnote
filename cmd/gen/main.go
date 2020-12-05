package main

import (
	"bytes"
	"fmt"
	"text/template"
)

func main() {
	if err := genTableSchema(); err != nil {
		panic(err)
	}
}

func genTableSchema() (err error) {
	const file = "db_add_table_template.tpl"

	var t *template.Template
	t, err = template.New("tpl").ParseFiles(file)
	if err != nil {
		return
	}

	w := new(bytes.Buffer)

	// 指定用户名，表名，特有字段
	for _, cas := range []struct {
		UserName      string
		TableName     string
		SpecialFields string
	}{
		{
			UserName:  "jd",
			TableName: "t_test",
			SpecialFields: `
    size integer NOT NULL DEFAULT 0::integer, 
    content bytea NOT NULL DEFAULT ''::bytea 
			`,
		},
		{
			UserName:  "jd",
			TableName: "t_target",
			SpecialFields: `
	name text NOT NULL DEFAULT ''::text, 
	file_content_id integer NOT NULL
			`,
		},
	} {
		err = t.ExecuteTemplate(w, file, cas)
		if err != nil {
			return
		}
	}

	fmt.Printf("=== gen table:\n")
	fmt.Printf("%s\n", w.Bytes())

	return
}
