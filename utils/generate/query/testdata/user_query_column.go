// Please don't edit this file, it's made by go generate!

package user

// QueryColumn collect column from list
func QueryColumn(list []User) []string {
	columns := make([]string, 0)
	for _, s := range list {
		columns = append(columns, s.Name)
	}
	return columns
}
