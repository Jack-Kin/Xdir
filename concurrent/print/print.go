package print

import (
	"fmt"
	"sync"
)
// 这里全局
var wg sync.WaitGroup

func Print( s string, times int) {
	// 函数退出时，调用Done, 通知main函数工作已经完成
	// 如果没有这个会报错fatal error: all goroutines are asleep - deadlock!
	// 为什么会发生死锁？goroutine 在退出前调用了 wg.Done() ，程序应该正常退出的。

	defer wg.Done()

	for i:= 0; i < times; i++ {
		fmt.Println(s, i)
	}
}

func PrintPrime(prefix string, times int) {

	defer wg.Done()

	// 写法骚气
	nextNum:
		for i := 2; i < times; i++ {
			for j := 2; j < i; j++ {
				if i%j == 0 {
					continue nextNum
				}
			}
			fmt.Printf("%s:%d\n", prefix, i)
		}
	fmt.Printf("complete %s\n", prefix)
}