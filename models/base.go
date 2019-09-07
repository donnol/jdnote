package model

import (
	"github.com/donnol/jdnote/utils/store/db"
)

// Base 基底
type Base struct {
	*db.Base
}

// NewBase 新建
func NewBase() *Base {
	return &Base{
		Base: db.New(defaultDB),
	}
}

// DB DB接口
type DB = db.DB
