package main_test

import (
	solution88 "LeetCodeSolution/src/88_MergeSortedArray/solution"
	"testing"
)

func TestCase1(t *testing.T) {

	content := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	solution88.Merge(content, m, nums2, n)
	expected := [6]int{1, 2, 2, 3, 5, 6}
	for i, v := range expected {
		if content[i] != v {
			t.Fatalf(`expected content "%v", got %v`, v, content[i])
		}
	}

}

func TestCase2(t *testing.T) {

	content := []int{1}
	m := 1
	var nums2 []int
	n := 0
	solution88.Merge(content, m, nums2, n)
	expected := [1]int{1}
	for i, v := range expected {
		if content[i] != v {
			t.Fatalf(`expected content "%v", got %v`, v, content[i])
		}
	}

}

func TestCase3(t *testing.T) {

	var content = []int{0}
	m := 0
	nums2 := []int{1}
	n := 1
	solution88.Merge(content, m, nums2, n)
	expected := [1]int{1}
	for i, v := range expected {
		if content[i] != v {
			t.Fatalf(`expected content "%v", got %v`, v, content[i])
		}
	}

}
