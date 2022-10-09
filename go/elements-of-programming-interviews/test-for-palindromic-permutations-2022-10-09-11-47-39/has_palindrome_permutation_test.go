package main

import "testing"

func TestHasPalindromePermutation(t *testing.T) {
	t.Run("It should return `true` if input string len < 2", func(t *testing.T) {
		if !hasPalindromePermutation("") {
			t.Error("Expected empty string to return true")
		}

		if !hasPalindromePermutation("a") {
			t.Error("Expected string with 1 character to return true")
		}
	})

	t.Run("It should return `false` if there are 2 unique characters in string len == 2", func(t *testing.T) {
		got := hasPalindromePermutation("ab")
		want := false

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("It should return `true` if there are 2 unique chars that make up a palindromic permutation", func(t *testing.T) {
		// "dad" => true
		got := hasPalindromePermutation("add")
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})

	t.Run("It should return `true` if there are more than 2 unique chars that make up a palindromic permutation", func(t *testing.T) {
		// "xzyyyyyzx" => true
		got := hasPalindromePermutation("xxzzyyyyy")
		want := true

		if got != want {
			t.Errorf("got %t want %t", got, want)
		}
	})
}
