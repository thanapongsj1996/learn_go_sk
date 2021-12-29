package main

import "fmt"

func main() {
	var i interface{}
	i = "TEXT"

	var s string

	// ถ้าให้ s = i จะเกิด error เพราะ i เป็น interface ส่วน s เป็น string
	// s = i

	// ต้องทำ type assertion ก่อน
	if v, ok := i.(string); ok {
		s = v
	}
	fmt.Println(s)

	// ---------- DEMO ----------
	var a, b Obj
	a = rectangle{width: 4, height: 4}
	b = triangle{base: 4, height: 4}

	fmt.Println(a.Area()) // 16
	fmt.Println(b.Area()) // 8

	if v, ok := b.(triangle); ok {
		v.base = 20
		fmt.Println(v.Area()) // 40
	}

}

type Obj interface {
	Area() float64
}
type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) Area() float64 {
	return r.width * r.height
}

type triangle struct {
	base   float64
	height float64
}

func (t triangle) Area() float64 {
	return 0.5 * t.base * t.height
}
