package unittests

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add two integers
func Add[T Number](a, b T) T {
	return a + b
}

// Find the maximum of two integers
func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Reverse a string
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Check if a number is even
func IsEven(n int) bool {
	return n%2 == 0
}

// Concatenate two strings
func Concat(s1, s2 string) string {
	return s1 + s2
}
