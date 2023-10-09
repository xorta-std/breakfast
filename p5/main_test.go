package main

import (
	"lc/rnd"
	"math/rand"
	"testing"
)

func TestIsPalindromic(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		{"cbbd", false},
		{"abc", false},
		{"abccba", true},
		{"b", true},
		{"alacazamazacala", true},
	}

	for i, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := isPalindromic([]rune(tc.input))
			if result != tc.expected {
				t.Errorf("[%d] Expected isPalindromic(%s) to be %v, but got %v", i, string(tc.input), tc.expected, result)
			}
		})
	}
}

func TestLongestPalindrome(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"abc", "a"},
		{"aba", "aba"},
		{"abccba", "abccba"},
		{"b", "b"},
		{"alacazamazacala", "alacazamazacala"},
		{"babad", "bab"},
		{"cbbd", "bb"},
	}

	for i, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := longestPalindrome(tc.input)
			if result != tc.expected {
				t.Errorf("[%d] Expected longestPalindrome(%s) to be %v, but got %v", i, tc.input, tc.expected, result)
			}
		})
	}
}

var testCase = rnd.RandomString(rand.Intn(10_000), rnd.LOWER_CASE)

func BenchmarkLongestPalindrome1(b *testing.B) {
	b.ResetTimer()
	longestPalindrome1(testCase)
}

func BenchmarkLongestPalindrome2(b *testing.B) {
	b.ResetTimer()
	longestPalindrome2(testCase)
}

func BenchmarkLongestPalindrome3(b *testing.B) {
	b.ResetTimer()
	longestPalindrome3(testCase)
}
