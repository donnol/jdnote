package config

import "testing"

func TestMakeConfigFromFile(t *testing.T) {
	conf, err := MakeConfigFromFile("testdata/config.toml")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", conf)
	t.Logf("%+v\n", conf.DB.String())
}
