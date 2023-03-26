package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	dataChannel := make(chan int)
	go func() {
		for i := 1; i < 9000000; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				dataChannel <- doWork()
			}()
		}
		wg.Wait()
		close(dataChannel)
	}()

	for i := range dataChannel {
		fmt.Println(i)
	}
}

func doWork() int {
	time.Sleep(time.Second)
	return rand.Intn(100)
}
