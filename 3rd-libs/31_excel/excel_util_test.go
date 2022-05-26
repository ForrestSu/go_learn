package main

import (
	"log"
	"os"
	"testing"

	"github.com/gocarina/gocsv"

	"github.com/kylelemons/godebug/pretty"
)

// CpUser user
type CpUser struct {
	NickName string `csv:"昵称"`
	ID       string `csv:"id"`
	TagName  string `csv:"标签名"`
	Price    int64  `csv:"报价"`
}

func TestExcelToCSV(t *testing.T) {
	f, err := os.Open("/Users/sq/Desktop/CP白名单.xlsx")
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()
	var list []*CpUser
	err = gocsv.UnmarshalCSV(NewXlsxReader(f), &list)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(pretty.Sprint(list))
}

func BenchmarkUnmarshalExcel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f, err := os.Open("/Users/sq/Desktop/CP白名单.xlsx")
		if err != nil {
			log.Println(err)
			return
		}
		defer f.Close()
		var list []CpUser
		err = gocsv.UnmarshalCSV(NewXlsxReader(f), &list)
		if err != nil {
			log.Println(err)
			return
		}
	}
}
