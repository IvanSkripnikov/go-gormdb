package gormdb

import (
	"strconv"
	"sync"

	logger "github.com/IvanSkripnikov/go-logger"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	mu            sync.RWMutex
	Client        *gorm.DB
	ClientName    string
	ClientCounter int
	Clients       map[string]*gorm.DB
)

func AddClient(name string, client *gorm.DB) {
	if name == "" {
		logger.Fatal("Client name cannot be empty string")
	}

	Client = client
	ClientName = name

	if Clients == nil {
		mu = sync.RWMutex{}
		Clients = make(map[string]*gorm.DB)
	}

	mu.Lock()
	Clients[name] = client
	mu.Unlock()
}

func GetClient(name string) *gorm.DB {
	mu.RLock()
	defer mu.RUnlock()

	if Clients == nil {
		logger.Fatal("No clients have been initialized")
	}

	if client, ok := Clients[name]; ok {
		return client
	}

	logger.Fatalf("Client '%s' has not been initialized", name)
	return nil
}

func GetDefaultClientName() string {
	mu.Lock()
	defer mu.Unlock()

	ClientCounter++
	return "Client" + strconv.Itoa(ClientCounter)
}

// CheckTables Проверить наличие таблиц в БД для текущего клиента.
func CheckTables(models ...schema.Tabler) {
	CheckTablesForClient(ClientName, models...)
}

// CheckTablesForClient Проверить наличие таблиц в БД.
func CheckTablesForClient(clientName string, models ...schema.Tabler) {
	var missingTables int

	client := GetClient(clientName)

	for _, model := range models {
		if !client.Migrator().HasTable(&model) {
			missingTables++
			logger.Errorf("Table `%s` is missing in the database.", model.TableName())
		}
	}

	if missingTables > 0 {
		logger.Fatal("The database is missing required tables")
	}
}

// ApplyMigrations Применить миграции для текущего клиента.
func ApplyMigrations(models ...schema.Tabler) {
	ApplyMigrationsForClient(ClientName, models...)
}

// ApplyMigrationsForClient Применить миграции.
func ApplyMigrationsForClient(clientName string, models ...schema.Tabler) {
	tablesModels := append([]schema.Tabler{Migration{}}, models...)

	client := GetClient(clientName)

	for _, model := range tablesModels {
		if createTable(client, model) {
			createTableMigrationMessage(client, model)
		}
	}
}

// ApplyAlterTablesMigrations Применить миграции на изменение существующих таблиц для текущего клиента.
func ApplyAlterTablesMigrations(migrations []NewMigration) {
	ApplyAlterTablesMigrationsForClient(ClientName, migrations)
}

// ApplyAlterTablesMigrationsForClient Применить миграции на изменение существующих таблиц.
func ApplyAlterTablesMigrationsForClient(clientName string, migrations []NewMigration) {
	client := GetClient(clientName)

	for _, migration := range migrations {
		for _, migrationField := range migration.Fields {
			if migrationField.Type == AddColumn && migrationField.addColumnMigration(client, migration.Model) {
				appliedMigrations.Inc()
				addColumnMigrationMessage(client, migration.Model, migrationField.FieldName)
			} else {
				failedApplyMigrations.Inc()
			}
		}
	}
}

// Применить миграции на создание новых таблиц.
func createTable(client *gorm.DB, model schema.Tabler) bool {
	var hasApplied bool

	if !client.Migrator().HasTable(&model) {
		if err := client.AutoMigrate(&model); err != nil {
			failedApplyMigrations.Inc()
			logger.Errorf("Failed to apply migration to create table `%s`. Error: %v", model.TableName(), err)
		} else {
			hasApplied = true
			appliedMigrations.Inc()
			logger.Infof("Migration to create `%s` table successfully applied.", model.TableName())
		}
	}

	return hasApplied
}
