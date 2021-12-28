package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		These are same
		->  m := make(map[string]string)
		->  m := map[string]string{}
		-----------------------------------------
		var m map[string]string --> nil
	*/

	m := map[string]string{}
	m["a"] = "Apple"

	m2 := map[string]string{
		"a": "Apple",
		"b": "Banana",
	}
	fmt.Println(m2) // map[a:Apple b:Banana]

	/*
		จงแก้ไข func wordCount ให้สามารถนับจำนวนคำต่างๆ ได้อย่างถูกต้อง เช่น
		ถ้า input = "Apple Banana Apple Banana apple"
		ต้อง return map[string]int{
			"Apple": 2,
			"Banana": 2,
			"apple": 1,
		}
	*/
	fmt.Println(wordCount("Apple Banana Apple Banana apple"))
}

func wordCount(s string) map[string]int {
	mySlice := strings.Split(s, " ")
	myResult := map[string]int{}
	for _, str := range mySlice {
		myResult[str] += 1
	}

	return myResult
}
