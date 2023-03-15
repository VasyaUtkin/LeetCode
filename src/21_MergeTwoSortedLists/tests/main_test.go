package main_test

import (
	solution21 "LeetCodeSolution/src/21_MergeTwoSortedLists/solution"
	"testing"
)

func TestCase1(t *testing.T) {

	list12 := solution21.ListNode{Val: 4, Next: nil}
	list11 := solution21.ListNode{Val: 2, Next: &list12}
	list1Head := solution21.ListNode{Val: 1, Next: &list11}

	list22 := solution21.ListNode{Val: 4, Next: nil}
	list21 := solution21.ListNode{Val: 3, Next: &list22}
	list2Head := solution21.ListNode{Val: 1, Next: &list21}

	content := solution21.MergeTwoLists(&list1Head, &list2Head)
	expected := [6]int{1, 1, 2, 3, 4, 4}
	counter := 0
	for content.Next != nil {
		if content.Val != expected[counter] {
			t.Fatalf(`expected content "%v", got %v`, expected[counter], content.Val)
		}
		counter++
		content = content.Next
	}

}

func TestCase2(t *testing.T) {

	var list1Head solution21.ListNode
	var list2Head solution21.ListNode

	content := solution21.MergeTwoLists(&list1Head, &list2Head)
	expected := 0
	if content.Val != expected {
		t.Fatalf(`expected content "%v", got %v`, expected, content.Val)
	}

}

func TestCase3(t *testing.T) {

	var list1Head solution21.ListNode
	list2Head := solution21.ListNode{Val: 0, Next: nil}

	content := solution21.MergeTwoLists(&list1Head, &list2Head)
	expected := [1]int{0}
	counter := 0
	for content.Next != nil {
		if content.Val != expected[counter] {
			t.Fatalf(`expected content "%v", got %v`, expected[counter], content.Val)
		}
		counter++
		content = content.Next
	}

}
