package main

import "fmt"

func main() {
	thisIsVariadic("a", "b", "c") // a,b,c,

	myStr := []string{"b", "o", "y"}
	thisIsVariadic(myStr...) // b,o,y,
}

func thisIsVariadic(s ...string) {
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i] + ",")
	}
}
