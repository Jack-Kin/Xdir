package mutex

import (
	"runtime"
	"sync"
)

var (
	counter int
	wgs sync.WaitGroup
	mutex sync.Mutex
)


// 互斥锁用于在代码当中建立一个临界区，保证同一时间只有一个G执行临界区的代码。
func incCounterMutex() {
	defer wgs.Done()

	for i:=0; i < 2000; i++{
		// 建立临界区， 同一时刻只允许一个goroutine进入这个临界区
		mutex.Lock()
		//{
		//	counter++
		//}
		{
			// 节目效果， 就算强制将当前goroutine退出当前线程，
			// 调度器还是会再次分配给这个goroutine继续运行
			val := counter
			runtime.Gosched()
			val++
			counter = val
		}
		mutex.Unlock()
	}
}