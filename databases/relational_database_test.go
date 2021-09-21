package databases

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"minesweeper-api/errors"
	"testing"
)

func TestNewConnection(t *testing.T) {
	connection := NewConnection(nil, nil)

	assert.Implements(t, (*RelationalDatabase)(nil), connection)
}

func TestNewConnection_error(t *testing.T) {
	dialector := postgres.New(postgres.Config{
		DSN:        "panic",
		DriverName: "postgres",
	})

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Not panic")
		}
	}()

	_ = NewConnection(gorm.Open(dialector, &gorm.Config{}))
}

func Test_handleMigrationErr(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("panic")
		}
	}()

	err := errors.NewError("panic")

	handleMigrationErr(err)
}

func Test_relationDatabaseImpl_DoMigration(t *testing.T) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Can´t recover...")
		}
	}()

	database := NewConnection(nil, nil)

	database.DoMigration()

	t.Errorf("fail: not panic")
}
