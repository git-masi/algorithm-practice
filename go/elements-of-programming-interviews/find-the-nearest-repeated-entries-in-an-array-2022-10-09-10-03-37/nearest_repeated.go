package main

import "math"

type Words map[string][]int

func (w Words) GetValueAtOffset(word string, offset int) int {
	return w[word][len(w[word])-offset]
}

type Nearest []int

func (n Nearest) GetDistance() int {
	len := len(n)

	if len < 2 {
		return math.MaxInt
	}

	return n[len-1] - n[len-2]
}

func nearestRepeated(s []string) []int {
	nearest := Nearest{}
	words := Words{}

	for i, word := range s {
		if _, ok := words[word]; !ok {
			words[word] = []int{}
		}

		words[word] = append(words[word], i)

		lenN := len(nearest)

		if len(words[word]) > 1 {
			if lenN == 0 || words.GetValueAtOffset(word, 1)-words.GetValueAtOffset(word, 2) < nearest.GetDistance() {
				nearest = Nearest{words.GetValueAtOffset(word, 2), words.GetValueAtOffset(word, 1)}
			}
		}
	}

	return nearest
}
