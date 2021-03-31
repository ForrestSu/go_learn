package dao

import (
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// db 数据库连对象
var db *gorm.DB = nil

var doOnce sync.Once

func init() {
	doOnce.Do(defaultInitMysql)
}

// defaultInitMysql  初始化mysql
func defaultInitMysql() {
	dsn := "root:123456@tcp(localhost:3306)/survey_db?charset=utf8&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: &schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	log.Println("init success")
}

// GetDB 获取数据库对象
func GetDB() *gorm.DB {
	return db
}
