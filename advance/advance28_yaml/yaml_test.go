package advance28_yaml

import (
	"testing"

	"github.com/kylelemons/godebug/pretty"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

const configInfo = `
2:
    title: 拾光-接单通知
    content: 帐号（XXX）收到任务名（XXX）的邀约，请于24小时内确认接受，超时未接单订单会自动失效
3:
    title: 拾光-上传脚本通知
    content: 已确认接受任务（XXX），请及时上传脚本
`

// MsgTemplate 消息模板
type MsgTemplate struct {
	Title   string `json:"title" yaml:"title"`     // 标题
	Content string `json:"content" yaml:"content"` // 内容
	// Status  string `json:"status" yaml:"status"`   // 状态
}

func TestYaml(t *testing.T) {
	var cfg = make(map[int32]*MsgTemplate)
	err := yaml.Unmarshal([]byte(configInfo), cfg)
	assert.Nil(t, err)

	t.Log(pretty.Sprint(cfg))
}
