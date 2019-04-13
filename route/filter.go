package route

// Filter 过滤器
type Filter interface {
	Filter() interface{}
}
