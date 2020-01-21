package atomic

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func Test(t *testing.T) {
	runtime.GOMAXPROCS(4)
	wgs.Add(2)
	t1 := time.Now()

	//// 非原子操作
	//go incCounterBad()
	//go incCounterBad()

	// 原子函数
	go incCounterAtomic()
	go incCounterAtomic()


	fmt.Println("adding")
	wgs.Wait()
	fmt.Printf("now counter is %d\n", counter)


	t2 := time.Now()
	fmt.Println("费时：",t2.Sub(t1))
}

