package main

/*
	goroutine里套goroutine
	一个时间只有一个runner
 */

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	baton := make(chan int)

	wg.Add(1)

	go Runner(baton)

	baton <- 1

	wg.Wait()
}

func Runner(baton chan int) {

	runner := <- baton

	fmt.Printf("Runner %d Running with Baton\n", runner)

	if runner != 4 {
		fmt.Printf("Runner %d to the Line\n", runner)
		runner = runner +1
		go Runner(baton)
	}

	time.Sleep(1000 * time.Millisecond)

	if runner == 4 {
		fmt.Printf("Runner %d finished, Race Over\n", runner)
		wg.Done()
		return
	}

	// 交接接力棒
	fmt.Printf("Runner %d exchange with runner %d\n", runner, runner)

	baton <- runner
}

/*
Output:
		Runner 1 Running with Baton
		Runner 1 to the Line
		Runner 2 exchange with runner 2
		Runner 2 Running with Baton
		Runner 2 to the Line
		Runner 3 exchange with runner 3
		Runner 3 Running with Baton
		Runner 3 to the Line
		Runner 4 finished, Race Over
 */