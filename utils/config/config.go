package config

import (
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
)

// Config 配置
type Config struct {
	// 服务器配置
	Server Server `toml:"server"`

	Pprof      Pprof      `toml:"pprof"`
	Prometheus Prometheus `toml:"prometheus"`

	// 数据库配置
	DB DB `toml:"database"`

	// jwt配置
	JWT JWT `toml:"jwt"`
}

type Pprof struct {
	Server
}

type Prometheus struct {
	Server
}

type Port int

func (port Port) Raw() int {
	return int(port)
}

func (port Port) String() string {
	return fmt.Sprintf("'%d'", port)
}

func (port Port) ToAddr() string {
	return fmt.Sprintf(":%d", port)
}

// Server 服务器配置
type Server struct {
	Port Port // 端口，如：8810
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

// JWT jwt配置
type JWT struct {
	Secret string
}

// MakeConfigFromFile 新建DB配置
func MakeConfigFromFile(file string) (Config, error) {
	var conf Config

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return conf, err
	}

	if _, err := toml.Decode(string(data), &conf); err != nil {
		return conf, err
	}

	return conf, nil
}

// String 字符串
func (db DB) String() string {
	var str string

	format := "%s://%s:%s@%s:%d/%s?sslmode=disable"
	str = fmt.Sprintf(format, db.Scheme, db.User, db.Password, db.Host, db.Port, db.Name)

	return str
}
