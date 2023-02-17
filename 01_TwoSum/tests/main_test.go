package main_test

import (
	"LeetCodeSolution/01_TwoSum/solution"
	"testing"
)

func TestCase1(t *testing.T) {

	content := solution.TwoSum([]int{2, 7, 11, 15}, 9)
	expected := []int{0, 1}
	if content[0] != expected[0] {
		t.Fatalf(`expected content "%d", got %d`, expected[0], content[0])
	}

	if content[1] != expected[1] {
		t.Fatalf(`expected content "%d", got %d`, expected[1], content[1])
	}

}

func TestCase2(t *testing.T) {

	content := solution.TwoSum([]int{3, 2, 4}, 6)
	expected := []int{1, 2}
	if content[0] != expected[0] {
		t.Fatalf(`expected content "%d", got %d`, expected[0], content[0])
	}

	if content[1] != expected[1] {
		t.Fatalf(`expected content "%d", got %d`, expected[1], content[1])
	}

}

func TestCase3(t *testing.T) {

	content := solution.TwoSum([]int{3, 3}, 6)
	expected := []int{0, 1}
	if content[0] != expected[0] {
		t.Fatalf(`expected content "%d", got %d`, expected[0], content[0])
	}

	if content[1] != expected[1] {
		t.Fatalf(`expected content "%d", got %d`, expected[1], content[1])
	}

}
