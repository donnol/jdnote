package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/donnol/jdnote/utils/generate"
)

var (
	typ       = flag.String("type", "", "类型名称")
	field     = flag.String("field", "", "字段名称")
	fieldType = flag.String("fieldType", "", "字段类型")
)

func main() {
	// 解析参数
	flag.Parse()
	if *typ == "" {
		log.Fatal("Bad param.")
	}

	// 生成
	content := queryColumn(*typ, *field, *fieldType)

	// 保存
	fileName := fmt.Sprintf("%s_query_column.go", strings.ToLower(*typ))
	generate.Save(fileName, content)
}

// queryColumn 获取列表的指定列
func queryColumn(typ, field, fieldType string) []byte {
	formatContent := generate.WrapTemplate(func(pkgName string) []byte {
		// 使用text/template，生成模板
		data := map[string]interface{}{
			"pkg":       pkgName,
			"typ":       typ,
			"field":     field,
			"fieldType": fieldType,
		}
		content := makeTemplate(data)
		return content
	})

	return formatContent
}

func makeTemplate(data map[string]interface{}) []byte {
	const temp = `
		package {{.pkg}}
	
		// QueryColumn collect column from list
		func QueryColumn(list []{{.typ}}) []{{.fieldType}} {
			columns := make([]{{.fieldType}}, 0)
			for _, s := range list {
				columns = append(columns, s.{{.field}})
			}
			return columns
		}
	`
	return generate.MakeTemplate(temp, data)
}
