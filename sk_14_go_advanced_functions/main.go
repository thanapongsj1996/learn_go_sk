package main

import "fmt"

func main() {
	// First class function
	var add = func(a, b int) int {
		return a + b
	}
	fmt.Println(add(1, 2)) // 3

	var newFuncName = double
	fmt.Println(newFuncName(4)) // 8

	// Higher order function (function ที่รับ params เป็น function ได้, หรือ return ออกมาเป็น function)
	s := HOFGreeting(func() string {
		return "Boy"
	})
	fmt.Println(s) // Hello, Boy

	fmt.Println(HOFGreeting(newNameFunc("Thanapong"))) // Hello, Thanapong

	counter01 := newCounterFunc()
	fmt.Println("counter01: ", counter01()) // 1
	fmt.Println("counter01: ", counter01()) // 2
	fmt.Println("counter01: ", counter01()) // 3

	counter02 := newCounterFunc()
	fmt.Println("counter02: ", counter02()) // 1
	fmt.Println("counter01: ", counter01()) // 4
	fmt.Println("counter02: ", counter02()) // 2
}

// First class function
func double(n int) int {
	return n * 2
}

// Higher order function
type nameFunc func() string

func HOFGreeting(nameFn nameFunc) string {
	return fmt.Sprintf("Hello, %s", nameFn())
}
func newNameFunc(name string) nameFunc {
	return func() string { return name }
}

// Closure function เก็บ state ในตัวเองได้
func newCounterFunc() func() int {
	var i int
	return func() int {
		i++
		return i
	}
}
