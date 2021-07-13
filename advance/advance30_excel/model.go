package main

// fileName 文件名
const fileName = "白名单.xlsx"

type ImportCpUser struct {
	NickName string `xlsx:"column(昵称)"`
	ID       string `xlsx:"column(id)"`
	TagName  string `xlsx:"column(属性标签)"`
	Price    int64  `xlsx:"column(报价)"`
}

// GetXLSXSheetName sheet name
func (m ImportCpUser) GetXLSXSheetName() string {
	return "Sheet1"
}
