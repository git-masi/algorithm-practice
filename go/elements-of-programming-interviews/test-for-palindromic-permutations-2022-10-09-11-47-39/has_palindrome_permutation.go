package main

func hasPalindromePermutation(s string) bool {
	// This only works for ASCII chars
	// Use len([]rune(s)) for other char sets
	if len(s) < 2 {
		return true
	}

	freq := map[rune]int{}
	count := 0

	for _, char := range s {
		if _, ok := freq[char]; !ok {
			freq[char] = 0
		}

		freq[char]++
	}

	for _, v := range freq {
		if v%2 != 0 {
			count++
		}

		if count > 1 {
			return false
		}
	}

	return true
}
