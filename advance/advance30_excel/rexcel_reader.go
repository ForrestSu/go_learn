package main

import (
	"io"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

// ReadExcel 将流数据读入为excel
func ReadExcel(r io.Reader) ([][]string, error) {
	fp, err := excelize.OpenReader(r)
	if err != nil {
		return nil, err
	}

	rows, err := fp.GetRows(fp.GetSheetName(0))
	if err != nil {
		return nil, err
	}
	return rows, nil
}
