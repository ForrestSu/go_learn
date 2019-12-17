package test_go_oo

import (
	"fmt"
	"testing"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestStruct(t *testing.T) {
	em := Employee{"1", "Alice", 26}
	em1 := Employee{Name: "Bob", Age: 25}

	em2 := new(Employee) //注意这里返回的引用/指针，相当于 e := &Employee{}
	em2.Id = "2"
	em2.Age = 30
	em2.Name = "Rose"

	t.Log(em)
	t.Log(em1)
	t.Log(em2)
	t.Logf("em %T", em)
	t.Logf("em3 %T", em2)
}

func (e Employee) toString() {
	fmt.Printf("Address &e = %p\n", &e)
	fmt.Printf("id:%s, name:%s, age:%d\n", e.Id, e.Name, e.Age)
}

//建议使用第二种,指针的方式，防止对象被拷贝
func (e *Employee) show() {
	fmt.Printf("&e = %p\n", e)
	fmt.Printf("id:%s, name:%s, age:%d\n", e.Id, e.Name, e.Age)
}

//Go测试对象行为
func TestAction(t *testing.T) {
	em := Employee{"1", "Alice", 26}
	fmt.Printf("&em = %p\n", &em)
	//em.toString()
	em.show()
}
