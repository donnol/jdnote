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
	if *typ == "" || *field == "" || *fieldType == "" {
		log.Fatal("Bad param.")
	}

	// 生成
	content := sliceToMap(*typ, *field, *fieldType)

	// 保存
	fileName := fmt.Sprintf("%s_map.go", strings.ToLower(*typ))
	generate.Save(fileName, content)
}

// sliceToMap slice转为map，[]typ => map[field]typ
func sliceToMap(typ, field, fieldType string) []byte {
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
	
		// {{.typ}}SliceToMap convert slice to map
		func {{.typ}}SliceToMap(list []{{.typ}}) map[{{.fieldType}}]{{.typ}} {
			m := make(map[{{.fieldType}}]{{.typ}})
			for _, s := range list {
				m[s.{{.field}}] = s
			}
			return m
		}
	`
	return generate.MakeTemplate(temp, data)
}
