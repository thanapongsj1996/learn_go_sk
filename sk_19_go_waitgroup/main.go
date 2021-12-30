package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	บางทีไม่รู้ว่าใน go routine อื่นทำเสร็จตอนไหน ใช้ sync.WaitGroup มาช่วยได้
*/

func main() {
	start := time.Now()

	wg := &sync.WaitGroup{}
	wg.Add(3)

	go doSomeThing(wg)
	go doSomeThing(wg)
	go doSomeThing(wg)

	wg.Wait()
	fmt.Println(time.Since(start))
	// 101.272016ms
}

func doSomeThing(wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("something...")
	wg.Done()
}
