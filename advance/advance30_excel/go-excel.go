package main

import (
	"fmt"
	"log"

	"github.com/kylelemons/godebug/pretty"
	excel "github.com/szyhf/go-excel"
)

func main() {
	var stdList []ImportCpUser
	err := excel.UnmarshalXLSX(fileName, &stdList)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println(pretty.Sprint(stdList))
}

func TestPlain() {
	conn := excel.NewConnecter()
	err := conn.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	var stdList []ImportCpUser
	// Generate an new reader of a sheet
	// sheetNamer: if sheetNamer is string, will use sheet as sheet name.
	//             if sheetNamer is int, will i'th sheet in the workbook, be careful the hidden sheet is counted. i ∈ [1,+inf]
	//             if sheetNamer is a object implements `GetXLSXSheetName()string`, the return value will be used.
	//             otherwise, will use sheetNamer as struct and reflect for it's name.
	// 			   if sheetNamer is a slice, the type of element will be used to infer like before.
	rd, err := conn.NewReader("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rd.Close()

	log.Println(conn.GetSheetNames())
	// readAll
	err = rd.ReadAll(&stdList)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println(pretty.Sprint(stdList))
}
