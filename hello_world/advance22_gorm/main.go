package main

import (
	"errors"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Product struct {
	gorm.Model
	Code   string
	Price  int
	Expire time.Time
	OpenID string `gorm:"not null; index; column:open_id; type:varchar(64);"`
	Sex    int8   `gorm:"comment:'注释'"`
	Text   string `gorm:"size:64"`
}

type User struct {
	gorm.Model
	Name      string `gorm:"size:64; comment:'用户名'"`
	OpenID    string `gorm:"size:64; index; comment:'微信OpenID'"`
	Birth     int32  `gorm:"comment:'出生年月YYYYMM'"`
	Sex       int8   `gorm:"comment:'性别'"`
	Marital   int8
	Education int8
	Job       int8
	Income    int8
	IsVip     bool   `gorm:"not null; default:false"` // 是否为腾讯视频会员
	QQ        string `gorm:"size:32"`
	WX        string `gorm:"size:32"`
	State     int8   `gorm:"comment:'有效状态'"`
	Tags      string `gorm:"comment:'用户标签ID逗号分隔'"`
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/survey_db?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: &schema.NamingStrategy{
			TablePrefix:   "t_",
			SingularTable: true,
		},
	})
	if err != nil {
		log.Println(err)
		return
	}

	// Migrate the schema
	db.AutoMigrate(&Product{}, &User{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	result := db.First(&product, 1) // find product with integer primary key

	_ = result.RowsAffected // returns found records count
	_ = result.Error        // returns error
	log.Printf("rows = %v, err = %+v \n", result.RowsAffected, result.Error)
	// check error ErrRecordNotFound
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		log.Printf("record not exist")
	}

	db.First(&product, "code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)

}
