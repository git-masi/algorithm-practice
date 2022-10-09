package main

import "testing"

func TestIsConstructable(t *testing.T) {
	t.Run("it should return true when `letter` is an empty string", func(t *testing.T) {
		got := isConstructable("", "")
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("it should return false when `magazine` is an empty string and `letter` is not", func(t *testing.T) {
		got := isConstructable("a", "")
		want := false

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("it should return true when the frequency of a single char in `magazine` is equal to the frequency of that char in `letter`", func(t *testing.T) {
		got := isConstructable("a", "a")
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("it should return false when the frequency of a single char in `magazine` is less than to the frequency of that char in `letter`", func(t *testing.T) {
		got := isConstructable("aa", "a")
		want := false

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("spaces should not effect the result", func(t *testing.T) {
		got := isConstructable("aa", "a  b")
		want := false

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("it should return false when there is a char in `letter` that is not in `magazine`", func(t *testing.T) {
		got := isConstructable("ab", "aa")
		want := false

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("it should return true when the frequency of all chars in `magazine` is equal to the frequency of all chars in `letter`", func(t *testing.T) {
		got := isConstructable("aabb", "aabb")
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("it should return true when the frequency of all chars in `magazine` is equal to the frequency of all chars in `letter` regardless of order", func(t *testing.T) {
		got := isConstructable("aabb", "bbaa")
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("it should return true when the frequency of all chars in `magazine` is equal to the frequency of all chars in `letter` regardless of order and spaces", func(t *testing.T) {
		got := isConstructable("aabb", " b b a a ")
		want := true

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
