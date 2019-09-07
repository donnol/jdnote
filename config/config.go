package config

import (
	"github.com/donnol/jdnote/utils/config"
)

// defaultConfig 默认配置
var defaultConfig config.Config

// Default 默认值
func Default() config.Config {
	return defaultConfig
}

func init() {
	defaultConfig = normal
}
