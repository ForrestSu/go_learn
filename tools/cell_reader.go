package main

import (
	"strings"

	"github.com/ForrestSu/go_learn/tools/utils"
	"github.com/xuri/excelize/v2"
)

type cellReader struct {
	Err       error
	file      *excelize.File
	sheetName string
}

func newCellReader(f *excelize.File, sheetName string) *cellReader {
	return &cellReader{
		file:      f,
		sheetName: sheetName,
	}
}

// 读取单元格的值
func (r *cellReader) readCell(axis string) string {
	if r.Err == nil {
		var value string
		value, r.Err = r.file.GetCellValue(r.sheetName, axis)
		return value
	}
	return ""
}

// 读取单元格的值，并格式化为 []string
func (r *cellReader) cellAsList(axis string) []string {
	var value = r.readCell(axis)
	// 读取成功, 字符串转数组
	if r.Err == nil {
		var list []string
		list, r.Err = utils.ReadFileToLines(strings.NewReader(value), true)
		return list
	}
	return nil
}
