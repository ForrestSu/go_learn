package main

// fileName 文件名
const fileName = "白名单.xlsx"

type ImportCpUser struct {
	NickName string `csv:"昵称"`
	ID       string `csv:"id"`
	TagName  string `csv:"标签"`
	Price    int64  `csv:"报价"`
}

// GetXLSXSheetName sheet name
func (m ImportCpUser) GetXLSXSheetName() string {
	return "Sheet1"
}
