package gormdb

import (
	"fmt"

	logger "github.com/IvanSkripnikov/go-logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
)

// InitMysql Инициализация подключения к Mysql.
func InitMysql(config Database) (*gorm.DB, error) {
	return AddMysql(GetDefaultClientName(), config)
}

// AddMysql Добавить подключение к Mysql.
func AddMysql(clientName string, config Database) (*gorm.DB, error) {
	dataSource := config.getMysqlDataSource()
	client, err := gorm.Open(mysql.Open(dataSource), &gorm.Config{
		Logger: gormLogger.Default.LogMode(gormLogger.Silent),
	})
	if err != nil {
		failedOpenConnectionOnDB.Inc()
		return nil, err
	}

	AddClient(clientName, client)
	openConnectionOnDB.Inc()

	logger.Infof("MySQL initialized (db: %s)", config.DB)

	return client, nil
}

// Получить строку соединения для БД Mysql.
func (config *Database) getMysqlDataSource() string {
	config.checkDbPort()

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.User, config.Password, config.Address, config.Port, config.DB)
}
