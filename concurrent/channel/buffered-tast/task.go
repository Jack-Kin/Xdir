package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const(
	numberGoroutines = 4
	taskLoad = 10
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	tasks := make(chan string, taskLoad)

	wg.Add(numberGoroutines)

	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task : %d", post)
	}

	// 通道关闭之后，goroutine可以从通道接收数据，但是不能发送数据
	// 通道关闭理解为数据只进不出
	close(tasks)

	wg.Wait()
}

func worker(tasks chan string, worker int) {
	defer wg.Done()

	for {
		task, ok := <- tasks
		if !ok {
			fmt.Printf("Worker %d : Shutting Down\n", worker)
			return
		}

		fmt.Printf("Worker %d : Starting %s\n", worker, task)

		// 这里做任务：睡觉！
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep) * time.Millisecond)

		// 等一段时间表示任务做完
		fmt.Printf("Worker %d : Completed %s\n", worker, task)
	}
}
