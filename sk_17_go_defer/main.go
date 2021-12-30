package main

import "fmt"

func main() {
	defer fmt.Println("last one from defer")
	defer fmt.Println("last one 2 from defer")
	fmt.Println("Hello")
	// Hello
	// last one 2 from defer
	// last one from defer

	dontForget(4)
	// 8
	// 4

	tricky()
	// 3
	// 3
	// 3

	tricky2()
	// 2
	// 1
	// 0
}

func dontForget(n int) {
	defer fmt.Println(n)
	defer func() {
		fmt.Println(n)
	}()
	n += n
}

func tricky() {
	for i := 0; i < 3; i++ {
		defer func() { fmt.Println(i) }()
	}
}

func tricky2() {
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}
}
