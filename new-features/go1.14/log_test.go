package go1_14

import (
	"log"
	"testing"
)

func TestCleanUp(t *testing.T) {
	// 测试完成,最后清理一些资源
	t.Cleanup(func() {
		log.Println("clean up do something...")
	})
	t.Log("test")
	t.Fatal("fatal...")
	// unreachable code
	t.Log("as")
}
