package main

/*
	Go test
	1. The suffix of filename is _test.go
	2. The name of the unit test function starts with Test
	3. The type of function with 1 parameter is *testing.T
	4. (optional) A package name can have _test as a suffix
*/

func sumOfFirst(n int) int {
	sum := 0
	for n > 0 {
		sum += n
		n--
	}

	return sum
}
