package main

import (
	solution88 "LeetCodeSolution/src/88_MergeSortedArray/solution"
	"fmt"
)

func main() {

	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	solution88.Merge(nums1, m, nums2, n)
	fmt.Print(nums1)
}
