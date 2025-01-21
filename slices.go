package main

import (
	"fmt"
	"slices"
)

func SlicesExamples() {
	exampleSlice := []string{"one", "two", "three", "four", "five"}
	exampleSliceAddSix := []string{"one", "two", "three", "four", "five", "six"}
	exampleSliceCopyTrue := []string{"one", "two", "three", "four", "five"}

	//Contains and Equal
	doesSliceContainsOne := slices.Contains(exampleSlice, "one")
	areSlicesEqualFalse := slices.Equal(exampleSlice, exampleSliceAddSix)
	areSlicesEqualTrue := slices.Equal(exampleSlice, exampleSliceCopyTrue)

	fmt.Printf("check if slice contains element, should be true: %t\n", doesSliceContainsOne)
	fmt.Printf("are slices equal, should be false: %t\n", areSlicesEqualFalse)
	fmt.Printf("are slices equal, should be true: %t\n", areSlicesEqualTrue)

	// Delete
	exampleSliceAddSixDeletable := []string{"one", "two", "three", "four", "five", "six"}
	newSliceAfterDelete := slices.Delete(exampleSliceAddSixDeletable, 0, 2)
	fmt.Println("slice after deletion last element: ", newSliceAfterDelete)
	fmt.Println("slice after deletion last element length: ", len(newSliceAfterDelete), cap(newSliceAfterDelete))
	fmt.Println("original slice and its length : ", exampleSliceAddSixDeletable, len(exampleSliceAddSixDeletable), cap(exampleSliceAddSixDeletable))

	// Clip
	newSliceAfterDelete = slices.Clip(newSliceAfterDelete)
	fmt.Println("slice after deletion last element length and fix it by Clip: ", len(newSliceAfterDelete), cap(newSliceAfterDelete))

	// Clone
	originalCloneableSlice := []string{"A", "B", "C", "D", "E", "F", "G"}
	clonedSliceFromOriginal := slices.Clone(originalCloneableSlice)

	fmt.Println("original:", originalCloneableSlice)
	fmt.Println("clone:", clonedSliceFromOriginal)

	clonedSliceFromOriginal[0] = "Q"
	fmt.Println("original:", originalCloneableSlice)
	fmt.Println("clone:", clonedSliceFromOriginal)

	// Compact
	sliceToBeCompacted := []int{1, 1, 2, 2, 3, 3, 4, 5, 6, 6, 6, 7, 7, 7, 7, 7, 7, 8}
	fmt.Println(sliceToBeCompacted)

	compactedSlice := slices.Compact(sliceToBeCompacted)
	fmt.Println(compactedSlice)
	fmt.Println(sliceToBeCompacted)

	// compact doesn't work in slice like
	stc := []int{1, 2, 3, 1, 2, 3, 1, 2, 3}
	cts := slices.Compact(stc)

	fmt.Println(stc)
	fmt.Println(cts)

	index5, found5 := slices.BinarySearch(sliceToBeCompacted, 5)
	fmt.Println("found 5: ", index5, found5)
	index7, found7 := slices.BinarySearch(sliceToBeCompacted, 7)
	fmt.Println("found 7: ", index7, found7)
	index8, found8 := slices.BinarySearch(sliceToBeCompacted, 8)
	fmt.Println("found 7: ", index8, found8)

	// Sort
	slices.Sort(stc)
	fmt.Println(stc)

	// Max Min
	maxValue := slices.Max(sliceToBeCompacted)
	fmt.Println("max value: ", maxValue)
	minValue := slices.Min(sliceToBeCompacted)
	fmt.Println("min value: ", minValue)

	// important Index doesn't have in ming slice len
	indexValue := slices.Index(sliceToBeCompacted, 8)
	fmt.Println("index: ", indexValue)
	fmt.Println("cap, len: ", cap(sliceToBeCompacted), len(sliceToBeCompacted))

	// Grow
	initialSliceToBeGrown := []int{1, 1, 2, 2, 3, 3, 4, 5, 6, 6, 6, 7, 7, 7, 7, 7, 7, 8}
	fmt.Println("cap, len before grow with 10: ", cap(initialSliceToBeGrown), len(initialSliceToBeGrown))
	grownSlice := slices.Grow(initialSliceToBeGrown, 25)
	fmt.Println("cap, len after grow with 10: ", cap(grownSlice), len(grownSlice))
	fmt.Println(grownSlice)

	// Reverse slice
	slices.Reverse(initialSliceToBeGrown)
	fmt.Println("reversed slice: ", initialSliceToBeGrown)

	// Replace
	names := []string{"Alice", "Bob", "Vera", "Zac"}
	names = slices.Replace(names, 1, 3, "Bill", "Billie", "Cat")
	fmt.Println(names)

	nums := []int{1, 2, 3, 4}
	nums = slices.Replace(nums, 1, 2, 44, 45)
	fmt.Println(nums)

	// Insert and IsSorted
	notSorted := []int{1, 3, 45, 2, 11, 5, 66, 778, 10}
	isSortedCheck := slices.IsSorted(notSorted)
	fmt.Println("check if is sorted: ", isSortedCheck)
	slices.Sort(notSorted)
	isSortedCheck = slices.IsSorted(notSorted)
	fmt.Println("check if is sorted (after sort): ", isSortedCheck)

	notSorted = slices.Insert(notSorted, 1, 100)
	fmt.Println("inserted 100: ", notSorted)
}
