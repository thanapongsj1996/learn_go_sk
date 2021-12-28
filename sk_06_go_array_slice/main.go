package main

import "fmt"

func main() {
	var s []int
	if s == nil {
		fmt.Println("empty slice is nil")
	}

	// %d ตัวเลข
	// %p เอา address ออกมา

	a := [...]int{1, 3, 5, 7, 9}
	s = a[:]

	fmt.Printf("%d %d %p %p \n", len(s), cap(s), s, &a)
	// 5 5 0xc000122060 0xc000122060

	s = append(s, 11, 13)
	fmt.Printf("%d %d %p %p \n", len(s), cap(s), s, &a)
	// 7 10 0xc000130000 0xc000122060
}
