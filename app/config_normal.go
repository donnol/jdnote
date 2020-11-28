package app

import "github.com/donnol/jdnote/utils/config"

var dev = func() config.Config {
	// 根据normal衍生而来，做部分修改
	newConfig := normal

	newConfig.Server.Port = 8890
	newConfig.Pprof.Server.Port = 6090
	newConfig.Prometheus.Server.Port = 6690
	newConfig.DB.Name = normal.DB.Name + "_dev"
	newConfig.JWT.Secret = "Xadfdfoere23242l2afasf34wraf090uadfrfdI123dfdfe1"

	return newConfig
}()

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
