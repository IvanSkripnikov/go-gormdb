package gormdb

import (
	"strconv"

	logger "github.com/IvanSkripnikov/go-logger"
)

type Database struct {
	Address  string
	Port     string
	User     string
	Password string
	DB       string
}

// Проверка что порт серевера DB является числовым.
func (config *Database) checkDbPort() {
	if _, err := strconv.Atoi(config.Port); err != nil {
		logger.Fatalf("Failed to parse on Database port. Error: %v", err)
	}
}
