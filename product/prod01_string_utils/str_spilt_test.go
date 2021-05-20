package prod01_string_utils

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Str2PidInfo(value string) (pid int64, pidType, chargeType int) {
	parts := strings.Split(value, "_")
	if len(parts) == 3 {
		v1, err1 := strconv.ParseInt(parts[0], 10, 64)
		v2, err2 := strconv.Atoi(parts[1])
		v3, err3 := strconv.Atoi(parts[2])
		if err1 == nil && err2 == nil && err3 == nil {
			return v1, v2, v3
		}
	}
	return
}

func TestSplit(t *testing.T) {

	value := "100_101_1,02"
	pid, pidType, chargeType := Str2PidInfo(value)
	fmt.Printf("pid=%d, pidType=%d, chargeType=%d\n", pid, pidType, chargeType)
}
