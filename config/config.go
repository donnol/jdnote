package config

import (
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

// DefaultConfig 默认配置
var DefaultConfig Config

// Config 配置
type Config struct {
	DB DB `toml:"database"`
}

func init() {
	// FIXME: 怎么使用这个配置文件比较方便呢？
	// DefaultConfig = MakeConfigFromFile("../../config/config.toml")
	DefaultConfig = Config{
		DB: DB{
			Scheme:   "postgres",
			Host:     "127.0.0.1",
			Port:     5432,
			User:     "jd",
			Password: "jd",
			Name:     "cicada",
		},
	}
}

// DB 数据库配置
type DB struct {
	Scheme   string // 协议
	Host     string // 地址
	Port     int    // 端口
	User     string // 用户
	Password string // 密码
	Name     string // 数据库名
	SSLmode  bool   // ssl模式
}

// MakeConfigFromFile 新建DB配置
func MakeConfigFromFile(file string) Config {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var conf Config
	if _, err := toml.Decode(string(data), &conf); err != nil {
		panic(err)
	}

	return conf
}

// String 字符串
func (db *DB) String() string {
	var str string

	format := "%s://%s:%s@%s:%d/%s?sslmode=disable"
	str = fmt.Sprintf(format, db.Scheme, db.User, db.Password, db.Host, db.Port, db.Name)

	return str
}
