package pipe_filter

import (
	"testing"
)

func TestStraightPipeline(t *testing.T) {
	// 多个过滤器
	splitter := NewSplitFilter(",")
	converter := NewToIntFilter()
	sum := NewSumFilter()

	// 组合多个过滤器
	sp := NewStraightPipeline("filter-demo-v1", splitter, converter, sum)
	ret, err := sp.Process("1,2,3")
	if err != nil {
		t.Fatal(err)
	}

	if ret != 6 {
		t.Fatalf("The expected is 6, but the actual is %d", ret)
	}
}
