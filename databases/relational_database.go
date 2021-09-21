package databases

import (
	"gorm.io/gorm"
	"minesweeper-api/models"
)

type RelationalDatabase interface {
	Get() *gorm.DB
	DoMigration()
}

type relationDatabaseImpl struct {
	database *gorm.DB
}

func NewConnection(database *gorm.DB, err error) RelationalDatabase {
	if err != nil {
		panic("[ERROR] there was an error while trying to connect database: " + err.Error())
	}

	return &relationDatabaseImpl{
		database: database,
	}
}

func (r relationDatabaseImpl) Get() *gorm.DB {
	return r.database
}

func (r relationDatabaseImpl) DoMigration() {
	migrator := r.Get().Migrator()

	modelsToAutoMigrate := []interface{}{
		&models.Settings{},
		&models.Field{},
		&models.Game{},
	}

	for _, model := range modelsToAutoMigrate {
		handleMigrationErr(migrator.AutoMigrate(model))
	}
}

func handleMigrationErr(err error) {
	if err != nil {
		panic(err.Error())
	}
}
