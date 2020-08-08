package main

import (
	"fmt"
	"sync"
	"time"
)

var wg = new(sync.WaitGroup)

func main3() {
	ch := make(chan struct{})
	ch1 := make(chan struct{})

	wg.Add(2)
	go printNum(ch)
	go printAlpha(ch1)

	// 不加睡眠的时候，会打印：1AB23CD45EF67GH89IJ10
	// 怎么样可以把这个睡眠去掉呢？
	time.Sleep(10 * time.Millisecond)

	for i := 0; i < 5; i++ {
		for _, s := range []int{
			0, 0, 1, 1,
		} {
			if s == 0 {
				ch <- struct{}{}
			} else {
				ch1 <- struct{}{}
			}
		}
	}

	wg.Wait()
}

func printNum(ch chan struct{}) {
	defer wg.Done()

	for i := 1; i <= 10; i++ {
		<-ch
		fmt.Print(i)
	}
}

func printAlpha(ch chan struct{}) {
	defer wg.Done()

	for i := 65; i <= 74; i++ {
		<-ch
		fmt.Printf("%c", i)
	}
}
