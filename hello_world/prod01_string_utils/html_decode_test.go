package prod01_string_utils

import (
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
