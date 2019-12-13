package try_test

import (
	"testing"
)

func TestMapInit(t *testing.T) {
	//01 way
	m1 := map[int]int{1: 100, 2: 200}
	// map[1:100 2:200] 2
	t.Log(m1, len(m1))

	// 02 way
	m2 := map[int]int{}
	m2[10] = 100
	//map[10:100] 1
	t.Log(m2, len(m2))

	// 03 way
	m3 := make(map[int]int, 10)
	//map[] len = 0
	t.Log(m3, len(m3))
}

/**(1) 和 c++的map不一样, c++直接访问会增加一个kv,
 * (2) 而Go里面的map不会, 直接访问不存在的 key 会返回0值；
 * (3) 也就是说,通过第一个返回值 == 0, 并不能真正确认key是否存在。
 */
func TestAccessNoneExistKey(t *testing.T) {
	m1 := map[int]int{1: 100, 2: 200}
	// m1[10] 将会输出 0
	t.Log(m1[10])

	// m1[10] = 0
	// t.Log(m1[10])

	/** Go 判断一个key,需要判断第二个值为true 还是false
	  m1[10]会返回2个值, ok为true，表示key存在
	*/
	if v, ok := m1[10]; ok {
		t.Log("key is exist!", v)
	} else {
		t.Log("key is not exist!")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[string]string{"apple": "100", "orange": "200"}
	for k, v := range m1 {
		t.Logf("%s = %s", k, v)
	}
}
