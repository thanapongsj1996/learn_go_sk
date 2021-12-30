package main

import (
	"fmt"
	"time"
)

// goroutine is called lightweight thread

func main() {
	start := time.Now()

	//doSomeThing()
	//doSomeThing()
	//doSomeThing()
	// 423.373358ms

	go doSomeThing()
	go doSomeThing()
	go doSomeThing()
	// 121.037297ms
	time.Sleep(120 * time.Millisecond)
	fmt.Println(time.Since(start))

	example()
	// --+--+-+--+--
	// มันอาจจะแย่งกันออกมาในบางจังหวะ

}

func doSomeThing() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("something...")
}

func example() {
	go func() {
		for {
			time.Sleep(100 * time.Millisecond)
			fmt.Print("-")
		}
	}()
	go func() {
		for {
			time.Sleep(200 * time.Millisecond)
			fmt.Print("+")
		}
	}()
	time.Sleep(time.Second)
}
