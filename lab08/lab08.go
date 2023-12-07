package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	doorStatus string
	handStatus string
	mtx        sync.Mutex // Mutex for synchronization
)

func hand(wg *sync.WaitGroup) {
	mtx.Lock() // Lock the critical section
	handStatus = "in"
	time.Sleep(time.Millisecond * 200)
	handStatus = "out"
	mtx.Unlock() // Unlock after modifying the shared resource
	wg.Done()
}

func door(wg *sync.WaitGroup) {
	mtx.Lock() // Lock the critical section
	doorStatus = "close"
	time.Sleep(time.Millisecond * 200)
	if handStatus == "in" {
		fmt.Println("夾到手了啦！")
	} else {
		fmt.Println("沒夾到喔！")
	}
	doorStatus = "open"
	mtx.Unlock() // Unlock after modifying the shared resource
	wg.Done()
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(2)
		go door(&wg)
		go hand(&wg)
		wg.Wait()
		time.Sleep(time.Millisecond * 200)
	}
}
