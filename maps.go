package main

import (
	"fmt"
	"maps"
	"strings"
)

func MapsDemo() {
	m := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	m2 := maps.Clone(m)

	fmt.Printf("first map: %v\n", m)
	fmt.Printf("second map: %v\n", m2)

	m2["one"] = 100
	fmt.Printf("first map: %v\n", m)
	fmt.Printf("second map: %v\n", m2)

	m3 := map[string][]int{
		"key": {1, 2, 3},
	}
	m4 := maps.Clone(m3)
	fmt.Printf("m4 first element of cloned array, %v\n", m4["key"][0])
	fmt.Printf("m4 whole cloned array, %v\n", m4)
	m4["key"][0] = 100
	fmt.Printf("m3 where key was changed to 100, %v\n", m3["key"][0])
	fmt.Printf("m4 where key was changed to 100, %v\n", m4["key"][0])
	fmt.Printf("m3 whole cloned array, %v\n", m3)
	fmt.Printf("m4 whole cloned array, %v\n", m4)

	// Copy
	m5 := map[string][]int{
		"one": {1, 2, 3},
		"two": {4, 5, 6},
	}
	m6 := map[string][]int{
		"one": {7, 8, 9},
	}

	maps.Copy(m5, m6)
	fmt.Println("m6 is:", m6)

	// Equal
	me1 := map[int]string{
		1:    "one",
		10:   "Ten",
		1000: "THOUSAND",
	}
	me2 := map[int]string{
		1:    "one",
		10:   "Ten",
		1000: "THOUSAND",
	}
	me3 := map[int]string{
		1:    "one",
		10:   "ten",
		1000: "thousand",
	}

	fmt.Println(maps.Equal(me1, me2))
	fmt.Println(maps.Equal(me1, me3))

	// Equal Func
	mef1 := map[int]string{
		1:    "one",
		10:   "Ten",
		1000: "THOUSAND",
	}
	mef2 := map[int][]byte{
		1:    []byte("One"),
		10:   []byte("Ten"),
		1000: []byte("Thousand"),
	}
	eq := maps.EqualFunc(mef1, mef2, func(v1 string, v2 []byte) bool {
		return strings.ToLower(v1) == strings.ToLower(string(v2))
	})
	fmt.Println(eq)

	// Delete Func
	mdf := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
	}
	maps.DeleteFunc(mdf, func(k string, v int) bool {
		return v%2 != 0 // delete odd values
	})
	fmt.Println(mdf)
}
