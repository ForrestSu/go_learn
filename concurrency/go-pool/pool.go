package gpool

import "sync"

type Pool struct {
	workers    []*Worker
	channelJob chan Job
	// 重用 waitGroup：再次调用 Add() 之前，需要在途的 Wait() 调用都返回之后；
	wg sync.WaitGroup
}

// NewPoolWithSize 构造线程池
func NewPoolWithSize(size int) *Pool {
	p := &Pool{
		workers:    make([]*Worker, 0, size),
		channelJob: make(chan Job),
	}
	for i := 0; i < size; i++ {
		p.workers = append(p.workers, NewWorker(i))
	}
	return p
}

func (p *Pool) Start() {
	for _, w := range p.workers {
		w.Start(p.channelJob)
	}
}

// AddJob 追加任务
func (p *Pool) AddJob(job Job) {
	p.wg.Add(1)
	p.channelJob <- func(id int) {
		defer p.wg.Done()
		job(id)
	}
}

// ExitWait 等待所有任务结束,退出前执行一次
func (p *Pool) ExitWait() {
	p.wg.Wait()
}

func (p *Pool) Stop() {
	for _, w := range p.workers {
		w.Stop()
	}
}
