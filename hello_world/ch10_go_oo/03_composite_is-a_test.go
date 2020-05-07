package test_go_oo

import (
	"fmt"
	"testing"
)

type Animal struct {
}
func (a *Animal) speak() {
	fmt.Println("animal...")
}
func (a *Animal) speakTo(name string) {
	a.speak()
	fmt.Println("animal call", name)
}

type Dog struct {
	Animal //is-a 匿名嵌套类型
	//p *Pet
}
//隐藏父类同名方法
func (d *Dog) speak() {
	fmt.Println("Wang")
}

func TestFn(t *testing.T) {
	a := new(Animal)
	a.speak()
	a.speakTo("Alice")

	dog := new(Dog)
	dog.speak()
	// 父类并不能调用子类Dog的speak方法，这里并不能达到继承多态的效果
	dog.speakTo("Bob")
	t.Log("exit")
}
