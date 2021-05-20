package main_test

import (
	"fmt"
	"io"
	"strconv"
	"testing"

	"gopkg.in/check.v1"
)

/*
func (s *SuiteType) SetUpSuite(c *C) - 在测试套件启动前执行一次
func (s *SuiteType) SetUpTest(c *C) - 在每个用例执行前执行一次
func (s *SuiteType) TearDownTest(c *C) - 在每个用例执行后执行一次
func (s *SuiteType) TearDownSuite(c *C) - - 在测试套件用例都执行完成后执行一次
*/

var a int = 1

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {
	check.TestingT(t)
}

type MySuite struct{}

var _ = check.Suite(&MySuite{})

func (s *MySuite) SetUpSuite(c *check.C) {
	str3 := "1) 初始化工作...."
	fmt.Println(str3)

}

func (s *MySuite) TearDownSuite(c *check.C) {
	str4 := "last) 最后清理工作"
	fmt.Println(str4)
}

// 每个用例执行前, 准备工作
func (s *MySuite) SetUpTest(c *check.C) {
	str1 := "第" + strconv.Itoa(a) + "条用例开始执行, ready"
	fmt.Println(str1)

}

// 每个用例执行后, 清理工作
func (s *MySuite) TearDownTest(c *check.C) {
	str2 := "第" + strconv.Itoa(a) + "条用例执行完成, clean"
	fmt.Println(str2)
	a = a + 1
}

func (s *MySuite) TestHelloWorld(c *check.C) {
	c.Assert(42, check.Equals, 42)
	c.Assert(io.ErrClosedPipe, check.ErrorMatches, "io: .*on closed pipe")
	c.Check(42, check.Equals, 42)
}

func (s *MySuite) TestHelloTerry(c *check.C) {
	c.Assert("terry", check.Equals, "terry")
	c.Assert(io.ErrClosedPipe, check.ErrorMatches, "io: .*on closed pipe")
	c.Check(42, check.Equals, 42)
}
