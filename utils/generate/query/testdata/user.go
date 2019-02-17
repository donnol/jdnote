//go:generate go run ../query_column.go -type=User -field=Name -fieldType=string

package user

// User 用户
type User struct {
	ID   int
	Name string
}
