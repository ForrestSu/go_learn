package main

import (
	"bytes"
)

// DailyNote daily
type DailyNote struct {
	Title string
	// 上周工作总结 (多行)
	PrevWork []string
	// 工作感悟 (多行)
	Conclusion []string
	// 本周工作计划 (多行)
	CurrentWork []string
}

func NewDailyNote(title string) *DailyNote {
	return &DailyNote{
		Title: title,
	}
}
func (d *DailyNote) Bytes() []byte {
	var sb = bytes.Buffer{}
	sb.Grow(512)
	sb.WriteString("## " + d.Title + "\n")
	sb.WriteString("- 上周工作总结\n")
	for _, s := range d.PrevWork {
		sb.WriteString("  - " + s)
		sb.WriteString("\n")
	}
	sb.WriteString("- 工作感悟\n")
	for _, s := range d.Conclusion {
		sb.WriteString("  - " + s)
		sb.WriteString("\n")
	}
	sb.WriteString("- 本周工作计划\n")
	for _, s := range d.CurrentWork {
		sb.WriteString("  - " + s)
		sb.WriteString("\n")
	}
	sb.WriteString("\n")
	return sb.Bytes()
}
