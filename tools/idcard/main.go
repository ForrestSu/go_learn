package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/ForrestSu/go_learn/utils"
)

// 校验位检查
var (
	Weights    = [...]int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	CheckCodes = "10X98765432"
)

func isValidIDCardCheckDigit(idCard string) error {
	// 检查长度
	if len(idCard) != 18 {
		return fmt.Errorf("ID Card 长度满足 18 位")
	}

	// 检查是否只包含数字和最后一位可能是 X/x
	if match, _ := regexp.MatchString(`^(\d{17})(\d|X)$`, idCard); !match {
		return fmt.Errorf("不满足正则规则检查")
	}

	// 提取生日部分
	birthYear, _ := strconv.Atoi(idCard[6:10])
	birthMonth, _ := strconv.Atoi(idCard[10:12])
	birthDay, _ := strconv.Atoi(idCard[12:14])

	// 检查生日是否有效
	birthDate := time.Date(birthYear, time.Month(birthMonth), birthDay, 0, 0, 0, 0, time.UTC)
	if birthDate.Year() != birthYear || int(birthDate.Month()) != birthMonth || birthDate.Day() != birthDay {
		return fmt.Errorf("出生年月日非法")
	}

	var sum int
	for i := 0; i < 17; i++ {
		num := int32(idCard[i]) - '0'
		sum += Weights[i] * int(num)
	}
	checkSum := CheckCodes[sum%11]
	if checkSum != idCard[17] {
		return fmt.Errorf("身份证校验非法! 期望:%c, 传入:%c", checkSum, idCard[17])
	}
	return nil
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Printf("Usage: %s <idcard1> <idcard2>...\n", os.Args[0])
		return
	}
	for _, idCard := range os.Args[1:] {
		if err := isValidIDCardCheckDigit(idCard); err != nil {
			fmt.Printf("身份证: %s => %s\n", idCard, utils.ErrPt.Sprint(err.Error()))
		} else {
			fmt.Printf("身份证: %s => %s\n", idCard, utils.TitlePt.Sprint("已通过"))
		}
	}
}
