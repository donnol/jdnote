package worker

import (
	"log"
	"os"
	"sync"
)

const (
	jobCount = 100
	errCount = 100
)

var logger = log.New(os.Stdout, "[Worker] ", log.LstdFlags|log.Lshortfile)

// Job 工作
type Job func() error

// Worker 工人
type Worker struct {
	// 所有管道都要有make, read, write, close操作
	limitChan chan struct{} // 并发控制管道
	stopChan  chan struct{} // 停止管道
	jobChan   chan Job      // 工作管道
	errChan   chan error    // 错误管道

	jobs []Job // 工作列表

	wg *sync.WaitGroup
}

// New 新建
func New(n int) *Worker {
	if n <= 0 {
		n = 10
	}
	return &Worker{
		limitChan: make(chan struct{}, n),
		stopChan:  make(chan struct{}),
		jobChan:   make(chan Job),
		errChan:   make(chan error, errCount),
		jobs:      make([]Job, 0),
		wg:        new(sync.WaitGroup),
	}
}

// Start 开始
func (w *Worker) Start() {
	go w.handleError()
	go w.start()
}

func (w *Worker) start() {
	for {
		select {
		case job, ok := <-w.jobChan: // 有工作
			if !ok {
				continue
			}

			w.do(job)

		case <-w.stopChan:
			w.close()
			return

		default:
			if len(w.jobs) != 0 {
				for _, job := range w.jobs {
					w.do(job)
				}
				// 清空
				w.jobs = make([]Job, 0)
			}
		}
	}
}

func (w *Worker) do(job Job) {
	// 占据一个坑
	w.limitChan <- struct{}{}
	w.wg.Add(1)

	// 开始工作
	go func() {
		defer func() {
			if r := recover(); r != nil {
				logger.Printf("job: %+v\n", r)
			}

			// 释放一个坑
			<-w.limitChan
			w.wg.Done()
		}()

		// 执行
		if err := job(); err != nil {
			w.errChan <- err
		}
	}()
}

func (w *Worker) handleError() {
	for {
		select {
		case err, ok := <-w.errChan:
			if !ok {
				return
			}
			logger.Printf("err is %v\n", err)
		}
	}
}

// Stop 停止
func (w *Worker) Stop() {
	w.wait()

	w.stopChan <- struct{}{}
}

func (w *Worker) close() {
	// close管道时，有可能panic
	defer func() {
		if r := recover(); r != nil {
			logger.Printf("close: %+v\n", r)
		}
	}()

	close(w.stopChan)
	close(w.errChan)
	close(w.jobChan)
	close(w.limitChan)
}

func (w *Worker) wait() {
	// 等待所有工作完成
	w.wg.Wait()
}

// Push 添加
func (w *Worker) Push(job Job) {
	// 如果添加的速度超过处理的速度，就会阻塞。这时，先把job存到jobs里
	select {
	case w.jobChan <- job:
	default:
		w.jobs = append(w.jobs, job)
		logger.Println("Full load. Append to jobs")
	}
}
