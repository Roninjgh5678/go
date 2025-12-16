package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now().UnixMilli()

	for i := 2; i < 1200000; i++ {
		var flag = true
		for num := 2; num < i; num++ {
			if i%num == 0 {
				flag = false
				break
			}
		}
		if flag {
			fmt.Println(i, "是素数")
		}
	}
	end := time.Now().UnixMilli()
	contuine := (end - start)
	fmt.Println("1运行的时间为", start)
	fmt.Println("2运行的时间为", end)
	fmt.Println("3运行的时间为", contuine)

}
