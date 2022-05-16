package app

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"sync"
	"time"
)

const (
	DSN = `root:zhangkaqiu520.@tcp(124.223.46.4:3306)/cats?parseTime=true`
)

var DB *gorm.DB
var Once = sync.Once{}

func NewDB() *gorm.DB {
	Once.Do(func() {
		var err error
		DB, err = gorm.Open(mysql.Open(DSN), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			log.Fatalln(err)
		}

		sqlDB, err := DB.DB()

		if sqlDB == nil || err != nil {
			log.Fatalln(err)
		}

		// SetMaxIdleConn 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(100)

		// SetMaxOpenConn 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(500)

		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Hour)
	})

	return DB
}
