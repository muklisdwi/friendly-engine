package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Channel Goroutine")
	fmt.Println("")
	// InfiniteLoopGoroutine()
	// ChannelChan()
	CekMutex()
}

func CekMutex() {
	var mutex = &sync.Mutex{}
	var x = 0
	fmt.Println(x)
	for {
		go func() {
			mutex.Lock()
			fmt.Println("Mutex A")
			x++
			fmt.Println(x)
			mutex.Unlock()
		}()

		go func() {
			mutex.Lock()
			fmt.Println("Mutex B")
			x++
			fmt.Println(x)
			mutex.Unlock()
		}()
	}
}

// channel melakukan blok saat mengirim dan menerima.
func ChannelChan() {
	ch := make(chan string, 2)
	ch <- "Ikan"
	ch <- "Bebek"

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func InfiniteLoopGoroutine() {
	go func() {
		for {
			fmt.Println("Ikan")
			time.Sleep(time.Second * 1)
		}
	}()

	rf := func() {
		for {
			fmt.Println("Bebek")
			time.Sleep(time.Second * 1)
		}
	}

	rf()
}
