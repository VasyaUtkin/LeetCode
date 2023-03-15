package main

import (
	solution21 "LeetCodeSolution/src/21_MergeTwoSortedLists/solution"
	"fmt"
)

func main() {

	list12 := solution21.ListNode{Val: 4, Next: nil}
	list11 := solution21.ListNode{Val: 2, Next: &list12}
	list1Head := solution21.ListNode{Val: 1, Next: &list11}

	list22 := solution21.ListNode{Val: 4, Next: nil}
	list21 := solution21.ListNode{Val: 3, Next: &list22}
	list2Head := solution21.ListNode{Val: 1, Next: &list21}

	listMerge := solution21.MergeTwoLists(&list1Head, &list2Head)

	for listMerge.Next != nil {
		fmt.Printf("%d ", listMerge.Val)
		listMerge = listMerge.Next
	}

}
