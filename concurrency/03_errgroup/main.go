package main

import (
	"fmt"
	"io"
	"net/http"

	"golang.org/x/sync/errgroup"
)

// 任务信息
type taskInfo struct {
	URL string
	rsp []byte
	err error
}

func main() {
	tasks := []taskInfo{
		{URL: "http://www.golang.org/"},
		{URL: "http://www.google.com/"},
		{URL: "http://www.somestupidname.com/"},
	}
	g := errgroup.Group{}
	for i := range tasks {
		func(task *taskInfo) {
			g.Go(func() error {
				// Fetch the URL.
				rsp, err := http.Get(task.URL)
				task.err = err
				if err != nil {
					fmt.Printf("err: %v\n", err)
					return err
				}
				defer rsp.Body.Close()
				task.rsp, task.err = io.ReadAll(rsp.Body)
				fmt.Printf("ulr %v, OK!  body len= %v\n", task.URL, len(task.rsp))
				return err
			})
		}(&tasks[i])
	}
	// Wait for all HTTP fetches to complete.
	if err := g.Wait(); err != nil {
		fmt.Printf("err: %v\n", err)
	}

}
