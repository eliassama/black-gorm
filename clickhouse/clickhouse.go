package ormClickhouse

import (
	ormInitialize "github.com/eliassama/black-gorm/extern/initialize"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"strconv"
	"strings"
)

// New 创建 clickhouse 链接
//
//goland:noinspection ALL
func New(config *ormInitialize.DatabaseConf) (*gorm.DB, error) {

	var dsn strings.Builder
	dsn.WriteString("tcp://")
	dsn.WriteString(config.Host)
	dsn.WriteString(":")
	dsn.WriteString(strconv.FormatInt(config.Port, 10))

	dsn.WriteString("?database=")
	dsn.WriteString(config.Database)

	dsn.WriteString("&username=")
	dsn.WriteString(config.User)

	dsn.WriteString("&password=")
	dsn.WriteString(config.Password)

	return ormInitialize.New(clickhouse.Open(dsn.String()))
}
