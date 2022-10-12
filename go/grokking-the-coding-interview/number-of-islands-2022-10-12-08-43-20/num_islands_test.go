package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNumIslandsRecursive(t *testing.T) {
	t.Run("It should return '0' for an empty matrix", func(t *testing.T) {
		got := numIslandsRecursive(Matrix{})
		want := 0

		assertInt(got, want, t)
	})

	t.Run("It should return '1' for a matrix with all land cells", func(t *testing.T) {
		got := numIslandsRecursive(Matrix{
			{1, 1},
			{1, 1},
		})
		want := 1

		assertInt(got, want, t)
	})

	t.Run("It should return '0' for a matrix with all water cells", func(t *testing.T) {
		got := numIslandsRecursive(Matrix{
			{0, 0},
			{0, 0},
		})
		want := 0

		assertInt(got, want, t)
	})

	t.Run("It should return '2' for a matrix with two diagonal land cells", func(t *testing.T) {
		got := numIslandsRecursive(Matrix{
			{0, 1},
			{1, 0},
		})
		want := 2

		assertInt(got, want, t)
	})

	t.Run("It should return '3' for three islands of different sizes", func(t *testing.T) {
		got := numIslandsRecursive(Matrix{
			{0, 0, 1, 0, 0},
			{0, 1, 1, 0, 0},
			{0, 0, 0, 1, 1},
			{1, 1, 0, 0, 0},
			{0, 1, 1, 1, 1},
		})
		want := 3

		assertInt(got, want, t)
	})
}

func TestNumIslandsIterative(t *testing.T) {
	t.Run("It should return '0' for an empty matrix", func(t *testing.T) {
		got := numIslandsIterative(Matrix{})
		want := 0

		assertInt(got, want, t)
	})

	t.Run("It should return '1' for a matrix with all land cells", func(t *testing.T) {
		got := numIslandsIterative(Matrix{
			{1, 1},
			{1, 1},
		})
		want := 1

		assertInt(got, want, t)
	})

	t.Run("It should return '0' for a matrix with all water cells", func(t *testing.T) {
		got := numIslandsIterative(Matrix{
			{0, 0},
			{0, 0},
		})
		want := 0

		assertInt(got, want, t)
	})

	t.Run("It should return '2' for a matrix with two diagonal land cells", func(t *testing.T) {
		got := numIslandsIterative(Matrix{
			{0, 1},
			{1, 0},
		})
		want := 2

		assertInt(got, want, t)
	})

	t.Run("It should return '3' for three islands of different sizes", func(t *testing.T) {
		got := numIslandsIterative(Matrix{
			{0, 0, 1, 0, 0},
			{0, 1, 1, 0, 0},
			{0, 0, 0, 1, 1},
			{1, 1, 0, 0, 0},
			{0, 1, 1, 1, 1},
		})
		want := 3

		assertInt(got, want, t)
	})
}

func assertInt(got int, want int, t *testing.T) {
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

// Benchmarks
var recursiveResult int
var iterativeResult int

var sizes map[int]Matrix = map[int]Matrix{
	1:      createMatrix(1),
	10:     createMatrix(10),
	100:    createMatrix(100),
	1_000:  createMatrix(1_000),
	10_000: createMatrix(10_000),
}

func BenchmarkNumIslandsRecursive(b *testing.B) {
	for s, m := range sizes {
		b.Run(fmt.Sprintf("Matrix size: %d", s), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r := numIslandsRecursive(m)
				recursiveResult = r
			}
		})
	}
}

func BenchmarkNumIslandsIterative(b *testing.B) {
	for s, m := range sizes {
		b.Run(fmt.Sprintf("Matrix size: %d", s), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				r := numIslandsIterative(m)
				iterativeResult = r
			}
		})
	}
}

func createMatrix(size int) Matrix {
	m := make(Matrix, 0, size)

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		s := make([]int, 0, size)

		for j := 0; j < size; j++ {
			s = append(s, rand.Intn(2))
		}

		m = append(m, s)
	}

	return m
}
