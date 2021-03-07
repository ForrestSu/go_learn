package main

import (
	"github.com/kylelemons/godebug/pretty"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
)

// 多对多-关系

// User 用户
type User struct {
	ID        uint        `gorm:"primary_key"`
	Name      string      `gorm:"not null"`
	Languages []*Language `gorm:"many2many:user_languages;"`
}

// Language 语言
type Language struct {
	ID    uint    `gorm:"primary_key"`
	Name  string  `gorm:"not null"`
	Users []*User `gorm:"many2many:user_languages;"`
}

func LinksTest() {
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
	//db.AutoMigrate(&User{}, &Language{})

	// db.Create(&Language{Name: "汉语"})
	// db.Create(&User{Name: "张三", Languages: []Language{{ID: 1}}})

	//var user1 User
	//db.Where("id = ?", "1").First(&user1)
	//
	//// 关联的关键代码
	//db.Model(&user1).Association("Role").Find(&user1.Role)
	//log.Println(pretty.Sprint(user1))

	// find all
	log.Println("=====")
	// 查出所有用户，以及他们会的语言
	var users []User
	db.Preload("Languages").Find(&users)
	log.Println(pretty.Sprint(users))

	// 查出所有语言，以及下面的用户
	var lans []Language
	db.Preload("Users").Find(&lans)
	log.Println(pretty.Sprint(lans))




}

func main() {
	LinksTest()
}
