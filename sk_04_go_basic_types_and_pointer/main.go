package main

import "fmt"

func main() {
	var ok bool = true
	fmt.Printf("%T %v \n", ok, ok)
	var s string = "Hi"
	fmt.Printf("%T %v \n", s, s)
	var i int = 1
	fmt.Printf("%T %v \n", i, i)
	// int8 int 16 int 32 int64
	// uint uint8 uint16 uint32 uint64 uintptr

	// byte === uint8
	var c byte = 1
	fmt.Printf("%T %v \n", c, c)

	// rune === int32 (represents a Unicode code point
	var r rune = 1
	fmt.Printf("%T %v \n", r, r)

	var f float32 = 2.3
	fmt.Printf("%T %v \n", f, f)
	// float32 float64

	var z complex64 = 1 + 2i
	fmt.Printf("%T %v \n", z, z)
	// complex64 ==> real number 32 bit and imaginary number 32 bit
	// complex128 ==> real number 64 bit and imaginary number 64 bit

	// --------------- Pointer ---------------
	var str string = "string value"
	fmt.Printf("%v \n", str)  // string value
	fmt.Printf("%v \n", &str) // 0xc000096230

	var str2 string
	var pStr2 *string
	pStr2 = &str2
	str2 = "Hello"
	fmt.Printf("%s \n", *pStr2) // Hello

	var age int = 25
	fmt.Printf("age before change : %d \n", age)

	double(&age)

	fmt.Printf("age after change : %d \n", age)

	var name string = "Thanapong"
	fmt.Printf("name before change : %s \n", name)
	appendGreeting(&name)
	fmt.Printf("name after change : %s \n", name)
}

/*
	จงแก้ไข func ต่อไปนี้ ให้สามารถทำงานได้อย่างถูกต้อง โดยที่

	func double จะเพิ่มค่า value ที่อยู่ใน pointer เป็น 2 เท่า
	func appendGreeting จะเติม "Hi, " เข้าไปข้างหน้า
*/
func double(d *int) {
	*d = *d * 2
}
func appendGreeting(s *string) {
	*s = "Hi, " + *s
	//*s = fmt.Sprintf("Hi, %s", *s)
}
