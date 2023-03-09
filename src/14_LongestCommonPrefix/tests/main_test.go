package main_test

import (
	"LeetCodeSolution/src/14_LongestCommonPrefix/solution"
	"testing"
)

func TestCase1(t *testing.T) {

	content := solution.LongestCommonPrefix([]string{"flower", "flow", "flight"})
	expected := "fl"
	if content != expected {
		t.Fatalf(`expected content "%v", got %v`, expected, content)
	}

}

func TestCase2(t *testing.T) {

	content := solution.LongestCommonPrefix([]string{"dog", "racecar", "car"})
	expected := ""
	if content != expected {
		t.Fatalf(`expected content "%v", got %v`, expected, content)
	}

}
