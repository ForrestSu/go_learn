package fastid

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLSB(t *testing.T) {
	assert.Equal(t, byte(0xfe), LSB(0x123456fe))
}

func TestFastID_GenID(t *testing.T) {
	var operatorID uint = 10
	var userID uint = 123
	sn := StdFastID.GenID(operatorID, userID)
	// 202105171838043898712414209
	t.Log(sn)
}

func TestFastID(t *testing.T) {
	var lastID int64 = 1415918331035375617
	var expectSeq int64 = 1
	var expectTime int64 = 41208647 // 11:26:48.647
	var expectMachineID byte = 12
	assert.Equal(t, expectSeq, GetSeqFromID(lastID))
	assert.Equal(t, expectTime, GetTimeFromID(lastID))
	assert.Equal(t, expectMachineID, GetMachineFromID(lastID))
}

func BenchmarkFastID_GenID(b *testing.B) {
	var operatorID uint = 10
	var userID uint = 123
	for i := 0; i < b.N; i++ {
		_ = StdFastID.GenID(operatorID, userID)
	}
}
