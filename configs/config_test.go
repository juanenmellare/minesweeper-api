package configs

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setMockEnvs() {
	_ = os.Setenv("DATABASE_HOST", "FOO-HOST")
	_ = os.Setenv("DATABASE_USER", "FOO-USER")
	_ = os.Setenv("DATABASE_PASS", "FOO-PASS")
	_ = os.Setenv("DATABASE_PORT", "FOO-PORT")
	_ = os.Setenv("DATABASE_NAME", "FOO-NAME")
}

func mockConfig() Config {
	setMockEnvs()
	return NewConfig()
}

var config = mockConfig()

func assertPanicMessage(t *testing.T, function func() Config, panicMessage string) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Not panic")
		} else {
			assert.Equal(t, r.(error).Error(), panicMessage)
		}
	}()
	function()
}

func assertEnvIsMissing(t *testing.T, env string) {
	setMockEnvs()
	_ = os.Unsetenv(env)

	assertPanicMessage(t, NewConfig, env+" doesn't exist.")
}

func TestConfigImpl_GetDatabaseHost(t *testing.T) {
	assert.Equal(t, config.GetDatabaseHost(), "FOO-HOST")
}

func TestConfigImpl_NewConfig_missing_DATABASE_HOST(t *testing.T) {
	assertEnvIsMissing(t, "DATABASE_HOST")
}

func TestConfigImpl_GetDatabaseUser(t *testing.T) {
	assert.Equal(t, config.GetDatabaseUser(), "FOO-USER")
}

func TestConfigImpl_NewConfig_missing_DATABASE_USER(t *testing.T) {
	assertEnvIsMissing(t, "DATABASE_USER")
}

func TestConfigImpl_GetDatabasePass(t *testing.T) {
	assert.Equal(t, config.GetDatabasePass(), "FOO-PASS")
}

func TestConfigImpl_NewConfig_missing_DATABASE_PASS(t *testing.T) {
	assertEnvIsMissing(t, "DATABASE_PASS")
}

func TestConfigImpl_GetDatabasePort(t *testing.T) {
	assert.Equal(t, config.GetDatabasePort(), "FOO-PORT")
}

func TestConfigImpl_NewConfig_missing_DATABASE_PORT(t *testing.T) {
	assertEnvIsMissing(t, "DATABASE_PORT")
}

func TestConfigImpl_GetDatabaseName(t *testing.T) {
	assert.Equal(t, config.GetDatabaseName(), "FOO-NAME")
}

func TestConfigImpl_NewConfig_missing_DATABASE_NAME(t *testing.T) {
	assertEnvIsMissing(t, "DATABASE_NAME")
}
