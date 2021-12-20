package main

import (
	"fmt"
	"strconv"
)

func main() {
	// --------------- Variables ---------------
	var i int
	fmt.Println(i) // 0

	var j int = 14
	fmt.Println(j) // 14

	var n0 string
	fmt.Printf("%T %s \n", n0, n0) // string

	var n1 string = "Boy"
	fmt.Printf("%T %s \n", n1, n1) // string Boy

	var ok bool
	fmt.Printf("%T %v \n", ok, ok) // bool false

	var ok2 bool
	fmt.Printf("%T %v \n", ok2, ok2) // bool true

	// --------------- Type interface ---------------
	var k = 15
	var s = "hello"
	var ok3 = true
	fmt.Printf("%T %v \n", k, k)     // int 15
	fmt.Printf("%T %v \n", s, s)     // string hello
	fmt.Printf("%T %v \n", ok3, ok3) // bool true

	a := 20
	b := "Bangkok"
	c := true
	fmt.Printf("%T %v \n", a, a) // int 20
	fmt.Printf("%T %v \n", b, b) // string Bangkok
	fmt.Printf("%T %v \n", c, c) // bool true

	// --------------- Constants ---------------
	const defaultValue int = 1
	const defaultTitle = "Go"
	const defaultStatus = true

	const (
		sunday    = iota // 0
		monday           // 1
		tuesday          // 2
		wednesday        // 3
		thursday         // 4
		friday           // 5
		saturday         // 6
	)
	fmt.Printf("%T %v \n", monday, monday) // int 1

	// --------------- Functions ---------------
	greeting("Thanapong", "somjai") // Hello Thanapong somjai
	fmt.Println(add(1, 3))          // 4
	fmt.Println(swap(2, 5))         // 5 2

	// --------------- Conditions ---------------
	var name string = "Boy"
	if name != "" {
		fmt.Printf("Hello, %s\n", name)
	} else {
		fmt.Println("Hello, what is your name ?")
	}

	// --- If with a short statement
	if n, err := strconv.Atoi("5s"); err != nil {
		fmt.Println("NaN:", err)
	} else {
		fmt.Println(n)
	}

	// --------------- Loops ---------------
	// --- a loop with 3 components
	for i := 0; i < 3; i++ {
		fmt.Printf("Loops %d \n", i)
	}

	// --- a loop with a condition
	p := 0
	for p < 3 {
		fmt.Printf("p = %d \n", p)
		p++
	}

	// --- an infinite loop
	q := 0
	for {
		fmt.Printf("infinite loop \n")
		q++
		if q == 4 {
			break
		}
	}
}

// --------------- Functions ---------------
// --- No return
func greeting(firstName, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

// --- Return 1 value
func add(a, b int) int {
	return a + b
}

// --- Return 2 values
func swap(a, b int) (int, int) {
	return b, a
}
