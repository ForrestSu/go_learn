package main

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestQuery(t *testing.T) {
	var db *sql.DB
	db, _ = sql.Open("mysql", "root:mysql.admin.pass@tcp(192.168.0.216:3306)/test?charset=utf8")
	defer db.Close()
	db.SetMaxOpenConns(1000)
	err := db.Ping()
	if err != nil {
		fmt.Println("Failed to connect to mysql,err:" + err.Error())
		os.Exit(1)
	}
	// 操作一：执行数据操作语句

	var insertSqlStr = "insert into students values (5, 'berry', 'berry@gmail.com')"
	result, err := db.Exec(insertSqlStr) //执行SQL语句
	if err != nil {
		fmt.Println(err)
	} else {
		n, _ := result.RowsAffected() //获取受影响的记录数
		fmt.Println("受影响的记录数是", n)
	}

	// 操作二：执行预处理
	/*
	   stu:=[2][2] string{{"3","ketty"},{"4","rose"}}
	   stmt,_:=db.Prepare("insert into students values (?,?)")      //获取预处理语句对象
	   for _,s:=range stu{
	       stmt.Exec(s[0],s[1])            //调用预处理语句
	   }
	*/

	//操作三：单行查询
	/*
	   var id,name string
	   rows:=db.QueryRow("select * from students where id=4")   //获取一行数据
	   rows.Scan(&id,&name)        //将rows中的数据存到id,name中
	   fmt.Println(id,"--",name)
	*/

	// 操作四：多行查询
	rows, _ := db.Query("select * from students") //获取所有数据
	var id int64
	var name, email string
	for rows.Next() { //循环显示所有的数据
		err = rows.Scan(&id, &name, &email)
		fmt.Printf("id = %d, name = %s, email = %s.\n", id, name, email)
	}
}

/**
goos: windows
goarch: amd64
BenchmarkQuery
BenchmarkQuery-12           2844        402803 ns/op        420 B/op       9 allocs/op
PASS
ok      command-line-arguments  1.448s
*/
func BenchmarkQuery(b *testing.B) {
	var db *sql.DB
	db, _ = sql.Open("mysql", "root:mysql.admin.pass@tcp(192.168.0.216:3306)/test?charset=utf8")
	db.SetMaxOpenConns(1000)
	defer db.Close()
	// start test
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rows, _ := db.Query("select * from students") //获取所有数据
		rows.Close()
	}
	b.StopTimer()
}
