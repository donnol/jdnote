package config

var normal = Config{
	Server: Server{
		Port: 8810,
	},
	DB: DB{
		Scheme:   "postgres",
		Host:     "127.0.0.1",
		Port:     5432,
		User:     "jd",
		Password: "jd",
		Name:     "cicada",
	},
	JWT: JWT{
		Secret: "Xadfdfoere23242l2afasf34wraf090uadfrfdIEJF039039",
	},
}
