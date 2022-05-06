package main

import (
	"io"

	"github.com/xuri/excelize/v2"
)

// ReadExcel 解析excel数据为二维数组
func ReadExcel(r io.Reader) ([][]string, error) {
	f, err := excelize.OpenReader(r)
	if err != nil {
		return nil, err
	}
	return f.GetRows(f.GetSheetName(0))
}

// XlsxReader 实现 CSVReader 相关接口
type XlsxReader struct {
	r io.Reader
}

// NewXlsxReader new
func NewXlsxReader(r io.Reader) *XlsxReader {
	return &XlsxReader{r}
}

// Read never used
func (x *XlsxReader) Read() ([]string, error) {
	panic("implements me!") // not used
}

// ReadAll read all
func (x *XlsxReader) ReadAll() ([][]string, error) {
	return ReadExcel(x.r)
}
