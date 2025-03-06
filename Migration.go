package gormdb

import (
	"fmt"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Migration struct {
	Version   string `gorm:"primarykey"`
	ApplyTime time.Time
}

func (model Migration) TableName() string {
	return "migration"
}

// Добавить сообщение для создния таблиц.
func createTableMigrationMessage(client *gorm.DB, model schema.Tabler) {
	message := fmt.Sprintf("%s %s", MigrationTypes[CreateTable], model.TableName())
	migration := Migration{
		Version:   message,
		ApplyTime: time.Now(),
	}

	client.Create(&migration)
}

// Добавить сообщение для добавления полей.
func addColumnMigrationMessage(client *gorm.DB, model schema.Tabler, fieldName string) {
	message := fmt.Sprintf("%s %s to %s table", MigrationTypes[AddColumn], fieldName, model.TableName())
	migration := Migration{
		Version:   message,
		ApplyTime: time.Now(),
	}

	client.Create(&migration)
}
