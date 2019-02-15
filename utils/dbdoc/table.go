package dbdoc

import (
	"fmt"
	"io"
	"strings"

	"github.com/donnol/jdnote/utils/reflectx"
)

// Table 表
type Table struct {
	Name        string  // 名字
	Comment     string  // 注释
	Description string  // 描述
	FieldList   []Field // 字段列表
	IndexList   []Index // 索引列表

	mapper     Mapper // 字段映射函数
	typeMapper Mapper // 字段类型映射函数

	doc []byte
}

// Mapper 映射函数类型
type Mapper func(string) string

// Field 字段
type Field struct {
	Name        string // 名字
	Type        string // 类型
	Nullable    bool   // 是否可null
	Primary     bool   // 是否主键
	Description string // 描述
	Index              // 索引
}

// Index 索引
type Index struct {
	Name      string  // 名字
	Unique    bool    // 是否唯一索引
	FieldList []Field // 涉及字段
}

// NewTable 新建表
func NewTable() *Table {
	return &Table{
		FieldList:  make([]Field, 0),
		IndexList:  make([]Index, 0),
		mapper:     fieldMapper,
		typeMapper: fieldTypeMapper,
	}
}

// New 新建
func (t *Table) New() *Table {
	return NewTable()
}

// Resolve 解析
func (t *Table) Resolve(v interface{}) *Table {
	var err error
	var vstruct reflectx.Struct
	vstruct, err = reflectx.ResolveStruct(v)
	checkError(err)

	vstructName := vstruct.Name
	nameList := strings.Split(vstructName, ".")
	t.Name = t.mapper(nameList[len(nameList)-1])
	t.Comment = vstruct.Comment
	t.Description = vstruct.Description

	var tf Field
	for _, sf := range vstruct.Fields {
		// 主键
		if sf.Name == "ID" {
			tf = Field{
				Name:        "id",
				Type:        t.typeMapper(sf.Type.Kind().String()),
				Primary:     true,
				Description: sf.Comment,
			}
		} else {
			tf = FieldByTag(sf.Tag)
			if tf.Name == "" {
				tf.Name = t.mapper(sf.Name)
			}
			if tf.Type == "" {
				tf.Type = t.typeMapper(sf.Type.Kind().String())
			}
			tf.Description = sf.Comment
		}

		t.FieldList = append(t.FieldList, tf)
	}

	return t
}

// SetComment 设置注释
func (t *Table) SetComment(comment string) *Table {
	t.Comment = comment
	return t
}

// SetDescription 设置描述
func (t *Table) SetDescription(description string) *Table {
	t.Description = description
	return t
}

// Write 写入f
func (t *Table) Write(w io.Writer) *Table {
	t = t.makeDoc()

	_, err := w.Write(t.doc)
	checkError(err)

	return t
}

// SetMapper 设置字段名映射方法
func (t *Table) SetMapper(f Mapper) *Table {
	t.mapper = f
	return t
}

// SetTypeMapper 设置类型映射方法
func (t *Table) SetTypeMapper(f Mapper) *Table {
	t.typeMapper = f
	return t
}

func (t *Table) makeDoc() *Table {

	leftAngle := "<"
	rightAngle := ">"
	format := "## %s\n\n%s\n\n%s\n%s索引：\n\n%s\n"
	fieldFormat := "| %s | %s | %v | %v | %s |\n"
	header := "| Field | Type | Nullable | Primary | Description |\n| :-: | :-: | :-: | :-: | :-: |\n"
	indexFormat := "* %s(%s: %s)\n"

	// 字段
	var field, index, fieldEnum string
	for _, tf := range t.FieldList {
		var nullableString, primaryString, description string
		if tf.Nullable {
			nullableString = "*"
		}
		if tf.Primary {
			primaryString = "*"
		}
		description = tf.Description
		if strings.Contains(description, leftAngle) &&
			strings.Contains(description, rightAngle) &&
			strings.Index(description, leftAngle) < strings.Index(description, rightAngle) {

			if fieldEnum != "" {
				fieldEnum += "\n"
			}
			fieldEnumLine := description[strings.Index(description, leftAngle)+1 : strings.Index(description, rightAngle)-1]
			fieldEnumList := strings.Split(fieldEnumLine, ":")
			for fei, fe := range fieldEnumList {
				if fei == 0 {
					fieldEnum += fe + ":\n\n"
				}
				if fei > 0 {
					fieldEnumValueList := strings.Split(fe, ";")
					for _, fev := range fieldEnumValueList {
						fieldEnum += fmt.Sprintf("* %s\n", strings.TrimRight(fev, ";"))
					}
				}
			}

			description = description[:strings.Index(description, leftAngle)]
		}
		field += fmt.Sprintf(fieldFormat, tf.Name, tf.Type, nullableString, primaryString, description)

		// 索引
		var indexName = tf.Index.Name
		var uniqueString string
		if tf.Index.Unique {
			uniqueString = "UNIQUE"
			if indexName == "" {
				indexName = tf.Name
			}
		}
		if indexName != "" {
			index += fmt.Sprintf(indexFormat, uniqueString, indexName, tf.Name)
		}
	}

	description := t.Description + "\n\n" + fieldEnum
	if description != "" {
		description += "\n\n"
	}
	content := fmt.Sprintf(format, t.Name, t.Comment, header+field, description, index)
	t.doc = []byte(content)

	return t
}
