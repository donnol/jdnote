package model

// CommonParam 通用参数
type CommonParam struct {
	BeginTime int64 `json:"beginTime"` // 开始时间
	EndTime   int64 `json:"endTime"`   // 结束时间

	PageIndex int `json:"pageIndex"` // 分页开始
	PageSize  int `json:"pageSize"`  // 分页大小
}

// DefaultCommonParam 默认值
var DefaultCommonParam = CommonParam{
	PageSize: 10,
}

// Result 结果
type Result struct {
	data interface{} // 数据
	err  error       // 错误
}

// SetData 设置数据
func (r *Result) SetData(data interface{}) *Result {
	r.data = data
	return r
}

// Data 获取数据
func (r *Result) Data() interface{} {
	return r.data
}

// SetErr 设置错误，如果已存在错误，则忽略
func (r *Result) SetErr(err ...error) *Result {
	if r.err != nil {
		return r
	}
	for _, e := range err {
		if e != nil {
			r.err = e
			break
		}
	}
	return r
}

// Err 获取错误
func (r *Result) Err() error {
	return r.err
}

// ErrIsNil 错误是否存在
func (r *Result) ErrIsNil() bool {
	return r.err == nil
}

// Unwrap 如果错误不为nil，则panic
// 由调用方决定要不要调Unwrap，而不是在方法里自己panic -- 将决定权留给调用方
func (r Result) Unwrap() interface{} {
	if r.err != nil {
		panic(r.err)
	}
	return r.data
}
