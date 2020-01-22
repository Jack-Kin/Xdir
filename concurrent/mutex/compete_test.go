package mutex

import (
	"fmt"
	"testing"
)

func TestMutex(t *testing.T) {
	wgs.Add(2)

	go incCounterMutex()
	go incCounterMutex()

	fmt.Println("adding...")
	wgs.Wait()
	fmt.Printf("now counter is: %d\n", counter)
}