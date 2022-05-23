package dao

import (
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// db 数据库连对象
var dbConn *gorm.DB

var doOnce sync.Once

func init() {
	doOnce.Do(func() {
		dsn := "root:123456@tcp(localhost:3306)/testdb?charset=utf8&parseTime=True&loc=Local"
		db, err := NewMySQLDB(dsn)
		if err != nil {
			panic(err)
		}
		dbConn = db
		log.Println("init success")
	})
}

// GetDB 获取数据库对象
func GetDB() *gorm.DB {
	return dbConn
}

// NewMySQLDB 创建 mysql 连接
func NewMySQLDB(dsn string) (*gorm.DB, error) {
	return gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: &schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
}

// MustSQLiteDB 适用于本地测试
func MustSQLiteDB(dst ...interface{}) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err = db.AutoMigrate(dst...); err != nil {
		panic(err)
	}
	return db
}
