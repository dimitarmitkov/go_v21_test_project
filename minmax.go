package main

import "fmt"

func MinMax() {
	mMin := min(10, 20)
	fmt.Printf("shows min of 10 and 20: %d\n", mMin)

	mMax := max(10, 20)
	fmt.Printf("shows max of 10 and 20: %d\n", mMax)

	m10 := max(9, 5, 10)
	fmt.Printf("shows max of 5 and 9, but at least 10: %d\n", m10)

	mm10 := min(9, 5, 2)
	fmt.Println("min of 9 and 5, but at least 2: ", mm10)

	s := max("fed", "def")
	fmt.Println("shows min of 'abc' and 'def'", s)

	sCap := max("ABC", "abc")
	fmt.Println("shows max between 'ABC' and 'abc'", sCap)

	sf := min("aaa", "def", "abc")
	fmt.Println("shows min of 'abc' and 'def'", sf)

	flt := min(1, 1.2)
	fmt.Println("min of float values: ", flt)

	fltmx := max(1, 1.2)
	fmt.Println("min of float values: ", fltmx)

	fltmx = max(1, 1.2, 3.2)
	fmt.Println("min of float values: (3.2) ", fltmx)
}
