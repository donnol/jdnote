package app

import "github.com/donnol/jdnote/utils/config"

var normal = config.Config{
	Server: config.Server{
		Port: 8810,
	},
	Pprof: config.Pprof{
		Server: config.Server{
			Port: 6060,
		},
	},
	Prometheus: config.Prometheus{
		Server: config.Server{
			Port: 6660,
		},
	},
	DB: config.DB{
		Scheme:   "postgres",
		Host:     "127.0.0.1",
		Port:     5432,
		User:     "jd",
		Password: "jd",
		Name:     "cicada",
	},
	JWT: config.JWT{
		Secret: "Xadfdfoere23242l2afasf34wraf090uadfrfdIEJF039039",
	},
}
