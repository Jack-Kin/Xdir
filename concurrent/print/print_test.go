package print

import (
	"fmt"
	"runtime"
	"testing"
)


func Test(t *testing.T) {
	//这里设P的数量， 改成1 2 3 4 5试试
	runtime.GOMAXPROCS(4)

	// 这里居然看的到print.go里定义的变量
	// 几个goroutine， add几个
	wg.Add(2)

	fmt.Println("Start goroutines")
	//go Print("goroutine A:", 10)
	//go Print("goroutine B:", 10)
	//go Print("goroutine C:", 10)

	// 光打印看不出来concurrency，需要一些计算
	go PrintPrime("goroutine A", 5000)
	go PrintPrime("goroutine B", 5000)
	//go PrintPrime("goroutine C", 5000)


	//等待goroutine结束
	wg.Wait()

	fmt.Println("Terminating program")
}
