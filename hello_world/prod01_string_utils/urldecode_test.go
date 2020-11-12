package prod01_string_utils

import (
	"fmt"
	"net/url"
	"testing"
)

func TestUrl(t *testing.T) {
	urlStr := "http://www.baidu.com/s?wd=自由度"
	fmt.Println(urlStr)

	encodeUrl := url.QueryEscape(urlStr)
	fmt.Println(encodeUrl)

	decodeUrl, err := url.QueryUnescape(encodeUrl)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(decodeUrl)
}
