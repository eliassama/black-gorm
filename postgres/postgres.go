package ormPostgres

import (
	ormInitialize "github.com/eliassama/black-gorm/extern/initialize"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

// New 创建 postgres 链接
//
//goland:noinspection ALL
func New(config *ormInitialize.DatabaseConf) (*gorm.DB, error) {
	var dsn strings.Builder

	dsn.WriteString("host=")
	dsn.WriteString(config.Host)

	dsn.WriteString(" user=")
	dsn.WriteString(config.User)

	dsn.WriteString(" password=")
	dsn.WriteString(config.Password)

	dsn.WriteString(" dbname=")
	dsn.WriteString(config.Database)

	dsn.WriteString(" port=")
	dsn.WriteString(strconv.FormatInt(config.Port, 10))

	dsn.WriteString(" sslmode=disable TimeZone=Asia/Shanghai")

	return ormInitialize.New(
		postgres.New(
			postgres.Config{DSN: dsn.String(), PreferSimpleProtocol: true},
		),
	)
}
