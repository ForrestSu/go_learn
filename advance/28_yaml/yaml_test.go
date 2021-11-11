package yaml_test

import (
	"testing"
	"time"

	"github.com/kylelemons/godebug/pretty"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

const configInfo = `
list_int: [1, 2, 3]
list_str: 
  - a
  - "text"
  - 123
timeout: 4m  # 4 minutes
map_data: 
  2:
    title: 拾光-接单通知
    content: 帐号（XXX）收到任务名（XXX）的邀约，请于24小时内确认接受，超时未接单订单会自动失效
  3:
    title: 拾光-上传脚本通知
    content: 已确认接受任务（XXX），请及时上传脚本
`

type YamlConfig struct {
	// 整形数组
	ListInt []int `yaml:"list_int"`
	// 字符串数组
	ListStr []string `yaml:"list_str"`
	// duration
	Timeout time.Duration `yaml:"timeout"`
	// Map
	MapData map[int32]*Template `yaml:"map_data"`
}

// Template 消息模板
type Template struct {
	Title   string `yaml:"title"`   // 标题
	Content string `yaml:"content"` // 内容
}

// yaml 解析time.Duration类型, 是直接使用的标准库函数 time.ParseDuration()
// 扩展 time.ParseDuration() 具体支持的时间单位如下：
// Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".

func TestYaml(t *testing.T) {
	var cfg = &YamlConfig{}
	err := yaml.Unmarshal([]byte(configInfo), cfg)
	assert.Nil(t, err)

	t.Log(pretty.Sprint(cfg))
	assert.Equal(t, 4*time.Minute, cfg.Timeout)
}
