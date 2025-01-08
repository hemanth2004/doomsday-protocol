package dday

import (
	"fmt"
	"time"
)

var numberChannel = make(chan int)

func PrintNumberGoroutine() {
	number := 0
	for {
		select {
		case n := <-numberChannel:
			number += n
		default:
			fmt.Println(number)
			time.Sleep(time.Second)
		}
	}
}

func AddToNumber(n int) {
	numberChannel <- n
}
