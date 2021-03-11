package main

import (
	"errors"
	"log"
	"time"

	"github.com/ForrestSu/go_learn/hello_world/advance22_gorm/dao"

	"gorm.io/gorm"
)

// Product struct
type Product struct {
	gorm.Model
	Code   string
	Price  int
	Expire time.Time
	OpenID string `gorm:"not null; index; column:open_id; type:varchar(64);"`
	Sex    int8   `gorm:"comment:'注释'"`
	Text   string `gorm:"size:64"`
}

func HelloWorldTest() {

	var db = dao.GetDB()
	// Migrate the schema
	db.AutoMigrate(&Product{})

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

//func main() {
//	// HelloWorldTest()
//	// OneToManyTest()
//	LinksTest()
//}
