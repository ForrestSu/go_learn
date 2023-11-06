package go1_21

import (
	"context"
	"go/types"
	"log/slog"
	"net"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// min、max
// slice、maps
// log/slog
// loopvar

func TestClear(t *testing.T) {
	m := map[int]string{1: "name", 2: "age"}
	assert.True(t, len(m) == 2)
	clear(m)
	assert.True(t, len(m) == 0)
}

func TestMaxMin(t *testing.T) {
	assert.Equal(t, 2, max(1, 2))
	assert.Equal(t, "a", min("a", "b"))
	p := types.NewPackage("github.com/stretchr/testify", "assert")
	t.Logf(p.GoVersion())
}

// GOEXPERIMENT=loopvar go test -v -run=TestForVar
func TestForVar(t *testing.T) {
	var all []*int
	for _, item := range []int{1, 10} {
		all = append(all, &item)
	}
	assert.Equal(t, 1, *all[0])
}

func TestSLog(t *testing.T) {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))
	slog.Info("hello", "name", "Al")
	slog.Error("ops", net.ErrClosed, "status", 500)
	slog.LogAttrs(context.Background(), slog.LevelError, "oops",
		slog.Int("status", 500), slog.Any("err", net.ErrClosed))
}

func TestAddressAble(t *testing.T) {
	var answerIDs []int64
	scan(&answerIDs)
}

func scan(dest interface{}) {
	reflectValue := reflect.ValueOf(dest)
	if reflectValue.Kind() == reflect.Slice {
		reflectValue.Set(reflect.MakeSlice(reflectValue.Type(), 0, 20))
	}
}
