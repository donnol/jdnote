package json

import (
	"bytes"
	"encoding/json"
	"io"
)

// Indent 格式化后输出
func Indent(w io.Writer, v interface{}) {
	var b []byte
	if vb, ok := v.([]byte); ok {
		b = vb
	} else {
		var err error
		b, err = json.Marshal(v)
		if err != nil {
			panic(err)
		}
	}
	var out bytes.Buffer
	if err := json.Indent(&out, b, "", "\t"); err != nil {
		panic(err)
	}
	if _, err := out.WriteTo(w); err != nil {
		panic(err)
	}
}
