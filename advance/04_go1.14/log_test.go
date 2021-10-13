package go1_14

import (
	"fmt"
	"testing"
)

func TestCleanUp(t *testing.T) {
	//测试完成,最后清理一些资源
	t.Cleanup(func() {
		fmt.Print("clean up do something...")
	})
	t.Log("test")
	//t.Fatal("fatal...")
	t.Log("as")
}
