//go:generate go run ../enum.go -type=Color

package enum

// 枚举
const (
	Red    Color = 1
	Green  Color = 2
	Yellow Color = 3
)

// Color 颜色
type Color int
