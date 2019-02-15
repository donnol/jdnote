package config

import "testing"

func TestMakeConfigFromFile(t *testing.T) {
	conf := MakeConfigFromFile("config.toml")
	t.Logf("%+v\n", conf)
	t.Logf("%+v\n", conf.DB.String())
}
