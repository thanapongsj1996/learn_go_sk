package main

import "fmt"

func main() {
	// switch case ใน go จะ break ให้เองจะเจอ condition ที่ตรง
	var name = "boy"
	switch name {
	case "girl":
		fmt.Println("Hi, Girl")
	case "boy":
		fmt.Println("Hi, Boy")
	default:
		fmt.Println("Hi, Everyone")
	}
	// Hi, Boy

	// ถ้าตรงการให้มันไม่ break ให้ใส่ fallthrough
	switch name {
	case "girl":
		fmt.Println("Hi, Girl")
	case "boy":
		fmt.Println("Hi, Boy")
		fallthrough
	default:
		fmt.Println("Hi, Everyone")
	}
	// Hi, Boy
	// Hi, Everyone

	// switch case สามารถใส่ condition ของตัวเองได้
	var num = 10
	switch {
	case num == 5:
		fmt.Println("five")
	case num*2 == 20:
		fmt.Println("num * 2 = 20 -> num = 10")
	}
	// num * 2 = 20 -> num = 10

	// switch case ทำ switch type ได้
	var money interface{}
	money = 20.23
	switch v := money.(type) {
	case int:
		fmt.Println("money is int ", v)
	case float32:
		fmt.Println("money is float32 ", v)
	case float64:
		fmt.Println("money is float64 ", v)
	default:
		fmt.Println("we don't know ", v)
	}
	// money is float64  20.23
}
