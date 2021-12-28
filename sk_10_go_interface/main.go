package main

import (
	"fmt"
	"math"
)

func main() {
	// Empty interface
	var a interface{}
	a = 1
	fmt.Printf("a: %v %T \n", a, a) // a: 1 int
	a = "Hello"
	fmt.Printf("a: %v %T \n", a, a) // a: Hello string

	/*
		 	มองว่า ข้างใน interface เก็บของไว้ 2 อย่าง
			- pointer ของ ของที่ถือ
			- type ของ ของที่ถืออยู่
	*/

}

// Implemented Implicitly
type Shape interface {
	GetArea() float32
}

type Circle struct {
	Radius float32
}

func (c *Circle) GetArea() float32 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float32
	Height float32
}

func (r *Rectangle) GetArea() float32 {
	return r.Width * r.Height
}
