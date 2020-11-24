package prod01_string_utils

import (
	"github.com/stretchr/testify/assert"
	"html"
	"testing"
)

func TestHtmlEncode(t *testing.T) {
	unescaped := `<script>alert(123);</script>`
	escaped := html.EscapeString(unescaped)
	t.Log(escaped)
}

func TestHtmlDecode(t *testing.T) {

	const s = `&quot;Fran &amp; Freddie&#39;s Diner&quot; &lt;tasty@example.com&gt;`
	unescaped  := html.UnescapeString(s)

	t.Log(unescaped)
}

func TestHtmlDecodeJson(t *testing.T){

	jsonstr := `{"app_ver":50202119,"cid":"rhjgk5bo0uxotgt","expansion":"","historyVid":"","history_duration":0,"isVip":0,"lid":"","outWebId":"","plat_bucketid":0,"pt":1,"sessionId":4,"vid":""}`

	val := html.EscapeString(jsonstr)
	unescaped  := html.UnescapeString(val)
	assert.Equal(t, jsonstr, unescaped)
	t.Log(val)
}