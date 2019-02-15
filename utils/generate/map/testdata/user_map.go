// Please don't edit this file, it's made by go generate!

package user

// UserSliceToMap convert slice to map
func UserSliceToMap(list []User) map[string]User {
	m := make(map[string]User)
	for _, s := range list {
		m[s.Name] = s
	}
	return m
}
