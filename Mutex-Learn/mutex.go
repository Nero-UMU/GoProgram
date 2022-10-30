package main

import (
	"fmt"
	"sync"
	"time"
)

var value = 0

func main() {
	lock := sync.Mutex{}
	for i := 0; i < 1000; i++ {
		go func() {
			lock.Lock()
			value++
			lock.Unlock()
		}()
	}
	//为了等待所有的子协程执行完毕
	time.Sleep(time.Second * 10)

	fmt.Println(value)
}
