//go:generate go run ../map.go -type=User -field=Name -fieldType=string

package user

// User 用户
type User struct {
	Name string
}
