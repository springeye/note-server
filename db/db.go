package db

import (
	"fmt"
	"github.com/springeye/note-server/config"
	"time"

	//"gorm.io/driver/sqlite" // 基于 GGO 的 Sqlite 驱动
	"github.com/glebarez/sqlite" // 纯 Go 实现的 SQLite 驱动, 详情参考： https://github.com/glebarez/sqlite
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Connection *gorm.DB

type dbLogger struct {
}

func (d *dbLogger) Printf(msg string, args ...interface{}) {
	//slog.Info(msg, args...)
	fmt.Printf(msg, args...)
}
func Setup() {
	if Connection != nil {
		return
	}
	var err error
	//Connection, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	c := logger.Config{
		SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值
		LogLevel:                  logger.Info,            // 日志级别
		IgnoreRecordNotFoundError: true,                   // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  true,                   // 禁用彩色打印
	}
	if config.DefaultConfig.Debug {

	} else {
		c.LogLevel = logger.Error
	}
	newLogger := logger.New(&dbLogger{}, c)

	config := gorm.Config{Logger: newLogger}
	Connection, err = gorm.Open(sqlite.Open(":memory:?_pragma=foreign_keys(1)"), &config)
	if err != nil {
		panic(err)
	}
	user := User{}
	note := Note{}
	err = Connection.AutoMigrate(&user, &note)
	if err != nil {
		panic(err)
	}
	for i := 0; i < 3; i++ {
		Connection.Create(&User{Username: fmt.Sprintf("testuser%d", i)})
	}
}
