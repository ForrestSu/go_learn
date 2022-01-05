package utils

import (
	"bufio"
	"io"
	"strings"
)

// ReadFileToLines 按行读取文件内容
// ignoreEmpty - ignore empty line if is true
func ReadFileToLines(r io.Reader, ignoreEmpty bool) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		cur := strings.TrimSpace(scanner.Text())
		if ignoreEmpty && (cur == "" || cur == "1") {
			continue
		}
		lines = append(lines, cur)
	}
	return lines, scanner.Err()
}
