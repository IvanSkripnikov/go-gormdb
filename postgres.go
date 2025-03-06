package gormdb

import (
	"fmt"

	logger "github.com/IvanSkripnikov/go-logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// InitPostgres Инициализация подключения к Postgres.
func InitPostgres(config Database) (*gorm.DB, error) {
	return AddPostgres(GetDefaultClientName(), config)
}

// AddPostgres Добавить подключение к Postgres.
func AddPostgres(clientName string, config Database) (*gorm.DB, error) {
	dataSource := config.getPostgresDataSource()
	client, err := gorm.Open(postgres.Open(dataSource), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})
	if err != nil {
		failedOpenConnectionOnDB.Inc()
		return nil, err
	}

	AddClient(clientName, client)
	openConnectionOnDB.Inc()

	logger.Infof("Postgres initialized (db: %s)", config.DB)

	return client, nil
}

// Получить строку соединения для БД Postgres.
func (config *Database) getPostgresDataSource() string {
	config.checkDbPort()

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.Address, config.User, config.Password, config.DB, config.Port)
}
