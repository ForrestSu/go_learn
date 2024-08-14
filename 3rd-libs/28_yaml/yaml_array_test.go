package yaml

import (
	"testing"

	"github.com/kylelemons/godebug/pretty"
	"gopkg.in/yaml.v3"
)

// 单行不是必须的，但你可以像下面这样换行，也可以使用多行进行等效表示
var yamlText = `
name: "test"
str_list: [1,
  2, 3
  ]
multi_str: 
 - x1
 - y1
 - z1
`

type dataMode struct {
	Name     string   `yaml:"name"`
	StrList  []string `yaml:"str_list"`
	MultiStr []string `yaml:"multi_str"`
}

func TestYamlArray(t *testing.T) {
	var d dataMode
	err := yaml.Unmarshal([]byte(yamlText), &d)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("d=", pretty.Sprint(d))
}
