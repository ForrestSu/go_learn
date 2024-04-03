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
set_name: ad.sz.* 
map_data: 
  2:
    title: 拾光-接单通知
    content: 帐号（XXX）收到任务名（XXX）的邀约，请于24小时内确认接受，超时未接单订单会自动失效
  3:
    title: 拾光-上传脚本通知
    content: 已确认接受任务（XXX），请及时上传脚本
# TV 合作渠道免广
# TV 合作渠道免广
tv_channel:
  16252: # 长虹电视
    ad_types: ["TV", "TL"]     # TL TV前贴 TH TV后贴
    threshold: 3m30s   # 3 minutess
# 频道页对应的焦点图偏移 (精选第 3 帧)
page_id_offset: # 频道页对应的焦点图偏移 (精选第 3 帧)
  "100101": 2
`

type YamlConfig struct {
	// 整形数组
	ListInt []int `yaml:"list_int"`
	// 字符串数组
	ListStr []string `yaml:"list_str"`
	// duration
	Timeout time.Duration `yaml:"timeout"`
	// 字符串
	SetName string `yaml:"set_name"`
	// Map
	MapData map[int32]*Template `yaml:"map_data"`
	// Map duration
	MapDuration map[int]TVManufacture `yaml:"tv_channel"`
	// 频道页对应的焦点图偏移 (精选第 3 帧)
	PageIDOffset map[string]int32 `yaml:"page_id_offset"`
}

// TVManufacture 合作厂商
type TVManufacture struct {
	// 免广的广告类型 (默认只有前贴)
	AdTypes []string `yaml:"ad_types"`
	// 免广的视频时长
	Threshold time.Duration `yaml:"threshold"`
}

// 是否免广 (视频时长+广告类型)
func (m *TVManufacture) freeAd(d time.Duration) bool {
	return d > 0 && d <= m.Threshold
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
	cfg.Timeout = 5 * time.Minute
	err := yaml.Unmarshal([]byte(configInfo), cfg)
	assert.Nil(t, err)
	assert.Equal(t, "ad.sz.*", cfg.SetName)

	t.Log(pretty.Sprint(cfg))
	assert.Equal(t, 4*time.Minute, cfg.Timeout)

	val := cfg.MapDuration[16252]
	assert.Equal(t, 3*time.Minute+30*time.Second, val.Threshold)
	// assert.True(t, val.freeAd(3*time.Minute))
	// assert.False(t, val.freeAd(4*time.Minute))
	assert.Equal(t, int32(2), cfg.PageIDOffset["100101"])
}
