package test_pal

import (
	"LeetCodeSolution/src/09_Palindrome/solution"
	"testing"
)

func TestCase1(t *testing.T) {

	content := solution.IsPalindrome(121)
	expected := true
	if content != expected {
		t.Fatalf(`expected content "%v", got %v`, expected, content)
	}

}

func TestCase2(t *testing.T) {

	content := solution.IsPalindrome(-121)
	expected := false
	if content != expected {
		t.Fatalf(`expected content "%v", got %v`, expected, content)
	}

}

func TestCase3(t *testing.T) {

	content := solution.IsPalindrome(10)
	expected := false
	if content != expected {
		t.Fatalf(`expected content "%v", got %v`, expected, content)
	}

}

func TestCad(t *testing.T) {
	content := solution.IsPalindrome(10)
	expected := false
	if content != expected {
		t.Fatalf(`expected content "%v", got %v`, expected, content)
	}
}
