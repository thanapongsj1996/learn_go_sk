package main

import "fmt"

func main() {
	// -------- Zero Value --------
	var number int
	fmt.Println(number) // 0

	var str string
	fmt.Println(str) // ""

	var boolean bool
	fmt.Println(boolean) // false

	var p *string
	fmt.Println(p) // nil

	// -------- Array --------
	var fourNum [4]int
	fourNum[0] = 1
	fourNum[2] = 3
	fmt.Printf("%T %v \n", fourNum, fourNum) // [4]int [1 0 3 0]

	// -------- Slice --------
	var fourNumSlice []int
	fmt.Printf("%T %v \n", fourNumSlice, fourNumSlice) // []int []

	fourNumSlice = make([]int, 4)

	fourNumSlice[0] = 1
	fourNumSlice[2] = 3
	fmt.Printf("%T %v \n", fourNumSlice, fourNumSlice) // []int [1 0 3 0]

	fourNumSlice = append(fourNumSlice, 15)
	fmt.Printf("%T %v \n", fourNumSlice, fourNumSlice) // []int [1 0 3 0 15]

	var s []int
	s = make([]int, 4, 6)
	// 4 is length of the slice
	// 6 is capacity of the slice
	fmt.Printf("%T %v \n", s, s) // []int [0 0 0 0]
	// --->  []int [0 0 0 0 nil nil]
	fmt.Println(len(s)) // 4
	fmt.Println(cap(s)) // 6

	var fiveNum [5]int
	var nums []int
	// | 0 | 0 | 0 | 0 | 0 |
	// 0   1   2   3   4   5
	nums = fiveNum[1:3]
	fmt.Println(len(nums)) // 2
	fmt.Println(cap(nums)) // 4

	var s2 []int
	s2 = append(s2, 1, 2, 3, 4, 5)
	fmt.Println(len(s2)) // 5
	fmt.Println(cap(s2)) // 6

	s2 = append(s2, 6)
	fmt.Println(len(s2)) // 6
	fmt.Println(cap(s2)) // 6

	s2 = append(s2, 7)
	fmt.Println(len(s2)) // 7
	fmt.Println(cap(s2)) // 12
}
