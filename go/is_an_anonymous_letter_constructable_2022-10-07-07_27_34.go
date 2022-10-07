package main

func isConstructable(letter, magazine string) bool {
	if letter == "" {
		return true
	}

	if magazine == "" {
		return false
	}

	freq := map[rune]int{}

	for _, char := range magazine {
		if _, ok := freq[char]; !ok {
			freq[char] = 0
		}

		freq[char]++
	}

	for _, char := range letter {
		if _, ok := freq[char]; !ok {
			return false
		}

		freq[char]--

		if freq[char] < 0 {
			return false
		}
	}

	return true
}
