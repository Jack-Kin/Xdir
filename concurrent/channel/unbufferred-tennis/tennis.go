package main

/*
	这代码太骚了
 */

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// 创建一个无缓冲的通道
	court := make(chan int)

	wg.Add(2)

	go player("Nadal", court)
	go player("Djokovic", court)

	// 发球
	court <- 1

	// 等待游戏结束
	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	// go没有while
	for {
		// 接球
		ball, ok := <-court

		// 判断通道是否关闭
		if !ok {
			fmt.Printf("Player %s win\n", name)
			return
		}
		// 判断击球是否miss
		n := rand.Intn(100)
		if n % 13 == 0 {
			fmt.Printf("Player %s miss\n", name)
			close(court)
			return
		}

		fmt.Printf("Player %s hit %d\n", name, ball)
		ball++

		// 把球打回去
		court <- ball
	}

}


/*
Output例子:
		Player Djokovic hit 1
		Player Nadal hit 2
		Player Djokovic hit 3
		Player Nadal hit 4
		Player Djokovic hit 5
		Player Nadal hit 6
		Player Djokovic hit 7
		Player Nadal hit 8
		Player Djokovic hit 9
		Player Nadal hit 10
		Player Djokovic hit 11
		Player Nadal hit 12
		Player Djokovic hit 13
		Player Nadal hit 14
		Player Djokovic hit 15
		Player Nadal hit 16
		Player Djokovic hit 17
		Player Nadal miss
		Player Djokovic win
 */