package main

import (
	"fmt"
)

func merge(channels []chan int) chan int {
	merged := make(chan int)
	//wg := sync.WaitGroup{}

	go func() {

		merge := func(ch <-chan int) {
			//defer wg.Done()
			for i := range ch {
				merged <- i
			}
		}
		cnt := 0

		for _, i := range channels {
			if cnt == len(channels) {
				close(merged)
				//return merged
			} else {
				go merge(i)
			}

			cnt++
		}

		//wg.Add(len(channels))
		//for _, i := range channels {
		//	go merge(i)
		//}
		//
		//wg.Wait()
		//defer close(merged)
	}()
	return merged
}

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)

	go func() {
		defer close(chan1)
		for i := 0; i < 10; i++ {
			chan1 <- i
		}
	}()

	go func() {
		defer close(chan2)
		for i := 10; i < 15; i++ {
			chan2 <- i
		}
	}()

	merged := merge([]chan int{chan1, chan2})

	for i := range merged {
		//fmt.Println(1)
		fmt.Println(i)
	}
}
