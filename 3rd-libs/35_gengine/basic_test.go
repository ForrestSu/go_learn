package test

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/bilibili/gengine/builder"
	"github.com/bilibili/gengine/context"
	"github.com/bilibili/gengine/engine"
)

// 定义想要注入的结构体
type User struct {
	Name string
	Age  int64
	Male bool
}

func (u *User) GetNum(i int64) int64 {
	return i
}

func (u *User) Print(s string) {
	fmt.Println(s)
}

func (u *User) Say() {
	fmt.Println("hello world")
}

// 定义规则
const rule1 = `
rule "name test" "i can"  salience 0
begin
		if 7 == User.GetNum(7){
			User.Age = User.GetNum(89767) + 10000000
			// User.Print("6666")
		}else{
			User.Name = "yyyy"
		}
end
`

func Test_Multi(t *testing.T) {
	user := &User{
		Name: "Calo",
		Age:  0,
		Male: true,
	}

	dataContext := context.NewDataContext()
	// 注入初始化的结构体
	dataContext.Add("User", user)

	// init rule engine
	ruleBuilder := builder.NewRuleBuilder(dataContext)

	start1 := time.Now()
	// 构建规则
	err := ruleBuilder.BuildRuleFromString(rule1) // string(bs)
	log.Printf("rules num:%d, load rules cost time:%v", len(ruleBuilder.Kc.RuleEntities), time.Since(start1))
	if err != nil {
		panic(err)
	}

	eng := engine.NewGengine()

	start := time.Now()
	// 执行规则
	err = eng.Execute(ruleBuilder, true)
	println(user.Age)
	if err != nil {
		log.Printf("execute rule error: %v", err)
	}
	log.Printf("execute rule cost %v \n", time.Since(start))
	log.Printf("user.Age=%d,Name=%s,Male=%t \n", user.Age, user.Name, user.Male)
}
