package main

import (
	"github.com/kylelemons/godebug/pretty"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
)

// MyUser has many CreditCards, UserID is the foreign key
type MyUser struct {
	gorm.Model
	Name        string
	CreditCards []CreditCard `gorm:"foreignKey:UserID"`
}

type CreditCard struct {
	gorm.Model
	TagName string
	UserID  uint
}

func OneToManyTest() {

	dsn := "root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: &schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println(err)
		return
	}

	// Migrate the schema
	db.AutoMigrate(&MyUser{}, &CreditCard{})

	// db.Create(&User{Name: "张三", CreditCards: []CreditCard{{TagName: "男"}, {TagName: "90后"}}})
	// db.Create(&User{Name: "李四", CreditCards: []CreditCard{{TagName: "男"}, {TagName: "90后"}}})
	var users []MyUser
	db.Find(&users)
	log.Println(pretty.Sprint(users))
}
