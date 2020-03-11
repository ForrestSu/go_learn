package ch17_obj_pool

import (
	"errors"
	"strconv"
	"time"
)

// reusable object
type ConnObj struct {
	id   int
	ip   string
	name string
}
type ObjPool struct {
	bufChan chan *ConnObj
}
func NewObjPool(size int) *ObjPool {
	pool := &ObjPool{}
	pool.bufChan = make(chan *ConnObj, size)
	for i := 0; i < size; i++ {
		pool.bufChan <- &ConnObj{
			id:   i,
			name: "conn-" + strconv.Itoa(i),
		}
	}
	return pool
}
func (pool *ObjPool) Take(timeout time.Duration) (*ConnObj, error) {
	select {
	case conn := <-pool.bufChan:
		return conn, nil
	case <-time.After(timeout): //超时控制
		return nil, errors.New("time out")
	}
}
func (pool *ObjPool) Release(conn *ConnObj) error {
	select {
	case pool.bufChan <- conn:
		return nil
	default:
		return errors.New("overflow")
	}
}