package main

import (
	"log"

	"github.com/kylelemons/godebug/pretty"
	"gorm.io/gorm"

	"github.com/ForrestSu/go_learn/advance/advance22_gorm/dao"
)

// 多对多-关系

// User 用户
type User struct {
	ID        uint        `gorm:"primary_key"`
	Name      string      `gorm:"not null"`
	Languages []*Language `gorm:"many2many:user_languages;"`
}

type Category struct {
	ID   int32
	Name string
}

// Language 语言
type Language struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"not null"`
	// belongs to
	CategoryID int32
	Category   *Category
	// has many
	Users []*User `gorm:"many2many:user_languages;"`
}

func insert(db *gorm.DB) {

	// Migrate the schema
	db.AutoMigrate(&User{}, &Language{})

	kind := &Category{ID: 1, Name: "语种"}
	var langCN = Language{Name: "汉语", CategoryID: kind.ID, Category: kind}
	db.Create(&langCN)
	var langEN = Language{Name: "English", CategoryID: kind.ID, Category: kind}
	db.Create(&langEN)

	// user
	db.Create(&User{Name: "张三", Languages: []*Language{{ID: 1}, {ID: 2}}})
	db.Create(&User{Name: "李四", Languages: []*Language{{ID: 2}}})
}

// QueryLinks ...
func QueryLinks(db *gorm.DB) {
	// 查询 id = 1 的用户信息
	var user1 User
	db.Find(&user1, 1)
	log.Println(pretty.Sprint(user1))

	// 关联的关键代码
	var userAssociation = db.Model(&user1).Association("Languages")

	// 替换所有关联 (为空则表示删除该用户的所有关联)
	userAssociation.Replace([]Language{{ID: 3}, {ID: 4}, {ID: 5}})

	// 再次查询关联
	var langs []Language
	db.Model(&user1).Association("Languages").Find(&langs)
	log.Println("langs = \n", pretty.Sprint(langs))

	// log.Println(pretty.Sprint(user1))

}

func QueryAll(db *gorm.DB) {
	// find all
	log.Println("== query all ===")
	// 查出所有用户，以及他们会的语言
	var users []User
	db.Preload("Languages", "category_id = ?", 1).Preload("Languages.Category").Find(&users)
	log.Println(pretty.Sprint(users))

	log.Println("delete ....")
	if len(users[0].Languages) > 0 {
		err := db.Model(&users[0]).Association("Languages").Delete(users[0].Languages)
		log.Println("err == ", err)
	}
	log.Println("will append ....")
	err := db.Model(&users[0]).Omit("Languages.*").
		Association("Languages").
		Append(&Language{ID: 1}, &Language{ID: 4}, &Language{ID: 5})
	log.Println("append OK", err)

	// 查出所有语言，以及下面的用户
	var lans []Language
	db.Preload("Users").Find(&lans)
	log.Println(pretty.Sprint(lans))
}

func main() {
	var db = dao.GetDB()
	// insert(db)
	QueryAll(db)
	// QueryLinks(db)
}
