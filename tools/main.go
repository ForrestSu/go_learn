package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"github.com/ForrestSu/go_learn/tools/utils"

	"github.com/xuri/excelize/v2"
)

func validateFormat(f *excelize.File, sheetName string, axis string, expect string) error {
	cell, err := f.GetCellValue(sheetName, axis)
	if err != nil {
		return fmt.Errorf("sheet %v 格式不正确， err=%v", sheetName, err)
	}
	if cell != expect {
		return fmt.Errorf("sheet %v 期望值为=%v, 实际值=[%v]", sheetName, expect, cell)
	}
	return nil
}

// Validate 1 检查格式
func Validate(f *excelize.File) {
	sheets := f.GetSheetList()
	for _, sheetName := range sheets {
		if _, err := utils.FormatTitle(sheetName); err != nil {
			log.Printf("err = %v\n", err)
			return
		}
		if err := validateFormat(f, sheetName, "B6", "上周工作总结"); err != nil {
			log.Printf("err = %v\n", err)
		}
		if err := validateFormat(f, sheetName, "B13", "工作感悟"); err != nil {
			log.Printf("err = %v\n", err)
		}
	}
}

func openFile(fileName string) {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// log.Printf("%v\n", f.GetSheetList())
	// 1 检查格式
	Validate(f)
	// 2 解析为对象
	var records []*DailyNote
	var sheets = f.GetSheetList()
	for i := len(sheets) - 1; i >= 0; i-- {
		note, err := ParseSheet(f, sheets[i])
		if err != nil {
			panic(err)
		}
		records = append(records, note)
	}
	// 3 格式化为 markdown
	var sb = &bytes.Buffer{}
	for i := len(records) - 1; i >= 0; i-- {
		sb.Write(records[i].Bytes())
	}
	if err := os.WriteFile(curDir+"daily.md", sb.Bytes(), 0644); err != nil {
		panic(err)
	}
}

func ParseSheet(f *excelize.File, sheetName string) (*DailyNote, error) {
	title, err := utils.FormatTitle(sheetName)
	if err != nil {
		return nil, err
	}
	r := newCellReader(f, sheetName)
	note := NewDailyNote(title)
	note.PrevWork = r.cellAsList("B9")
	note.Conclusion = r.cellAsList("C13")
	note.CurrentWork = r.cellAsList("B17")
	if r.Err != nil {
		return nil, r.Err
	}
	return note, nil
}

const curDir = "/Users/sq/go_workspace/go_learn/tools/"

func main() {
	fileName := curDir + "工作周报201507-201611.xlsx"
	openFile(fileName)
}
