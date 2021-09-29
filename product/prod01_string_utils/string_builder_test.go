package prod01_string_utils

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func Test_StringBuilder(t *testing.T) {

	var b strings.Builder
	t.Log(b.Len())
	t.Log(b.Cap())
	b.Grow(255)
	t.Log(b.Len())
	t.Log(b.Cap())
}

func Benchmark_Sprintf(b *testing.B) {
	uin := "3192695574"
	qqMusicKey := "@DNNPx7IFA"
	for i := 0; i < b.N; i++ {
		_ = fmt.Sprintf("qqmusic_uin=%s; qqmusic_key=%s; qq=%s; authst=%s; p_uin=%s; p_skey=%s; uin=%s; skey=%s;"+
			" qm_keyst=%s; lskey=%s; p_luin=%s; p_lskey=%s;", uin, qqMusicKey, uin, qqMusicKey, uin, qqMusicKey, uin, qqMusicKey,
			qqMusicKey, qqMusicKey, uin, qqMusicKey)
	}
}

func Benchmark_BytesBuffer(b *testing.B) {
	uin := "3192695574"
	qqMusicKey := "@DNNPx7IFA"
	uinNames := [...]string{"qqmusic_uin", "qq", "p_uin", "uin", "qm_keyst", "p_luin"}
	keyNames := [...]string{"qqmusic_key", "authst", "p_skey", "skey", "lskey", "p_lskey"}
	// b.Grow(255)
	for i := 0; i < b.N; i++ {
		var b bytes.Buffer
		for i := range uinNames {
			b.WriteString(uinNames[i])
			b.WriteString("=")
			b.WriteString(uin)
			b.WriteString("; ")

			b.WriteString(keyNames[i])
			b.WriteString("=")
			b.WriteString(qqMusicKey)
			b.WriteString("; ")
		}
		_ = b.String()
	}
}

// Benchmark_BytesBuffer_Grows
// Benchmark_BytesBuffer_Grows-4   	 2072634	       586.9 ns/op
func Benchmark_BytesBuffer_Grows(b *testing.B) {
	uin := "3192695574"
	qqMusicKey := "@DNNPx7IFA"
	uinNames := [...]string{"qqmusic_uin", "qq", "p_uin", "uin", "qm_keyst", "p_luin"}
	keyNames := [...]string{"qqmusic_key", "authst", "p_skey", "skey", "lskey", "p_lskey"}
	// b.Grow(255)
	for i := 0; i < b.N; i++ {
		var b bytes.Buffer
		b.Grow(255)
		for i := range uinNames {
			b.WriteString(uinNames[i])
			b.WriteString("=")
			b.WriteString(uin)
			b.WriteString("; ")

			b.WriteString(keyNames[i])
			b.WriteString("=")
			b.WriteString(qqMusicKey)
			b.WriteString("; ")
		}
		_ = b.String()
	}
}

// Benchmark_StringBuilder
// Benchmark_StringBuilder-4   	 2727388	       436.7 ns/op
func Benchmark_StringBuilder(b *testing.B) {
	uin := "3192695574"
	qqMusicKey := "@DNNPx7IFA"
	uinNames := [...]string{"qqmusic_uin", "qq", "p_uin", "uin", "qm_keyst", "p_luin"}
	keyNames := [...]string{"qqmusic_key", "authst", "p_skey", "skey", "lskey", "p_lskey"}
	for i := 0; i < b.N; i++ {
		var b strings.Builder
		for i := range uinNames {
			b.WriteString(uinNames[i])
			b.WriteString("=")
			b.WriteString(uin)
			b.WriteString("; ")

			b.WriteString(keyNames[i])
			b.WriteString("=")
			b.WriteString(qqMusicKey)
			b.WriteString("; ")
		}
		_ = b.String()
	}
}

const (
	// Q音-新的票据格式
	TicketKeyPrefixQQ = "Q_H_L_"
	TicketKeyPrefixWX = "W_X_"
	// 旧的票据格式 eg: @DNNPx7IFA
	TicketKeyPrefixOld = "@"
)

var (
	uinTags = [...]string{"qqmusic_uin", "qq", "p_uin", "uin", "qm_keyst", "p_luin"}
	keyTags = [...]string{"qqmusic_key", "authst", "p_skey", "skey", "lskey", "p_lskey"}
)

func handleCookies(uin string, musicKey string) string {
	// 新的票据格式
	if strings.HasPrefix(musicKey, TicketKeyPrefixQQ) || strings.HasPrefix(musicKey, TicketKeyPrefixWX) {
		return fmt.Sprintf("qqmusic_uin=%s; qqmusic_key=%s;", uin, musicKey)
	}

	// 否则, 统一按旧的票据处理
	var b strings.Builder
	b.Grow(255)
	for i := range uinTags {
		// uin
		b.WriteString(uinTags[i])
		b.WriteString("=")
		b.WriteString(musicKey)
		b.WriteString("; ")

		// musicKey
		b.WriteString(keyTags[i])
		b.WriteString("=")
		b.WriteString(musicKey)
		b.WriteString("; ")
	}
	result := b.String()
	// trim last space
	result = result[:len(result)-1]
	return result
}

func Benchmark_StringBuilder_Grows(b *testing.B) {
	// expects := "qqmusic_uin=3192695574; qqmusic_key=@DNNPx7IFA; qq=3192695574; authst=@DNNPx7IFA; p_uin=3192695574; p_skey=@DNNPx7IFA; uin=3192695574; skey=@DNNPx7IFA; qm_keyst=@DNNPx7IFA; lskey=@DNNPx7IFA; p_luin=3192695574; p_lskey=@DNNPx7IFA;"

	uin := "3192695574"
	qqMusicKey := "@DNNPx7IFA"

	for i := 0; i < b.N; i++ {
		_ = handleCookies(uin, qqMusicKey)
	}
	// assert.Equal(t, expects, actual)
}
