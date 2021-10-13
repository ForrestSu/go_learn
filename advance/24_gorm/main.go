package main

import (
	"database/sql"
	"errors"
	"log"
	"reflect"

	"github.com/ForrestSu/go_learn/advance/24_gorm/dao"

	"github.com/kylelemons/godebug/pretty"
	"gorm.io/gorm"
)

// Product struct
type Product struct {
	gorm.Model
	Name    string
	Price   uint64
	Expired sql.NullTime
	OpenID  string `gorm:"index; column:open_id; type:varchar(64);"`
	QQ      int64
}

func main() {
	// HelloWorldTest()
	// OneToManyTest()
	// LinksTest()
	// TestInsertOmit()
	// TestUpdateDifferentType()
	TestSqlNullTime()
}

func TestSqlNullTime() {
	var db = dao.GetDB()
	// Migrate the schema
	db.AutoMigrate(&Product{})
	// var product = &Product{Name: "Alice", Price: 100, Expired: sql.NullTime{
	// 	Time:  time.Now(),
	// 	Valid: true,
	// }}
	// var err = db.Create(product).Error
	// log.Println(err)
	// {
	// 	var products []Product
	// 	var result = db.Where("expired < ?", time.Now()).Find(&products)
	// 	log.Println(result.RowsAffected)
	// 	log.Println(pretty.Sprint(products))
	// }
	{
		var products []Product
		var result = db.Where("(qq is NULL OR qq < ?)", 10000).
			Where("id > 0").
			Find(&products)
		log.Println("records = ", result.RowsAffected)
		log.Println(pretty.Sprint(products))
	}

}

type ScopeFunc func(db *gorm.DB) *gorm.DB

func TestUpdateDifferentType() {
	var db = dao.GetDB()
	// Migrate the schema
	db.AutoMigrate(&Product{})

	//var product = &Product{Name: "D42", Price: 100}
	// db.Create(product)
	//db.Model(&product).Update("Price", "100")
	// 注意：虽然sql支持为int类型的列，更新时使用string 类型的值；
	// 但还是建议
	var ret = &Product{}
	db.Scopes(FilterEqual("price", 0)).Find(ret)

}

func FilterEqual(column string, value interface{}) ScopeFunc {
	return func(db *gorm.DB) *gorm.DB {
		if reflect.ValueOf(value).IsZero() {
			return db // 不过滤
		}
		return db.Where(column+" = ? ", value)
	}
}

// TestInsertOmit 测试 Insert时，显式 omit字段
func TestInsertOmit() {
	var db = dao.GetDB()
	// Migrate the schema
	db.AutoMigrate(&Product{})
	//var filter = func(db *gorm.DB) *gorm.DB {
	//	return db.Omit("text", "code")
	//}
	var omitFields []string
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
