package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func fn1(ch chan int) {
	defer wg.Done()
	for i := 0; i < 11; i++ {
		ch <- i
		fmt.Println("写入数据", i)

	}
	close(ch)

}
func fn2(ch chan int) {
	defer wg.Done()
	for v := range ch {
		fmt.Println("[读取数据]", v)
	}
}

func main() {
	var ch = make(chan int, 11)
	wg.Add(1)
	go fn1(ch)
	wg.Add(1)
	go fn2(ch)

	wg.Wait()

}
