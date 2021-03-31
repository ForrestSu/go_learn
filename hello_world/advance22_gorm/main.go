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
	Name   string
	Price  int
	Expire time.Time
	OpenID string `gorm:"index; column:open_id; type:varchar(64);"`
}

func main() {
	// HelloWorldTest()
	// OneToManyTest()
	// LinksTest()
	TestInsertOmit()
}

func TestInsertOmit() {
	var db = dao.GetDB()
	// Migrate the schema
	db.AutoMigrate(&Product{})
	//var filter = func(db *gorm.DB) *gorm.DB {
	//	return db.Omit("text", "code")
	//}
	var omitFields []string
	if omitFields == nil {
		log.Println("omitFields is nil")
	}
	db.Omit(omitFields...).Create(&Product{Name: "sunquan", Price: 100})
}

func HelloWorldTest() {

	var db = dao.GetDB()
	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Name: "D42", Price: 100})

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
	db.Model(&product).Updates(Product{Price: 200, Name: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - delete product
	db.Delete(&product, 1)
}

//func main() {
//	// HelloWorldTest()
//	// OneToManyTest()
//	LinksTest()
//}
