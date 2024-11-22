package main

import (
	"fmt"
	"time"
)

func main() {
	one := make(chan string)
	two := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		one <- "Hey"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		two <- "Heyoo"
	}()

	for x := 0; x < 2; x++ {
		select {
		case rec1 := <-one:
			fmt.Println("Received from one: ", rec1)
		case rec2 := <-two:
			fmt.Println("Received from two: ", rec2)
		}
	}
}
