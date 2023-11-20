package consumer_producer

import (
	"log"
	"testing"
	"time"

	gpool "github.com/ForrestSu/go_learn/concurrency/go-pool"
)

func init() {
	// flags are there: http://golang.org/pkg/log/#pkg-constants
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

// 生产者 消费者 模型
func TestConsumerProducer(t *testing.T) {
	var jobChan = make(chan gpool.Job, 2)
	wk := gpool.NewWorker(1)
	wk.Start(jobChan)
	for i := 0; i < 10; i++ {
		taskID := i
		jobChan <- func(id int) {
			log.Printf("执行任务=%d\n", taskID)
			time.Sleep(3 * time.Second)
		}
		log.Printf("add task %d\n", taskID)
	}
}
