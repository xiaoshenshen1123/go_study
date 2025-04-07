package main

import (
	"fmt"
	"sync"
	"time"
)

var idnames = make(map[int]string)
var lock sync.RWMutex

// map不允许同时读写，如果有不同go程同时操作map，需要对map上锁
// concurrent map read and map write
func main() {
	go func() {
		for {
			lock.Lock()
			idnames[0] = "duke"
			lock.Unlock()
		}
	}()
	go func() {
		for {
			lock.Lock()
			name := idnames[0]
			fmt.Println("name:", name)
			lock.Unlock()
		}
	}()
	for {
		fmt.Println("OVER!")
		time.Sleep(1 * time.Second)
	}
}
