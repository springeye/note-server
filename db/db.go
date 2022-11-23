package db

import (
	//"gorm.io/driver/sqlite" // 基于 GGO 的 Sqlite 驱动
	"github.com/glebarez/sqlite" // 纯 Go 实现的 SQLite 驱动, 详情参考： https://github.com/glebarez/sqlite
	"gorm.io/gorm"
)

var Connection *gorm.DB

func init() {
	var err error
	//Connection, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	Connection, err = gorm.Open(sqlite.Open(":memory:?_pragma=foreign_keys(1)"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	user := User{}
	note := Note{}
	err = Connection.AutoMigrate(&user, &note)
	if err != nil {
		panic(err)
	}
}
