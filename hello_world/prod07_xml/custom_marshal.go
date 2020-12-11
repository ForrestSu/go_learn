package main

import (
	"encoding/xml"
	"fmt"
	"time"
)

// 将一个字段 Marshal 为 2 个xml元素

type Appointment struct {
	DateTime time.Time
}

// inner
type appointmentExport struct {
	XMLName struct{} `xml:"appointment"`
	Date    string   `xml:"date"`
	Time    string   `xml:"time"`
}

// 实现 MarshalXML 接口
func (a *Appointment) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	n := &appointmentExport{
		Date: a.DateTime.Format("2006-01-02"),
		Time: a.DateTime.Format("15:04"),
	}
	return e.Encode(n)
}

func main() {
	a := &Appointment{time.Now()}
	output, _ := xml.MarshalIndent(a, "", "    ")
	fmt.Println(string(output))
}

// prints:
// <appointment>
//     <date>2016-04-15</date>
//     <time>17:43</time>
// </appointment>

/**
 * 将一个字段 Marshal 为 2 个xml元素
 * https://stackoverflow.com/questions/36599180/marshal-single-field-into-two-tags-in-golang
 */
