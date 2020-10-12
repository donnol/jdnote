package exp

import (
	"log"
	"reflect"
)

// Model Model
type Model struct {
	Base
}

// Srv Srv
type Srv struct {
	Base

	AM Model
	BM Model
}

// Base Base
type Base struct {
	DB DB
}

// Begin Begin
func (b *Base) Begin() error {
	newDB := &db{
		isTx: true,
	}
	log.Printf("%p\n", newDB)
	b.DB = newDB
	return nil
}

// Init Init
func (b *Base) Init() error {
	log.Printf("%p\n", b)
	refType := reflect.TypeOf(b)
	log.Printf("%+v\n", refType) // 反射也拿不到外面的类型的信息

	return nil
}

// Data Data
type Data struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// DB DB
type DB interface {
	Get(id uint) (Data, error)
}

type db struct {
	isTx bool
}

// Get Get
func (db *db) Get(id uint) (Data, error) {
	return Data{ID: 1, Name: "jd"}, nil
}
