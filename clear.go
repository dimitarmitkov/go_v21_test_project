package main

import "fmt"

func ClearAliceAndMap() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	fmt.Printf("map before clear: %v\n", m)

	clear(m)
	fmt.Printf("map after clear: %v\n", m)
	fmt.Printf("map len: %v\n", len(m))

	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice before clear: %v\n", s)

	clear(s)
	fmt.Printf("slice after clear: %v\n", s)
	fmt.Printf("slice len: %v\n", len(s))
}
