package runner

/*
	演示使用通道通信、同步等待，监控程序
	拓展：
		1. 交给定时任务去执行，比如cron
			E.g: https://www.jianshu.com/p/e629d637bf4c
		2. 更高效率的并发，更多灵活的控制程序的生命周期，更高效的监控等
*/

import (
	"errors"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	interrupt chan os.Signal
	// 如果执行任务发生错误，发回一个error接口类型的值； 不错误，发回一个nil
	complete chan error
	// 如果这个通道接收到time.Time的值，这个程序会试图清理状态并停止工作
	// 表示为timeout是个通道，it can only be used to receive time.Time
	timeout <-chan time.Time
	tasks []func(int)
}

var ErrTimeout = errors.New("received timeout")

var ErrInterrupt = errors.New("received interrupt")

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt:	make(chan os.Signal, 1),
		complete: 	make(chan error),
		timeout: 	time.After(d),
		//tasks字段的零值是nil, 没有必要明确初始化
	}
}

func (r *Runner) Add(tasks...func(int)) {
	r.tasks = append(r.tasks, tasks...)
}

// Start执行所有任务，并监视通道时间
func (r *Runner) Start() error {
	// 接收所有中断信号
	// 把os.Interrupt转发到r,interrupt
	signal.Notify(r.interrupt, os.Interrupt)

	go func() {
		r.complete <- r.run()
	}()

	select {
	// 当任务完成时发出的信号
	case sth := <- r.complete:
		return sth
	// 当处理任务运行超时时发出的信号
	case <- r.timeout:
		return ErrTimeout
	}
}


func (r *Runner) run() error {
	for id, task := range r.tasks {
		if r.gotInterrupt() {
			return ErrInterrupt
		}
		task(id)
	}
	return nil
}

func (r *Runner) gotInterrupt() bool {
	select {
	case <-r.interrupt:
		signal.Stop(r.interrupt)
		return true
	// 继续正常运行
	default:
		return false
	}
}