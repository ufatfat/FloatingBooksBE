package db

import (
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"FloatingBooks/config"
	"time"
)

var Redis *redis.Client
var Mysql *gorm.DB
var err error

func DBInit () {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,         // Disable color
		},
	)
	Redis = redis.NewClient(&redis.Options{
		Addr: config.REDIS_ADDR,
		Password: config.REDIS_PASS,
		DB: config.REDIS_DB,
	})
	Mysql, err = gorm.Open(mysql.Open(config.MYSQL_CONN_STR), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}