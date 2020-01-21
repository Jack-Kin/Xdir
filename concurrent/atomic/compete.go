package atomic

import (
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	// 所有 goroutine需要增加counter的值
	counter int32

	// 用来等待程序结束
	wgs sync.WaitGroup
)

// 增加包里counter的值
// 非原子操作加法
func incCounterBad() {

	defer wgs.Done()

	for i:= 0; i < 2000; i++ {
		// 捕获counter的值
		val := counter
		// 当前goroutine从线程退出，放回到队列
		// 强制切换goroutine，节目效果， 目的是让竞争状态更加明显
		runtime.Gosched()
		// 增加本地value
		val ++
		// 保存值，返回给counter
		counter = val
	}
}

// 原子函数以操作系统底层的枷锁机制同步访问变量
func incCounterAtomic(){

	defer wgs.Done()

	for i:=0; i < 2000; i++ {
		// 安全的加1; 读和写是LoadInt StoreInt
		atomic.AddInt32(&counter, 1)
		runtime.Gosched()
	}
}