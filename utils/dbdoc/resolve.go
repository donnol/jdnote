package dbdoc

import "io"

// Resolve 解析多个结构体，并将它们写到w
func Resolve(w io.Writer, v ...interface{}) error {
	var table = NewTable()

	for _, s := range v {
		table.New().Resolve(s).Write(w)
	}

	return nil
}
