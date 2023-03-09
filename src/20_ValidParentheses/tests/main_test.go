package main_test

import (
	"LeetCodeSolution/src/20_ValidParentheses/solution"
	"testing"
)

func TestCase1(t *testing.T) {

	content := solution.IsValid("()")
	expected := true
	if content != expected {
		t.Fatalf(`expected content "%v", got %v`, expected, content)
	}

}

func TestCase2(t *testing.T) {

	content := solution.IsValid("()[]{}")
	expected := true
	if content != expected {
		t.Fatalf(`expected content "%v", got %v`, expected, content)
	}

}

func TestCase3(t *testing.T) {

	content := solution.IsValid("(]")
	expected := false
	if content != expected {
		t.Fatalf(`expected content "%v", got %v`, expected, content)
	}

}

func TestCase4(t *testing.T) {

	content := solution.IsValid("[")
	expected := false
	if content != expected {
		t.Fatalf(`expected content "%v", got %v`, expected, content)
	}

}
