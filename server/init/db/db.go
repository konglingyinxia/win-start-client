package db

import (
	"fmt"
	"github.com/konglingyinxia/win-start-client/server/constant"
	"github.com/konglingyinxia/win-start-client/server/global"
	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite" // Sqlite driver based on CGO
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func Init() {
	fullPath := constant.HomePath + "/" + constant.DbFile
	if _, err := os.Stat(fullPath); err != nil {
		f, err := os.Create(fullPath)
		if err != nil {
			global.LOG.Fatal(fmt.Sprintf("sqlite-初始化数据库文件失败，错误: %v", err))
		}
		_ = f.Close()
	}
	newLogger := logger.New(log.Default(), logger.Config{})
	db, err := gorm.Open(sqlite.Open(fullPath), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   newLogger,
	})
	if err != nil {
		panic(err)
	}
	_ = db.Exec("PRAGMA journal_mode = WAL;")
	sqlDB, dbError := db.DB()
	if dbError != nil {
		panic(dbError)
	}
	sqlDB.SetConnMaxIdleTime(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	global.DB = db
	global.LOG.Info("sqlite-数据库初始化完成...")
}
