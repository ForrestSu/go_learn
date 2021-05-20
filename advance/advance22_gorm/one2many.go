package main

import (
	"log"

	"github.com/ForrestSu/go_learn/advance/advance22_gorm/dao"
	"github.com/kylelemons/godebug/pretty"
	"gorm.io/gorm"
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

	var db = dao.GetDB()
	// Migrate the schema
	db.AutoMigrate(&MyUser{}, &CreditCard{})

	// db.Create(&User{Name: "张三", CreditCards: []CreditCard{{TagName: "男"}, {TagName: "90后"}}})
	// db.Create(&User{Name: "李四", CreditCards: []CreditCard{{TagName: "男"}, {TagName: "90后"}}})
	var users []MyUser
	db.Find(&users)
	log.Println(pretty.Sprint(users))
}
