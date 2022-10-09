package main

import (
	"reflect"
	"testing"
)

func TestNearestRepeated(t *testing.T) {
	t.Run("It should return an empty slice if the input is an empty slice", func(t *testing.T) {
		got := nearestRepeated([]string{})
		want := 0

		if len(got) != want {
			t.Errorf("got %d want %d", len(got), want)
		}
	})

	t.Run("It should return two indexes if there is on pair of repeated words", func(t *testing.T) {
		got := nearestRepeated([]string{"a", "fox", "a", "box"})
		want := []int{0, 2}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("It should return an empty slice if there are not repeated words", func(t *testing.T) {
		got := nearestRepeated([]string{"the", "man", "with", "a", "plan"})
		want := 0

		if len(got) != want {
			t.Errorf("got %d want %d", len(got), want)
		}
	})

	t.Run("It should return the correct indexes if there is more than one repeated word", func(t *testing.T) {
		got := nearestRepeated([]string{"a", "b", "c", "b", "a"})
		want := []int{1, 3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
