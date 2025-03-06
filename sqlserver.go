package gormdb

import (
	"fmt"

	logger "github.com/IvanSkripnikov/go-logger"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// InitSqlServer Инициализация подключения к Sql Server.
func InitSqlServer(config Database) (*gorm.DB, error) {
	return AddSqlServer(GetDefaultClientName(), config)
}

// AddSqlServer Добавить подключение к Sql SeDrver.
func AddSqlServer(clientName string, config Database) (*gorm.DB, error) {
	dataSource := config.getMssqlDataSource()
	client, err := gorm.Open(sqlserver.Open(dataSource), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})
	if err != nil {
		failedOpenConnectionOnDB.Inc()
		return nil, err
	}

	AddClient(clientName, client)
	openConnectionOnDB.Inc()

	logger.Infof("SQL Server initialized (db: %s)", config.DB)

	return client, nil
}

// Получить строку соединения для БД Sql Server.
func (config *Database) getMssqlDataSource() string {
	config.checkDbPort()

	return fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		config.User, config.Password, config.Address, config.Port, config.DB)
}
