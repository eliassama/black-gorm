package ormInitialize

import (
	"github.com/eliassama/black-zap/gormlogger"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

// DatabaseConf 数据库配置
type DatabaseConf struct {
	Host     string
	Port     int64
	User     string
	Password string
	Database string
}

func New(dialector gorm.Dialector, levels ...gormLogger.LogLevel) (*gorm.DB, error) {

	level := gormLogger.Info

	if levels != nil && len(levels) > 0 && levels[0] > 0 && levels[0] < 5 {
		level = levels[0]
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 使用单数表名，如果启用此选项，`User` 的表将是 `user`
		},
		Logger: gormlogger.New(level),
	})

	if err != nil {
		return nil, err
	}

	// 连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err = sqlDB.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
