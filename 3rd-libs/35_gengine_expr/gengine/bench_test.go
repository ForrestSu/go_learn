package gengine

import (
	"fmt"
	"testing"

	"github.com/bilibili/gengine/builder"
	"github.com/bilibili/gengine/context"

	"github.com/bilibili/gengine/engine"
)

const source = `
rule "ott_rule"
begin
  // 创维酷开 2022-06-10 
  if ChID == 303 {
	if in("TL", AdTypes) && Vid == "y00420wuzxh" && time() <= 1654790403 {
	  return 44
	}
  }
  return 0
end`

type MyService struct {
	Pool *engine.GenginePool
}

// 初始化业务服务
func NewMyService(poolMinLen, poolMaxLen int64, em int, rulesStr string) *MyService {
	pool, e := engine.NewGenginePool(poolMinLen, poolMaxLen, em, rulesStr, nil)
	if e != nil {
		panic(fmt.Sprintf("初始化gengine失败，err:%+v", e))
	}
	return &MyService{Pool: pool}
}
func BenchmarkExecSelected(b *testing.B) {
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
	ruleBuilder.BuildRuleFromString(rule1)

	eng := engine.NewGengine()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		eng.Execute(ruleBuilder, true)
	}
}
