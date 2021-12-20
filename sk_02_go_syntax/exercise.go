package main

import "fmt"

/*
	Variables, Constants, Functions
	จงแก้ไข func ต่อไปนี้ ให้สามารถ return ค่าออกมาได้อย่างถูกต้อง
		- greeting("Pallat") -> "Hello, Pallat"
		- greetingWithAge("Pallat", 30) -> "Hello, Pallat. You are 30 years old."
		- greetingWithAgeAndDrink("Pallat", 30, "Cola") -> "Hello, Pallat. You are 30 years old and your favorite drink is Cola."
*/

func greetingExercise(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

func greetingWithAge(name string, age int) string {
	return fmt.Sprintf("Hello, %s. You are %d years old.", name, age)
}

func greetingWithAgeAndDrink(name string, age int, drink string) string {
	return fmt.Sprintf("Hello, %s. You are %d years old and your favorite drink is %s.", name, age, drink)
}

/*
	Conditions
	จงแก้ไข func isOdd ให้สามารถ return boolean ได้อย่างถูกต้อง โดยที่ถ้าใส่ค่าเป็นเลขคี่จะ return true และเลขคู่จะ return false
*/
func isOdd(n int) bool {
	if n%2 == 0 {
		return false
	}
	return true
}

/*
	Loops
	จงแก้ไข func ต่อไปนี้ ให้สามารถทำงานได้อย่างถูกต้อง
	- sumOfFirst(3) should return 3+2+1=6

	- isPrime()
	isPrime(1) should return false
	isPrime(2) should return true
	isPrime(3) should return true
	isPrime(4) should return false
*/
func sumOfFirst(n int) int {
	start := 0
	sum := 0
	for start <= n {
		sum += start
		start++
	}

	return sum
}

func isPrime(n int) bool {
	if n == 1 {
		return false
	}

	var i = 2
	for i <= n/2 {
		if n%i == 0 {
			return false
		}
		i += 1
	}

	return true
}
