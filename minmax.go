package main

import "fmt"

func MinMax() {
	m := min(10, 20)
	fmt.Printf("shows min of 10 and 20: %d\n", m)

	m = max(10, 20)
	fmt.Printf("shows max of 10 and 20: %d\n", m)

	m = max(9, 5, 10)
	fmt.Printf("shows max of 5 and 9, but at least 10: %d\n", m)
}
