package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/bytedance/sonic"
)

func main() {
	fmt.Printf("当前 Go 版本: %s\n", runtime.Version())
	fmt.Println("开始 Sonic + Go 1.24+ Map 内存踩踏复现测试...")

	// 1. 在局部创建一个空的 Map
	m := make(map[string]interface{})

	// 2. 准备一段稍微复杂一点的 JSON，促使 sonic 分配较多的 key，
	// 增加覆盖 / 破坏 Map 内部控制字节 (ctrl bytes) 或桶指针 (buckets) 的概率。
	jsonStr := `{"k1": 1, "k2": "v2", "k3": {"sub": "data"}, "k4": [1,2,3], "k5": 99.9, "k6": "trigger_crash", "k7": true, "k8": null}`

	// 3. 使用 sonic 反序列化。
	// 【危险动作发生】：在这个瞬间，由于 sonic 版本与 Go 1.24 的 Swiss Table 布局不匹配，
	// sonic 底层的 JIT 汇编代码算错了指针偏移量，直接将解析出来的数据写到了 Map 的“系统控制区”，
	// 此时底层 bucket 指针已经被篡改，但因为还没有人去读它，所以程序还活着。
	err := sonic.ConfigFastest.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		fmt.Printf("Unmarshal error: %v\n", err)
		return
	}

	fmt.Println("--> sonic.Unmarshal 成功！(但底层内存已经被悄悄破坏)")
	fmt.Println("--> 准备遍历 Map 触发 Panic...")
	time.Sleep(1 * time.Second) // 停顿一下，增加戏剧效果

	// 4. 遍历 Map
	// 此时 Go 的运行时迭代器 internal/runtime/maps.(*Iter).Next 启动，
	// 尝试去读取那些已经被 sonic 写坏的底层指针，瞬间触发 nil pointer dereference。
	count := 0
	for k, v := range m {
		// 如果内存未被破坏，这里会正常打印
		fmt.Printf("读取到 Key: %s, Value: %v\n", k, v)
		count++
	}

	fmt.Printf("成功遍历了 %d 个元素（如果你能看到这句话，说明没有发生 Panic）\n", count)
}
