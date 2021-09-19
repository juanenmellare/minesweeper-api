package mocks

import (
	"minesweeper-api/configs"
	"os"
)

func setMockEnvs() {
	_ = os.Setenv("DATABASE_HOST", "FOO-HOST")
	_ = os.Setenv("DATABASE_USER", "FOO-USER")
	_ = os.Setenv("DATABASE_PASS", "FOO-PASS")
	_ = os.Setenv("DATABASE_PORT", "FOO-PORT")
	_ = os.Setenv("DATABASE_NAME", "FOO-NAME")
}

func Config() configs.Config {
	setMockEnvs()
	return configs.NewConfig()
}
