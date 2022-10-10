package main

import (
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	t.Run("It should do nothing if the input matrix is empty", func(t *testing.T) {
		m := Matrix{}

		floodFill(m, 1, 0, 0)

		if !reflect.DeepEqual(m, Matrix{}) {
			t.Error("Expected an empty matrix")
		}
	})

	t.Run("It should change all of the values to '1'", func(t *testing.T) {
		m := Matrix{{0}}
		want := Matrix{{1}}

		floodFill(m, 1, 0, 0)

		if !reflect.DeepEqual(want, m) {
			t.Errorf("want %v got %v", want, m)
		}
	})

	t.Run("It should change all of the values to '2'", func(t *testing.T) {
		m := Matrix{
			{0, 0},
			{0, 0},
		}
		want := Matrix{
			{2, 2},
			{2, 2},
		}

		floodFill(m, 2, 0, 0)

		if !reflect.DeepEqual(want, m) {
			t.Errorf("want %v got %v", want, m)
		}
	})

	t.Run("It should change one of the values to '3'", func(t *testing.T) {
		m := Matrix{
			{0, 2},
			{0, 0},
		}
		want := Matrix{
			{0, 3},
			{0, 0},
		}

		floodFill(m, 3, 0, 1)

		if !reflect.DeepEqual(want, m) {
			t.Errorf("want %v got %v", want, m)
		}
	})

	t.Run("It should change two of the values to '4'", func(t *testing.T) {
		m := Matrix{
			{0, 3},
			{0, 3},
		}
		want := Matrix{
			{0, 4},
			{0, 4},
		}

		floodFill(m, 4, 0, 1)

		if !reflect.DeepEqual(want, m) {
			t.Errorf("want %v got %v", want, m)
		}
	})

	t.Run("It should change one of the values to '5'", func(t *testing.T) {
		// Diagonal cells with the same value are not considered to be adjacent
		m := Matrix{
			{0, 4},
			{4, 0},
		}
		want := Matrix{
			{0, 5},
			{4, 0},
		}

		floodFill(m, 5, 0, 1)

		if !reflect.DeepEqual(want, m) {
			t.Errorf("want %v got %v", want, m)
		}
	})
}
