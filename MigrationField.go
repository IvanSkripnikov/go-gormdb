package gormdb

import (
	logger "github.com/IvanSkripnikov/go-logger"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type MigrationField struct {
	FieldName string
	Type      int
}

type NewMigration struct {
	Model  schema.Tabler
	Fields []MigrationField
}

// Миграция на добавление полей и внешних связей.
func (thisModel *MigrationField) addColumnMigration(client *gorm.DB, model schema.Tabler) bool {
	var hasApplied bool

	if !client.Migrator().HasColumn(&model, thisModel.FieldName) {
		if errField := client.Migrator().AddColumn(&model, thisModel.FieldName); errField != nil {
			logger.Errorf("Failed to apply migration to add `%s` field to `%s` table. Error: %v",
				thisModel.FieldName, model.TableName(), errField)
		} else if errAuto := client.AutoMigrate(&model); errAuto != nil {
			logger.Errorf("Failed to apply auto migration: to %s table. Error: %v", model.TableName(), errAuto)
		} else {
			hasApplied = true
			logger.Infof("Migration to add the `%s` field to the `%s` table successfully applied.",
				thisModel.FieldName, model.TableName())
		}
	}

	return hasApplied
}
