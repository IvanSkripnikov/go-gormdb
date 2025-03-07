### Пакет реализующий функционал для работы с базами данных с использованием ORM Gorm


Использование пакета в проекте


Настройки подключения к базе данных


Для передачи настроек подключения используется структура gormdb.Database

```
config := gormdb.Database{
	Address:  os.Getenv("DB_ADDRESS"),
	Port:     os.Getenv("DB_PORT"),
	User:     os.Getenv("DB_USER"),
	Password: os.Getenv("DB_PASSWORD"),
	DB:       os.Getenv("DB_NAME"),
}
```

### Примеры подключения к базам данных


Для добавления подключения необходимо воспользоваться следующими функциями:

MS SQL Server
```
client, err := gormdb.AddSqlServer("[ClientName]", config);
```

MySQL
```
client, err := gormdb.AddMysql("[ClientName]", config);
```

Postgres
```
client, err := gormdb.AddPostgres("[ClientName]", config);
```

### Получение нужного подключения по наименованию


Все установленные подключения хранятся в глобальной переменной gormdb.Clients

Для получения нужного подключения можно использовать функцию gormdb.GetClient("[ClientName]")

*Пример:*
```
db := gormdb.GetClient("MySQL")
err := db.Where(...).Find(...).Error
```

### Упрощенная работа


В случае если в проекте используется только одна база данных, можно воспользоваться следующими функциями:

MS SQL Server
```
client, err := gormdb.InitSqlServer(config);
```

MySQL
```
client, err := gormdb.InitMysql(config);
```

Postgres
```
client, err := gormdb.InitPostgres(config);
```

Получение текущего подключения


Получить текущее подключение можно из глобальной переменной gormdb.Client

*Пример:*
```
db := gormdb.Client
err := db.Where(...).Find(...).Error
```