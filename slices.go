package main

import (
	"fmt"
	"slices"
)

func SlicesExamples() {
	exampleSlice := []string{"one", "two", "three", "four", "five"}
	exampleSliceAddSix := []string{"one", "two", "three", "four", "five", "six"}
	exampleSliceCopyTrue := []string{"one", "two", "three", "four", "five"}

	doesSliceContainsOne := slices.Contains(exampleSlice, "one")
	areSlicesEqualFalse := slices.Equal(exampleSlice, exampleSliceAddSix)
	areSlicesEqualTrue := slices.Equal(exampleSlice, exampleSliceCopyTrue)

	fmt.Printf("check if slice contains element, should be true: %t\n", doesSliceContainsOne)
	fmt.Printf("are slices equal, should be false: %t\n", areSlicesEqualFalse)
	fmt.Printf("are slices equal, should be true: %t\n", areSlicesEqualTrue)
}
