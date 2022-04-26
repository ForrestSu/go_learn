package main

import (
	"fmt"
	"log"
)

type VisitorFunc func(*Info, error) error

// Visitor
// 这个模式是一种将算法与操作对象的结构分离的一种方法。
// 这种分离的实际结果是能够在不修改结构的情况下向现有对象结构添加新操作，是遵循开放/封闭原则的一种方法
type Visitor interface {
	Visit(VisitorFunc) error
}

// Info 数据结构
type Info struct {
	Namespace   string
	Name        string
	OtherThings string
}

func (info *Info) Visit(fn VisitorFunc) error {
	return fn(info, nil)
}

// NameVisitor
// 声明了一个 NameVisitor 的结构体，这个结构体里有一个 Visitor 接口成员，这里意味着多态。
// 在实现 Visit() 方法时，其调用了自己结构体内的那个 Visitor的 Visitor() 方法:
//   这其实是一种修饰器的模式，用另一个Visitor修饰了自己
type NameVisitor struct {
	visitor Visitor
}

// Visit implement Visitor
func (v *NameVisitor) Visit(fn VisitorFunc) error {
	inner := func(info *Info, err error) error {
		fmt.Println("===> NameVisitor() before")
		err = fn(info, err)
		if err == nil {
			fmt.Printf("Name=%s, NameSpace=%s\n", info.Name, info.Namespace)
		}
		fmt.Println("<=== NameVisitor() after")
		return err
	}
	return v.visitor.Visit(inner)
}

type OtherThingsVisitor struct {
	visitor Visitor
}

func (v *OtherThingsVisitor) Visit(fn VisitorFunc) error {
	inner := func(info *Info, err error) error {
		fmt.Println("===> OtherThingsVisitor() before")
		err = fn(info, err)
		if err == nil {
			fmt.Println("OtherThings=", info.OtherThings)
		}
		fmt.Println("<=== OtherThingsVisitor() after")
		return err
	}
	return v.visitor.Visit(inner)
}

// 实现简单的拦截器
func main() {
	var v Visitor = &Info{}
	v = &NameVisitor{v}
	v = &OtherThingsVisitor{v}
	innerFunc := func(info *Info, err error) error {
		info.Name = "Alice"
		info.Namespace = "Production"
		info.OtherThings = "hello world!"
		log.Println("innerFunc()")
		return nil
	}
	_ = v.Visit(innerFunc)
}

/* Output:
===> NameVisitor() before
===> OtherThingsVisitor() before
2022/04/19 16:01:39 innerFunc()
OtherThings = hello world!
<=== OtherThingsVisitor() after
Name=Alice, NameSpace=Production
<=== NameVisitor() after
*/
