package test_go_oo

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) speak() {
	fmt.Println("pet...")
}
func (p *Pet) speakTo(name string) {
	p.speak()
	fmt.Println("pet call", name)
}

//匿名嵌套类型
type Cat struct {
	Pet
	//p *Pet
}

func (c *Cat) speak() {
	fmt.Println("Wang")
}

func TestFn(t *testing.T) {
	pet := new(Pet)
	pet.speak()
	pet.speakTo("Alice")

	//c := (*Pet)(new(Cat))
	c := new(Cat)
	c.speak()
	// 并不能调用Cat的 speak 方法，这里并不能达到继承多态的效果
	c.speakTo("Bob")
	t.Log("exit")
}
