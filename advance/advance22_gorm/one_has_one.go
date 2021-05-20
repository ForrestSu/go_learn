package main

import (
	"log"

	"github.com/ForrestSu/go_learn/advance/advance22_gorm/dao"
	"github.com/kylelemons/godebug/pretty"
	"gorm.io/gorm"
)

/**
以用户、角色为例，这里规定，
一个用户属且属于一个角色，而同一个角色可以赋予多个用户，因此用户与角色的关系是 n:1
*/

// User
type Users struct {
	gorm.Model
	Name   string
	RoleID uint
	Role   Role
}

// Role 用户角色
type Role struct {
	ID   uint `gorm:"primarykey"`
	Name string
}

func LinksTest() {

	var db = dao.GetDB()

	// Migrate the schema
	//db.AutoMigrate(&User{}, &Role{})

	//db.Create(&User{Name: "王五", Role: Role{Name: "终端用户"}})
	// db.Create(&User{Name: "李四", RoleID: 1})
	//var user1 User
	//db.Where("id = ?", "1").First(&user1)
	//
	//// 关联的关键代码
	//db.Model(&user1).Association("Role").Find(&user1.Role)
	//log.Println(pretty.Sprint(user1))

	// find all
	log.Println("=====")
	var users []Users
	db.Joins("Role").Find(&users)

	// db.Preload("Role").Find(&users)
	log.Println(pretty.Sprint(users))
}
