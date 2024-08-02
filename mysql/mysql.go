package ormMysql

import (
	ormInitialize "github.com/eliassama/black-gorm/extern/initialize"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

// New 创建 mysql 链接
//
//goland:noinspection ALL
func New(config *ormInitialize.DatabaseConf) (*gorm.DB, error) {
	var dsn strings.Builder

	dsn.WriteString(config.User)
	dsn.WriteString(":")
	dsn.WriteString(config.Password)

	dsn.WriteString("@tcp(")
	dsn.WriteString(config.Host)
	dsn.WriteString(":")
	dsn.WriteString(strconv.FormatInt(config.Port, 10))
	dsn.WriteString(")/")

	dsn.WriteString(config.Database)
	dsn.WriteString("?charset=utf8&parseTime=True&loc=Local")

	return ormInitialize.New(
		mysql.New(mysql.Config{
			DSN:                       dsn.String(), // DSN data source name
			DefaultStringSize:         255,          // string 类型字段的默认长度
			DisableDatetimePrecision:  true,         // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
			DontSupportRenameIndex:    true,         // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
			SkipInitializeWithVersion: false,        // 根据当前 MySQL 版本自动配置
		}),
	)
}
