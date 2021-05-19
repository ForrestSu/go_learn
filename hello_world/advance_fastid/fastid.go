package fastid

import (
	"fmt"
	"sync/atomic"
	"time"
)

// 日期格式
const dateLayout = "20060102"

const (
	seqBits     uint8 = 11 // 1毫秒内可生成的seq序号的二进制位数
	userBits    uint8 = 8  // user_id 二进制位数
	opBits      uint8 = 8  // operator_id 二进制位数
	machineBits uint8 = 8  // machine_id 二进制位数
	timeBits    uint8 = 28 // 时间毫秒
	// bit shift
	userShift    = seqBits
	opShift      = userShift + userBits
	machineShift = opShift + opBits
	timeShift    = machineShift + machineBits
	// mask
	seqMask  = ^(int64(-1) << seqBits)
	timeMask = ^(int64(-1) << timeBits)
)

/**
+-+------------------------+---------------+---------------+---------------+---+---------------+
|0|  timestamps(28 bits)   | machine_id (8)|  operator (8) | user_id (8)   |      seq (11)     |
+-+------------------------+---------------+---------------+---------------+---+---------------+
*/

// FastID 分布式唯一ID
type FastID struct {
	lastID int64
}

// NewFastID new
func NewFastID() *FastID {
	return &FastID{}
}

// StdFastID 默认的fastID
var StdFastID = NewFastID()

// LSB last byte
func LSB(v uint) byte {
	return byte(v)
}

// GetSeqFromID 从ID中获取seq
func GetSeqFromID(id int64) int64 {
	return id & seqMask
}

// GetTimeFromID 从ID中获取time
func GetTimeFromID(id int64) int64 {
	return (id >> timeShift) & timeMask
}

// GetMachineFromID 从ID中获取machineID
func GetMachineFromID(id int64) byte {
	if id > 0 {
		return byte(id >> machineShift)
	}
	return GetMachineByIP("127.0.0.1")
}

// GenID 生成唯一ID (格式：YYYYMMDD+ID，27位数字)
func (f *FastID) GenID(operatorID, userID uint) string {
	for {
		var preLastID = atomic.LoadInt64(&f.lastID)
		var seq = GetSeqFromID(preLastID)
		var preLastTime = GetTimeFromID(preLastID)
		var machineID = GetMachineFromID(preLastID)
		var tm = time.Now()
		var now = todayMills(&tm)
		if int64(now) != preLastTime {
			seq = 1
		} else if seq >= seqMask {
			// wait until the next millisecond
			var sleepNs = int(time.Millisecond) - tm.Nanosecond()%int(time.Millisecond)
			time.Sleep(time.Duration(sleepNs))
			continue
		} else {
			seq++
		}
		newID := generate(now, machineID, operatorID, userID, seq)
		if atomic.CompareAndSwapInt64(&f.lastID, preLastID, newID) {
			return datePrefix(&tm, newID)
		}
	}
}

// today millisecond
func todayMills(t *time.Time) int {
	return (t.Hour()*3600+t.Minute()*60+t.Second())*1000 + t.Nanosecond()/int(time.Millisecond)
}

// 使用位运算,生成int64 id
func generate(now int, machineID byte, operatorID, userID uint, seq int64) int64 {
	var p1 = seq
	var p2 = int64(LSB(userID)) << userShift
	var p3 = int64(LSB(operatorID)) << opShift
	var p4 = int64(machineID) << machineShift
	var p5 = int64(now) << timeShift
	var newID = p1 | p2 | p3 | p4 | p5
	return newID
}

// add date prefix
func datePrefix(t *time.Time, id int64) string {
	return t.Format(dateLayout) + fmt.Sprintf("%019d", id)
}
